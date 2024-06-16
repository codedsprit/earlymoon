package main

import (
	"earlymoon/internal/dns"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <domain> [type]", os.Args[0])
	}
	domain := os.Args[1]
	recordType := "A"
	if len(os.Args) > 2 {
		recordType = os.Args[2]
	}

	response, err := dns.Query(domain, recordType)
	if err != nil {
		log.Fatalf("Failed to query DNS: %v", err)
	}

	fmt.Println(response)
}
