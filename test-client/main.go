package main

import (
	"log"
	"os"

	"github.com/recurly/recurly-go"
)

var (
	errLog = log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds)
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// INSERT YOUR API KEY AND SITE ID OR SUBDOMAIN HERE
	recurly.APIKey = ""
	// recurly.Site = recurly.SiteIdentifier{ID: 12345678}
	// or
	recurly.Site = recurly.SiteIdentifier{Subdomain: "SUBDOMAIN"}

	client := recurly.DefaultClient()

	account, err := client.GetAccountByAccountCode("robot-overlord")
	if err != nil {
		errLog.Printf("Failed to retrieve account: %v", err)
	} else {
		log.Printf("Received account: %v", account)
	}
}
