package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain. hasMX,hasSPF,sprRecord,hasDAMARC,dmarchRecord\n")
	for scanner.Scan(){
		checkDoman(scanner.Text())
	}

	if err := scanner.Err(); err != nil{
		log.Fatal("Error: could not read from the input: %v\n",err)
	}
}

func checkDoman(domain string){
	var hasMX,hasSPF,hasDAMARC bool
	var spfRecord, dmarchRecord string

	mxRecords,err := net.LookupMX(domain)

	if err != nil{
		log.Printf("Error: %v\n",err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n",err)
	}

	for _, recrod := range txtRecords{
		if strings.HasPrefix(recrod,"v=spf1"){
			hasSPF = true
			spfRecord = recrod
			break
		}
	}

	dmarchRecords,err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error %v" , err)
	}
	for _,record := range dmarchRecords{
		if strings.HasPrefix(record,"v=DMARC1"){
			hasDAMARC = true
			dmarchRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v",domain,hasMX,hasSPF,spfRecord , hasDAMARC, dmarchRecord)
}