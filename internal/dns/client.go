package dns

import (
	"fmt"
	"github.com/miekg/dns"
)

// Query performs a DNS query for the specified domain and record type.
func Query(domain, recordType string) (string, error) {
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), dns.StringToType[recordType])
	m.RecursionDesired = true

	c := new(dns.Client)
	in, _, err := c.Exchange(m, "8.8.8.8:53")
	if err != nil {
		return "", err
	}
	if len(in.Answer) == 0 {
		return "", fmt.Errorf("no records found for %s", domain)
	}

	response := formatResponse(in.Answer)
	return response, nil
}
