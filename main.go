package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Items struct {
	ProductName string
	UnitPrice   int
	Amount      int
}

type Emails []string

type List []Items

func (l List) SplitBill(emails Emails) map[string]int {
	billingList := make(map[string]int)
	numberOfPeople := len(emails)

	var sum int
	var avarage int

	for _, list := range l {
		sum += list.UnitPrice * list.Amount
	}

	var rest = sum % numberOfPeople
	var amountToSplit = sum - rest

	avarage = amountToSplit / numberOfPeople

	for _, email := range emails {
		billingList[email] = avarage
	}

	rand.Seed(time.Now().UnixNano())
	personToPayRest := emails[rand.Intn(numberOfPeople)]

	billingList[personToPayRest] += rest

	return billingList
}

func main() {
	fmt.Println()
	var barCheck = List{
		{"Cerveja", 11, 4},
		{"Petisco", 50, 1},
	}
	var emailList = Emails{
		"a@email.com",
		"b@email.com",
		"c@email.com",
	}
	fmt.Println(barCheck.SplitBill(emailList))
}
