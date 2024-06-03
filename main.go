package main

import (
	"fmt"
	"net"
)

func main() {
	var website = "facebook.com"
	iprecords, _ := net.LookupIP(website)
	cname, _ := net.LookupCNAME(website)
	nameserver, _ := net.LookupNS(website)
	for _, ip := range iprecords {
		fmt.Println(ip)
	}
        fmt.Println(cname)

	for _, name := range nameserver {
                fmt.Println(name)
	}
        mxrecords, _ := net.LookupMX(website)
	for _, mx := range mxrecords {
		fmt.Println(mx.Host, mx.Pref)
	}
         txtrecords, _ := net.LookupTXT(website)
 
	for _, txt := range txtrecords {
		fmt.Println(txt)
	}



}
