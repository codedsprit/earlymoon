package main

import (
        "flag"
        "fmt"
)

func main() {
	// Parse the flags
	flag.Parse()

	// Get the remaining arguments after flag parsing
	args := flag.Args()

	// Check if at least two arguments are provided
	if len(args) <  {
		fmt.Println("Error: Please provide at least one domain and one record type.")
		flag.Usage()
		return
	}

	// The first argument can be either a domain or a record type
	firstArg := args[0]
	secondArg := args[1]

	// Determine if the first argument is a domain (heuristically)
	// This is a simplistic heuristic; in a real-world scenario, you might want to improve it
	isDomain := func(arg string) bool {
		return len(arg) > 0 && !containsOnlyLetters(arg)
	}

	// Determine domain and record types
	var domain string
	var recordTypes []string

	if isDomain(firstArg) {
		domain = firstArg
		recordTypes = args[1:]
	} else if isDomain(secondArg) {
		domain = secondArg
		recordTypes = append(recordTypes, firstArg)
		recordTypes = append(recordTypes, args[2:]...)
	} else {
		fmt.Println("Error: Unable to determine the domain. Please provide a valid domain.")
		flag.Usage()
		return
	}

	// Print the values
	fmt.Printf("Domain: %s\n", domain)
	fmt.Printf("Record Types: %v\n", recordTypes)

	// Here you can add the logic to query the DNS records based on the provided domain and record types
}

// containsOnlyLetters checks if a string contains only letters (simplistic heuristic for record types)
func containsOnlyLetters(s string) bool {
	for _, r := range s {
		if r < 'A' || r > 'Z' {
			return false
		}
	}
	return true
}

