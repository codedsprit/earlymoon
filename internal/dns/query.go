package dns

import (
	"fmt"
	"github.com/miekg/dns"
)

func formatResponse(answers []dns.RR) string {
	var response string
	for _, ans := range answers {
		response += fmt.Sprintf("%v\n", ans)
	}
	return response
}
