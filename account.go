package main

import (
	"errors"
	"fmt"
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

func (accounts *Accounts) createNewAccount(accountHolderName string, accountNum string, balance int) error {
	if err := accounts.findAccount(accountNum); err == nil {
		account := Account{
			User: accountHolderName,
			AccountNo: accountNum,
			Balance: balance,
		}
		*accounts = append(*accounts, account)
	}
	return errors.New("this account already exists")
}

func (accounts *Accounts) addBalance(accountNum string, balance int) error {
	if err := accounts.findAccount(accountNum); err != nil {
		(*accounts)[*err].Balance += balance
		record.save(&(*accounts)[*err], balance)
		return nil
	}
	return errors.New("error account does not exist")
}

func (accounts *Accounts) spendBalance(accountNum string, balance int) error {
	if err := accounts.findAccount(accountNum); err != nil {
		if balance > (*accounts)[*err].Balance {
			fmt.Println("you don't have enough money")
			return errors.New("you don't have enough balance")
		}
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

func (accounts *Accounts) deleteAccount(accountNum string) error{
	if err := accounts.findAccount(accountNum); err != nil {
		a := *accounts
		*accounts = append(a[:*err], a[*err+1:]...) //start from 0 to before index, and skip specify index to last index
		record.deleteRecords(accountNum)
		return nil
	}
	return errors.New("account does not exist")
}

func (accounts *Accounts) changeAccountUsername(username string, accountNum string) error{
	acc := *accounts
	err := accounts.findAccount(accountNum)
	if err == nil {
		fmt.Println("cannot change username")
		return errors.New("cannot change username")
	}
	acc[*err].User = username
	return nil
}

func (accounts *Accounts) changeAccountNumber(accountNumOld string, accountNumNew string) error {
	acc := *accounts
	errNew := accounts.findAccount(accountNumNew); if errNew != nil {
		fmt.Println("account already exists")
		return nil
	}

	errOld := accounts.findAccount(accountNumOld); if errOld != nil {
		acc[*errOld].AccountNo = accountNumNew
	}
	return nil
}




