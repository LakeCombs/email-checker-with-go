package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDNS \n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func checkDomain(domain string) {

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error:  %v\v", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	spfRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range spfRecords {
		if strings.Contains(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.Contains(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v \n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

}
