package pinger

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

type Ping struct {
	conn              *icmp.PacketConn
	addr              string
	retry             int
	countSuccessReply int
	rttList           []time.Duration
}

// send send ping
func (p *Ping) send() {
	start := time.Now()
	bytes := p.prepareMessage()

	addr := p.ipAddress()
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

// prepareMessage preparing message for ping
func (p *Ping) prepareMessage() []byte {
	message := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{ID: os.Getegid() & 0xffff, Seq: 1, Data: []byte("PING")},
	}

	bytes, err := message.Marshal(nil)
	errs.HandleFatal(err, "Error marshalling message")

	return bytes
}

// reply get reply from target addr
func (p *Ping) reply() (int, []byte, net.Addr) {
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

// ipAddress check if input is ip addr return ip else exec dnsLookUp and return first ip
func (p *Ping) ipAddress() net.IP {
	ip := net.ParseIP(p.addr)
	if ip != nil {
		return ip
	} else {
		ipAddr := dnsLookUp(p.addr)
		return ipAddr
	}
}

// dnsLookUp search ip V4 by domain
func dnsLookUp(target string) net.IP {
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

// initInternalData set internal retries and init rtt slice
func (p *Ping) initInternalData(retry int) {
	if retry <= 0 {
		p.retry = 4
	} else {
		p.retry = retry
	}
	p.countSuccessReply = _const.DefaultSuccessReply

	p.rttList = []time.Duration{}
}

func (p *Ping) Close() error {
	return p.conn.Close()
}
