package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jej10e5/Go_Cook/chapter5/dns"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <address>\n", os.Args[0])
		os.Exit(1)
	}
	address := os.Args[1]
	lookup, err := dns.LookupAddress(address)
	if err != nil {
		log.Panicf("failed to lookup: %s", err.Error())
	}
	fmt.Println(lookup)
}

//커밋
