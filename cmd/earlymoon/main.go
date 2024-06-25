package main

import (
	"earlymoon/internal/dns"
	"earlymoon/internal/help"
	"log"
        "fmt"
)

func main() {
	// Parse arguments
	args := help.ParseArgs()
	
	// Handle arguments
	help.HandleArgs(args)

	// Query DNS
	response, err := dns.Query(args.Domain, args.RecordType)
	if err != nil {
		log.Fatalf("Failed to query DNS: %v", err)
	}

	if len(response) == 0 {
		log.Printf("Maybe wrong args parsed, try help")
		return
	}

	// Print response
	fmt.Println(response)
}

