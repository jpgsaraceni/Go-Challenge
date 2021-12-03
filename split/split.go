package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/jpgsaraceni/Go-Challenge/brlparser"
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

type (
	EmailList      []string
	ItemListString []ItemString
	ItemList       []Item
)

type Shuffler interface {
	Shuffle(EmailList) EmailList
}

type DefaultShuffle struct{}

var ErrRepeatedEmails = errors.New("lista contém emails repetidos")

func (i ItemList) SplitBill(emails EmailList, shuffler Shuffler) (map[string]int, error) {
	billingList := make(map[string]int)

	numberOfPeople := len(emails)
	sum := i.sumItems()
	rest := sum % numberOfPeople
	baseValueOwed := (sum - rest) / numberOfPeople

	emails = shuffler.Shuffle(emails)

	for i, email := range emails {
		billingList[email] = baseValueOwed

		if i < rest {
			billingList[email]++
		}
	}

	return billingList, nil
}

func (DefaultShuffle) Shuffle(emails EmailList) EmailList {
	rand.Shuffle(len(emails), func(i, j int) {
		emails[i], emails[j] = emails[j], emails[i]
	})

	return emails
}

func (e EmailList) checkForRepeated() error {
	existentEmails := make(map[string]int)

	for _, email := range e {
		if existentEmails[email] > 0 {
			return ErrRepeatedEmails
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
	rand.Seed(time.Now().UnixNano())

	barCheck := ItemListString{
		{"Cerveja", "11", "42"},
		{"Petisco", "50", "1"},
	}
	emailList := EmailList{
		"a@email.com",
		"b@email.com",
		"c@email.com",
	}

	err := emailList.checkForRepeated()
	if err != nil {
		fmt.Println(err)

		return
	}

	barCheckParsed := make(ItemList, 0)

	for _, item := range barCheck {
		var priceInCents int

		priceInCents, err = brlparser.RealToCents(item.UnitPrice)
		if err != nil {
			fmt.Println(err)

			return
		}

		var amountInt int
		amountInt, err = strconv.Atoi(item.Amount)
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

	d := DefaultShuffle{}
	resultingMap, err := barCheckParsed.SplitBill(emailList, d)
	if err != nil {
		fmt.Println(err)

		return
	}
	for email, amount := range resultingMap {
		parsedValue, err := brlparser.CentsToReal(amount)
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Printf("Email: %s. Valor a pagar: %s\n", email, parsedValue)
	}
}
