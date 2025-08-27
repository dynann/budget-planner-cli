package main

import (
	"errors"
	// "fmt"
	"os"
	"strconv"

	"github.com/aquasecurity/table"
)

type Account struct {
	User    string
	AccountNo string
	Balance int
}


type Accounts []Account

var record = Records{}

func NewAccount(user string, balance int) *Account {
	return &Account{
		User: user,
		Balance: balance,
	}
}

// func (account *Account) addBalance(balance int) error {
// 	if balance > 0 {
// 		account.Balance += balance
// 		record.save(account, balance)
// 		return nil
// 	}

// 	return errors.New("cannot add zero or negative balance")
// }

// func (account *Account) spendBalance(balance int) error {
// 	if balance > 0 {
// 		account.Balance -= balance
// 		record.save(account, -balance)
// 		return nil
// 	}
// 	return errors.New("cannot spend zero or negative balance")
// }

// func (account *Account) info() {
// 	table := table.New(os.Stdout)
// 	table.SetRowLines(false)
// 	fmt.Println("hello world")
// 	table.SetHeaders("Account Holder Name", "Available Balance")
// 	table.AddRow(account.User, strconv.Itoa(account.Balance))
// 	table.Render()
// }


func (accounts *Accounts) findAccount(accountNum string) *int{
	for index, account := range *accounts {
		if accountNum == account.AccountNo {
			return &index
		}
	}
	return nil
}


// -> find account via number
func (accounts *Accounts) find(accountNum string) {
	for _, account := range *accounts {
		if account.AccountNo == accountNum {
			accounts.showInfo(account)
		}
	}
}

// -> find account 
func (accounts *Accounts) showInfo(account Account) {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("Account Holder Name", "Account Number", "Available Balance")
	table.AddRow(account.User, account.AccountNo, strconv.Itoa(account.Balance))
	table.Render()
}

func (accounts *Accounts) createNewAccount(accountHolderName string, accountNum string, balance int) {
	account := Account{
		User: accountHolderName,
		AccountNo: accountNum,
		Balance: balance,
	}
	*accounts = append(*accounts, account)
}

func (accounts *Accounts) addBalance(accountNum string, balance int) error {
	if err := accounts.findAccount(accountNum); err != nil {
		(*accounts)[*err].Balance += balance
		record.save(&(*accounts)[*err], balance)
		return nil
	}
	return errors.New("error account does not exist")
}

func (accounts *Accounts) spendBalance(accountNum string, balance int) error{
	if err := accounts.findAccount(accountNum); err != nil {
		(*accounts)[*err].Balance -= balance
		record.save(&(*accounts)[*err], -balance)
		return nil
	}
	return errors.New("error account does not exist")
}

func (accounts *Accounts) printAllAccount() {
	table := table.New(os.Stdout)
	table.SetRowLines(true)
	table.SetHeaders("#", "Account Holder Name", "Account Number", "Available Balance")
	for index, a := range *accounts {
		name := a.User
		number := a.AccountNo
		balance := strconv.Itoa(a.Balance)
		table.AddRow(strconv.Itoa(index), name, number, balance)
	}

	table.Render()
}


