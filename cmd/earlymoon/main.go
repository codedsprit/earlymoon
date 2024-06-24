package main

import (
	"earlymoon/internal/dns"
        "earlymoon/internal/help"
	"flag"
	"fmt"
	"log"
)

func main() {
	var domain string
	var recordType string

	flag.StringVar(&domain, "domain", "", "The domain to query (required)")
	flag.StringVar(&domain, "d", "", "The domain to query (required)")
	flag.StringVar(&recordType, "type", "", "The type of DNS record to query (required)")
	flag.StringVar(&recordType, "t", "", "The type of DNS record to query (required)")

	flag.Parse()

	if domain == "" || recordType == "" {
		flag.Usage()
		log.Fatalf("Both domain and type are required")
                help.PrintBanner()
	}

	response, err := dns.Query(domain, recordType)
	if err != nil {
		log.Fatalf("Failed to query DNS: %v", err)
	}

	if len(response) == 0 {
		fmt.Printf("No %s records found for %s\n", recordType, domain)
		return
	}

	fmt.Println(response)

                help.PrintBanner()
}
