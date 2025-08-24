/*
simple script to scan for open ports within a provided range.
it has no concurrency and it could be better written, but it
does the job.
*/

package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func printUsage() {
	fmt.Println("usage: sganner <host> <start> <end>")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 4 {
		printUsage()
	}

	host := os.Args[1]
	start, startErr := strconv.Atoi(os.Args[2])
	end, endErr := strconv.Atoi(os.Args[3])

	if startErr != nil || endErr != nil || start <= 0 || end <= 0 || end < start {
		printUsage()
	}

	for port := start; port <= end; port++ {
		addr := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", addr, 1*time.Second)

		if err == nil {
			fmt.Printf("port %d is open\n", port)
			conn.Close()
		}
	}
}
