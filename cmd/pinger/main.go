package main

import (
	"flag"
	"go-pinger/internal/errs"
	"go-pinger/internal/pinger"
	"os"
	"strings"
)

func main() {

	var retry int
	var filePath string
	var target string

	flag.StringVar(&target, "t", "", "target address")
	flag.IntVar(&retry, "r", 0, "count of ping (default 4)")
	flag.StringVar(&filePath, "f", "", "path to file with ip`s or domains")
	flag.Parse()

	pinger := pinger.New()
	defer pinger.Close()

	if filePath != "" {
		bytes, err := os.ReadFile(filePath)
		errs.HandleFatal(err, "can not read file, maybe not such file or directory")

		targets := parseFile(bytes)
		pinger.SendFewPings(targets, retry)
	} else {
		pinger.SendPing(target, retry)
	}

}

func parseFile(bytes []byte) []string {
	lines := strings.Split(string(bytes), "\n")
	var result []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		result = append(result, line)
	}

	return result
}
