package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"netthrottle/pkg/limiter"
	"netthrottle/pkg/proxy"
)

func main() {
	listenAddr := flag.String("l", ":8080", "Listen address")
	remoteAddr := flag.String("r", "", "Remote destination (required)")
	latency := flag.Int("delay", 0, "Injected latency (ms)")
	bandwidth := flag.Int("kb", 1024, "Bandwidth limit (KB/s)")
	flag.Parse()

	if *remoteAddr == "" {
		log.Fatal("Error: Please specify remote address with -r")
	}

	rate := *bandwidth * 1024
	lb := limiter.NewTokenBucket(rate, rate)

	ln, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("🚀 NetThrottle active: %s -> %s\n", *listenAddr, *remoteAddr)
	fmt.Printf("⚙️  Config: %d KB/s | %d ms delay\n", *bandwidth, *latency)

	for {
		lConn, err := ln.Accept()
		if err != nil {
			continue
		}

		go func(local net.Conn) {
			defer local.Close()
			remote, err := net.Dial("tcp", *remoteAddr)
			if err != nil {
				return
			}
			defer remote.Close()

			done := make(chan bool)
			go func() {
				proxy.Transfer(local, remote, lb, *latency)
				done <- true
			}()
			go func() {
				proxy.Transfer(remote, local, lb, *latency)
				done <- true
			}()
			<-done
		}(lConn)
	}
}