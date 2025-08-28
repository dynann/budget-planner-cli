package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add         string
	Spend       string
	Info        string
	ListAccounts string
	ListRecords bool
	AllAccounts bool
	DeleteAccount string
	ChangeUsername string
	ChangeAccountNum string
	New 		string
}

func NewCmdFlag() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add balance to account account-number:balance ")
	flag.StringVar(&cf.Spend, "spend", "", "Add spending from your account account-number:balance")
	flag.StringVar(&cf.Info, "info", "", "print out the account info")
	flag.StringVar(&cf.ListAccounts, "list", "", "list all the records specify account-number ")
	flag.BoolVar(&cf.ListRecords, "record", false, "list all records")
	flag.BoolVar(&cf.AllAccounts, "acc", false, "list all available account")
	flag.StringVar(&cf.New, "new", "", "create new account by username:account-number:initail-balance")
	flag.StringVar(&cf.DeleteAccount, "del", "", "delete account by account number")
	flag.StringVar(&cf.ChangeAccountNum, "changeNum", "", "change account number old:new")
	flag.StringVar(&cf.ChangeUsername, "username", "", "change account holdername name:number")
	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute(records *Records, accounts *Accounts) {
	switch {
	case cf.ListAccounts != "":
		accounts.printAllAccount()
	case cf.Info != "":
		accounts.find(cf.Info)
	case cf.Add != "":
		parts := strings.SplitN(cf.Add, ":", 2)
		if len(parts) != 2 {
			fmt.Println("error, invalid format use account-number:balance")
			os.Exit(1)
		}

		acount_number := parts[0]
		balance, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("invalid balance value")
			os.Exit(1)
		}

		accounts.addBalance(acount_number, balance)
		// accounts.addBalance
	case cf.Spend != "":
		parts := strings.SplitN(cf.Spend, ":", 2)
		if len(parts) != 2 {
			fmt.Println("error, invalid format use account-number:balance")
			os.Exit(1)
		}

		account_number := parts[0]
		balance, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("invalid balance value")
		}
		accounts.spendBalance(account_number, balance)
	case cf.ListRecords:
		records.print()
	case cf.New != "":
		parts := strings.SplitN(cf.New, ":", 3)
		if len(parts) != 3 {
			fmt.Println("error, invalid format ")
		}
		account_number := parts[1]
		username := parts[0]
		initial_balance, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("invalid initial balance")
			return
		}
		accounts.createNewAccount(username, account_number, initial_balance)
	case cf.AllAccounts:
		accounts.printAllAccount()
	case cf.DeleteAccount != "":
		if err := accounts.deleteAccount(cf.DeleteAccount); err != nil {
			fmt.Println("account does not exist")
		}
	case cf.ChangeAccountNum != "":
		parts := strings.SplitN(cf.ChangeAccountNum, ":", 2)
		accounts.changeAccountNumber(parts[0], parts[1])
	case cf.ChangeUsername != "":
		parts := strings.SplitN(cf.ChangeUsername, ":", 2)
		accounts.changeAccountUsername(parts[0], parts[1] )
	default:
		fmt.Println("invalid command")
	}	
}