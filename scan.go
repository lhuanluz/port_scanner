package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func scanPort(host string, port int, wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, strconv.Itoa(port)), time.Second)
	if err == nil {
		fmt.Printf("Port %d is open\n", port)
		conn.Close()
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: portscanner <host>")
		os.Exit(1)
	}

	host := os.Args[1]
	startPort := 1
	endPort := 1024

	var wg sync.WaitGroup
	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go scanPort(host, port, &wg)
	}
	wg.Wait()
}
