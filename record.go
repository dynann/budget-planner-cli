package main

import (
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Record struct {
	Amount  int
	Account *Account
	Date    time.Time
}

type Records []Record

func (records *Records) save(account *Account, amount int) {
	record := Record{
		Amount: amount,
		Account: account,
		Date: time.Now(),
	}
	*records = append(*records, record)
}

func (records *Records) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(true)
	table.SetHeaders("#", "Account", "Amount", "Date")
	// fmt.Print(*records)
	for index, t := range *records {
		name := t.Account.User
		amount := t.Amount
		date := t.Date.Format(time.RFC1123)
		table.AddRow(strconv.Itoa(index), name, strconv.Itoa(amount), date)
	}
	table.Render()
}