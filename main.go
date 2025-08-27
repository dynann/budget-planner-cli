package main

import (
	"fmt"
)

func main() {
	// account := NewAccount("dynann", 100)
	storageRecord := NewStorage[Records]("records.json")
	// storageAccount := NewStorage[Account]("account.json")
	storageAccounts := NewStorage[Accounts]("accounts.json")

	// account := &Account{}
	// if err := storageAccount.load(account); err != nil {
	// 	account = NewAccount("dynann", 100)
	// }

	accounts := Accounts{}
	if err := storageAccounts.load(&accounts); err != nil {
		fmt.Println("please create a new account")
	}

	// storageAccount.load(account)
	storageRecord.load(&record)
	command := NewCmdFlag()
	command.Execute(&record, &accounts)
	// storageAccount.save(*account)
	storageRecord.save(record)
	storageAccounts.save(accounts)
}