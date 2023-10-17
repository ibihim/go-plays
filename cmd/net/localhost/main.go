package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if err := app(); err != nil {
		panic(err)
	}
}

func app() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("usage: %s <address>", os.Args[0])
	}

	cfg := getConfig()

	isLoopback, err := isLoopbackAddress(cfg.address)
	if err != nil {
		return fmt.Errorf("failed to check if address is loopback: %w", err)
	}

	fmt.Printf("isLoopback: %v\n", isLoopback)

	return nil
}

func getConfig() *config {
	return &config{
		address: os.Args[1],
	}
}

type config struct {
	address string
}

func isLoopbackAddress(address string) (bool, error) {
	ip := net.ParseIP(address)
	if ip != nil {
		return ip.IsLoopback(), nil
	}

	resolver := net.Resolver{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ips, err := resolver.LookupIPAddr(ctx, address)
	if err != nil {
		return false, fmt.Errorf("failed to lookup ip: %w", err)
	}

	for _, ip := range ips {
		if !ip.IP.IsLoopback() {
			return false, nil
		}
	}

	return true, nil
}
