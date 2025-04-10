package pinger

import (
	"fmt"
	"time"
)

func (p *Ping) calcPacketLost() float32 {
	if p.retry == 0 {
		return 0
	}
	lost := float32(p.retry - p.countSuccessReply)
	return (lost / float32(p.retry)) * 100
}

func (p *Ping) calcAvgTime() time.Duration {
	if len(p.rttList) == 0 {
		return 0
	}
	var sum time.Duration
	for _, rtt := range p.rttList {
		sum += rtt
	}
	return sum / time.Duration(len(p.rttList))
}

func (p *Ping) calcMinAndMaxTime() (min time.Duration, max time.Duration) {
	if len(p.rttList) == 0 {
		return 0, 0
	}

	min, max = p.rttList[0], p.rttList[0]

	for _, rtt := range p.rttList[1:] {
		if rtt < min {
			min = rtt
		}
		if rtt > max {
			max = rtt
		}
	}
	return min, max
}

func (p *Ping) showInfo() {
	minTime, maxTime := p.calcMinAndMaxTime()

	fmt.Printf("\n══════════════════════════════════════════════════════════════════════\n")
	fmt.Printf("🏓  --- Ping statistics for %s ---  🏓\n", p.addr)
	fmt.Printf("══════════════════════════════════════════════════════════════════════\n")
	fmt.Printf("  ▪ %d packets transmitted, %d received, %.2f%% packet loss\n",
		p.retry,
		p.countSuccessReply,
		p.calcPacketLost(),
	)
	fmt.Printf("  ▪ Round-Trip Time (RTT):\n")
	fmt.Printf("    • Min RTT:  %.2f ms\n", float64(minTime.Microseconds())/1000)
	fmt.Printf("    • Avg RTT:  %.2f ms\n", float64(p.calcAvgTime().Microseconds())/1000)
	fmt.Printf("    • Max RTT:  %.2f ms\n", float64(maxTime.Microseconds())/1000)
	fmt.Printf("══════════════════════════════════════════════════════════════════════\n")
}
