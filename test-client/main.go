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
	recurly.SiteID = recurly.SiteIdentifier{Subdomain: "SUBDOMAIN"}

	client := recurly.DefaultClient()

	sites, err := client.ListSites(nil)
	if err != nil {
		errLog.Printf("Failed to retrieve sites: %v", err)
	} else {
		for i, site := range sites.Data {
			log.Printf("Site %3d: %d, %s, %s",
				i,
				site.ID.ToInt64(),
				site.Subdomain,
				site.Mode,
			)
		}
	}

	listParams := &recurly.AccountListParams{
		ListParams: recurly.ListParams{
			Sort:  recurly.SortCreatedAt,
			Order: recurly.ListAscending,
		},
	}
	accounts, err := client.ListAccounts(listParams)
	if err != nil {
		errLog.Printf("Failed to retrieve accounts: %v", err)
	} else {
		for i, account := range accounts.Data {
			log.Printf("Account %3d: %d, %s, %s",
				i,
				account.ID.ToInt64(),
				account.Code,
				*account.Email,
			)
		}
	}

	account, err := client.GetAccountByAccountCode("robot-overlord", nil)
	if err != nil {
		errLog.Printf("Failed to retrieve account: %v", err)
	} else {
		log.Printf("Received account: %v", account)
	}

	// --------------
	// Create account

	accountParams := &recurly.CreateAccountAttributes{
		Code:  "woody-123456",
		Email: "test@example.com",
	}
	newAccount, err := client.CreateAccount(accountParams)
	if err != nil {
		errLog.Printf("Failed to retrieve account: %v", err)
	} else {
		log.Printf("Received account: %v", newAccount)
	}

	// --------------
	// Update account

	accountUpdateParams := &recurly.UpdateAccountAttributes{
		Email:     "test-1234@example.com",
		FirstName: "Woody",
		LastName:  "Bird",
	}
	updatedAccount, err := client.UpdateAccount(newAccount.ID, accountUpdateParams)
	if err != nil {
		errLog.Printf("Failed to update account: %v", err)
	} else {
		log.Printf("Received updated account: %v", updatedAccount)
	}
}
