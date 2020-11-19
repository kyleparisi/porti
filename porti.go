package main

import (
	"fmt"
	"net"
	"os"
)
import "time"
import "flag"

func recurse(address string) {
	_, err := net.Dial("tcp", address)
	if err != nil {
		time.Sleep(600)
		recurse(address)
	}
}

func worker(address string, results chan<- string) {
	recurse(address)
	results <- address
}

func main() {
	helpPtr := flag.Bool("help", false, "usage help")
	hPtr := flag.Bool("h", false, "usage help")

	flag.Parse()
	if *helpPtr == true || *hPtr == true {
		fmt.Println("usage: porti localhost:3000")
		os.Exit(0)
	}

	addresses := flag.Args()
	if len(addresses) == 0 {
		fmt.Println("No addresses provided.")
		os.Exit(0)
	}
	results := make(chan string, len(addresses))
	for _, address := range addresses {
		go worker(address, results)
	}
	for range addresses {
		<-results
	}
}
