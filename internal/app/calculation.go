package app

import "time"

func (p *pinger) calcPacketLost() float32 {
	if p.retry == 0 {
		return 0
	}
	lost := float32(p.retry - p.countSuccessReply)
	return (lost / float32(p.retry)) * 100
}

func (p *pinger) calcAvgTime() time.Duration {
	if len(p.rttList) == 0 {
		return 0
	}
	var sum time.Duration
	for _, rtt := range p.rttList {
		sum += rtt
	}
	return sum / time.Duration(len(p.rttList))
}

func (p *pinger) calcMinAndMaxTime() (min time.Duration, max time.Duration) {
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
