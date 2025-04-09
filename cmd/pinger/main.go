package main

import (
	"flag"
	"go-pinger/internal/app"
)

func main() {
	var target string
	var retry int

	flag.StringVar(&target, "t", "", "target address")
	flag.IntVar(&retry, "r", 0, "count of ping (default 4)")
	flag.Parse()

	pinger := app.New()
	defer pinger.Close()
	pinger.SendPing(target, retry)
}
