package main

import (
	"earlymoon/internal/dns"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <domain> <type1> [<type2> ...]", os.Args[0])
	}
	domain := os.Args[1]
	recordTypes := os.Args[2:]

	for _, recordType := range recordTypes {
		response, err := dns.Query(domain, recordType)
		if err != nil {
			log.Printf("Failed to query DNS for %s: %v", recordType, err)
			continue
		}
		fmt.Printf("Record Type: %s\n%s\n", recordType, response)
	}
}
