package app

import (
	"fmt"
	_const "go-pinger/internal/const"
	"go-pinger/internal/errs"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"log"
	"net"
	"os"
	"time"
)

type pinger struct {
	conn              *icmp.PacketConn
	retry             int
	countSuccessReply int
	rttList           []time.Duration
}

func New() *pinger {
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	errs.HandleFatal(err, "Error of creating connection")

	return &pinger{conn: conn}
}

func (p *pinger) SendPing(target string, retry int) {
	p.initInternalData(retry)

	if retry <= 0 {
		retry = 4
	}

	for i := retry; i > 0; i-- {
		p.send(target)
		time.Sleep(500 * time.Millisecond)
	}

	minTime, maxTime := p.calcMinAndMaxTime()

	fmt.Printf("--- %s ping statistics ---\n", target)
	fmt.Printf("%d packets transmitted, %d received, %.2f%% packet loss\n",
		p.retry,
		p.countSuccessReply,
		p.calcPacketLost(),
	)
	fmt.Printf("rtt min/avg/max = %.2f/%.2f/%.2f ms\n",
		float64(minTime.Microseconds())/1000,
		float64(p.calcAvgTime().Microseconds())/1000,
		float64(maxTime.Microseconds())/1000,
	)
}

func (p *pinger) send(target string) {
	start := time.Now()
	bytes := p.prepareMessage()

	addr := p.ipAddress(target)
	_, err := p.conn.WriteTo(bytes, &net.IPAddr{IP: addr})
	errs.HandleFatal(err, "Error sending packet")

	n, reply, peer := p.reply()

	duration := time.Since(start)

	parsed, err := icmp.ParseMessage(_const.ProtocolICMP, reply[:n])
	errs.HandleFatal(err, "Error parsing response")

	switch parsed.Type {
	case ipv4.ICMPTypeEchoReply:
		p.countSuccessReply++
		p.rttList = append(p.rttList, duration)
		fmt.Printf("Received reply from %v in %v\n", peer, duration)
	default:
		fmt.Printf("Unexpected ICMP message: %+v\n", parsed)
	}
}

func (p *pinger) prepareMessage() []byte {
	message := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{ID: os.Getegid() & 0xffff, Seq: 1, Data: []byte("PING")},
	}

	bytes, err := message.Marshal(nil)
	errs.HandleFatal(err, "Error marshalling message")

	return bytes
}

func (p *pinger) reply() (int, []byte, net.Addr) {
	reply := make([]byte, _const.MTUDefaultSize)
	err := p.conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	if err != nil {
		log.Println(fmt.Errorf("failed to read dedline: %v", err))
		return -1, nil, nil
	}

	n, peer, err := p.conn.ReadFrom(reply)
	errs.HandleFatal(err, "request timed out or error")

	return n, reply, peer

}

func (p *pinger) ipAddress(target string) net.IP {
	ip := net.ParseIP(target)
	if ip != nil {
		return ip
	} else {
		ipAddr := p.dnsLookUp(target)
		return ipAddr
	}
}

func (p *pinger) dnsLookUp(target string) net.IP {
	ips, err := net.LookupIP(target)
	if err != nil || len(ips) == 0 {
		log.Fatalf("DNS lookup failed for %s: %v", target, err)
	}

	for _, ip := range ips {
		if ip.To4() != nil {
			return ip
		}
	}

	log.Fatalf("no IPv4 address found for %s", target)
	return nil
}

func (p *pinger) initInternalData(retry int) {
	if retry <= 0 {
		p.retry = 4
	} else {
		p.retry = retry
	}
	p.countSuccessReply = _const.DefaultSuccessReply

	p.rttList = make([]time.Duration, retry)
}

func (p *pinger) Close() error {
	return p.conn.Close()
}
