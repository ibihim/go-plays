package main

import (
	"log"
	"net"
	"strconv"
	"strings"
)

var subnetToScan = "192.168.178"

func main() {
	activeThreads := 0
	doneChannel := make(chan bool)

	for ip := 0; ip <= 255; ip++ {
		fullIP := subnetToScan + "." + strconv.Itoa(ip)
		go resolve(fullIP, doneChannel)
		activeThreads++
	}

	// Wait for all threads to finish
	for activeThreads > 0 {
		<-doneChannel
		activeThreads--
	}
}

func resolve(ip string, doneChannel chan bool) {
	addresses, err := net.LookupAddr(ip)
	if err == nil {
		log.Printf("%s - %s\n", ip, strings.Join(addresses, ", "))
	}

	doneChannel <- true
}
