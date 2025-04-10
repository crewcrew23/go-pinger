package pinger

import (
	_const "go-pinger/internal/const"
	"go-pinger/internal/errs"
	"golang.org/x/net/icmp"
	"sync"
	"time"
)

type pinger struct {
	ping        []*Ping
	globalRetry int
}

func New() *pinger {
	return &pinger{
		ping:        []*Ping{},
		globalRetry: _const.DefaultRetryCount,
	}
}

func (p *pinger) SendPing(target string, retry int) {
	pp := p.preparePing(target, retry)
	p.ping = append(p.ping, pp)
	p.ping[0].initInternalData(retry)

	if retry <= 0 {
		retry = _const.DefaultRetryCount
	}

	for i := retry; i > 0; i-- {
		p.ping[0].send()
		time.Sleep(500 * time.Millisecond)
	}

	p.ping[0].showInfo()
}

func (p *pinger) preparePing(addr string, retry int) *Ping {

	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	errs.HandleFatal(err, "Error of creating connection")

	var r int
	if retry <= 0 {
		r = _const.DefaultRetryCount
	} else {
		r = retry
	}
	pp := &Ping{
		conn:  conn,
		addr:  addr,
		retry: r,
	}
	pp.initInternalData(p.globalRetry)
	return pp
}

func (p *pinger) SendFewPings(addrs []string, retry int) {
	if retry <= 0 {
		p.globalRetry = _const.DefaultRetryCount
	}

	var wg sync.WaitGroup

	for _, addr := range addrs {
		pp := p.preparePing(addr, retry)
		p.ping = append(p.ping, pp)

		wg.Add(1)
		go func(p *Ping) {
			defer wg.Done()
			for i := p.retry; i > 0; i-- {
				p.send()
				time.Sleep(500 * time.Millisecond)
			}
		}(pp)
	}

	wg.Wait()
	p.showAllInfo()
}

func (p *pinger) showAllInfo() {
	for _, p := range p.ping {
		p.showInfo()
	}
}

func (p *pinger) Close() {
	if len(p.ping) > 0 {
		for i := 0; i < len(p.ping); i++ {
			p.ping[i].Close()
		}
	}
}
