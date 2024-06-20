package main

import (
	"earlymoon/internal/dns"
	"flag"
	"fmt"
	"log"
	"os"
)

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "  %s -d <domain> -t <type>\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Options:\n")
	fmt.Fprintf(os.Stderr, "  -d, -domain <domain>   The domain to query (required)\n")
	fmt.Fprintf(os.Stderr, "  -t, -type <type>       The type of DNS record to query (required)\n")
	fmt.Fprintf(os.Stderr, "  -h                     Show this help message\n")
	fmt.Fprintf(os.Stderr, "\nExample:\n")
	fmt.Fprintf(os.Stderr, "  %s -d example.com -t A\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\n")
}

func main() {
	var domain string
	var recordType string

	flag.StringVar(&domain, "domain", "", "The domain to query (required)")
	flag.StringVar(&domain, "d", "", "The domain to query (required)")
	flag.StringVar(&recordType, "type", "", "The type of DNS record to query (required)")
	flag.StringVar(&recordType, "t", "", "The type of DNS record to query (required)")

	flag.Usage = printUsage
	flag.Parse()

	if domain == "" || recordType == "" {
		flag.Usage()
		log.Fatalf("Both domain and type are required")
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
}
