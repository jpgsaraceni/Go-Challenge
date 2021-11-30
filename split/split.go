package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/jpgsaraceni/Go-Challenge/brlParser"
)

type Item struct {
	ProductName string
	UnitPrice   int
	Amount      int
}

type ItemString struct {
	ProductName string
	UnitPrice   string
	Amount      string
}

type EmailList []string

type ItemListString []ItemString

type ItemList []Item

var errRepeatedEmails = errors.New("lista contém emails repetidos")

func (i ItemList) SplitBill(emails EmailList) (map[string]int, error) {
	billingList := make(map[string]int)

	err := emails.checkForRepeated()
	if err != nil {
		return billingList, err
	}

	numberOfPeople := len(emails)
	sum := i.sumItems()
	rest := sum % numberOfPeople
	baseValueOwed := (sum - rest) / numberOfPeople

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(numberOfPeople, func(i, j int) {
		emails[i], emails[j] = emails[j], emails[i]
	})

	for i, email := range emails {
		billingList[email] = baseValueOwed

		if i < rest {
			billingList[email] += 1
		}
	}

	return billingList, nil
}

func (e EmailList) checkForRepeated() error {
	existentEmails := make(map[string]int)

	for _, email := range e {
		if existentEmails[email] > 0 {
			return errRepeatedEmails
		}

		existentEmails[email]++
	}
	return nil
}

func (i ItemList) sumItems() int {

	var sum int

	for _, item := range i {
		sum += item.UnitPrice * item.Amount
	}
	return sum
}

func main() {
	var barCheck = ItemListString{
		{"Cerveja", "11", "42"},
		{"Petisco", "50", "1"},
	}
	var emailList = EmailList{
		"a@email.com",
		"a@email.com",
		"c@email.com",
	}
	barCheckParsed := make(ItemList, 0)

	// var errInput = errors.New("invalid input")

	for _, item := range barCheck {
		priceInCents, err := brlParser.RealToCents(item.UnitPrice)

		if err != nil {
			fmt.Println(err)
			return
		}

		amountInt, err := strconv.Atoi(item.Amount)

		if err != nil {
			fmt.Println(err)
			return
		}

		barCheckParsed = append(barCheckParsed, Item{
			item.ProductName,
			priceInCents,
			amountInt,
		})
	}

	resultingMap, err := barCheckParsed.SplitBill(emailList)

	if err != nil {
		fmt.Println(err)
		return
	}
	for email, amount := range resultingMap {
		parsedValue, err := brlParser.CentsToReal(amount)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Email: %s. Valor a pagar: %s\n", email, parsedValue)
	}
}
