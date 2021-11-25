package main

import (
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

func (i ItemList) SplitBill(emails EmailList) (map[string]int, error) {

	numberOfPeople := len(emails)

	sum := i.sumItems()

	valuesOwed := i.getValuesOwed(sum, numberOfPeople)
	billingList := i.distributeValuesOwed(emails, valuesOwed)
	// billingList := i.distributeValues(emails, sum, numberOfPeople)

	return billingList, nil
}

func (i ItemList) sumItems() int {

	var sum int

	for _, item := range i {
		sum += item.UnitPrice * item.Amount
	}
	return sum
}

func (i ItemList) getValuesOwed(sum, numberOfPeople int) []int {
	var valuesOwed = make([]int, numberOfPeople)
	var rest int = sum % numberOfPeople
	var baseValueOwed = (sum - rest) / numberOfPeople

	for i := 0; i < numberOfPeople; i++ {
		valuesOwed[i] = baseValueOwed

		if i < rest {
			valuesOwed[i] += 1
		}
	}

	return valuesOwed
}

// don't know how to shuffle values in the same loop as they are assigned (if possible)

// func (i ItemList) distributeValues(emails EmailList, sum, numberOfPeople int) map[string]int {
// 	billingListValues := make(map[string]int)

// 	var valuesOwed = make([]int, numberOfPeople)
// 	var rest int = sum % numberOfPeople
// 	var baseValueOwed = (sum - rest) / numberOfPeople

// 	for i, email := range emails {
// 		valuesOwed[i] = baseValueOwed

// 		if i < rest {
// 			valuesOwed[i] += 1
// 		}

// 		rand.Seed(time.Now().UnixNano())
// 		rand.Shuffle(len(valuesOwed), func(i, j int) {
// 			valuesOwed[i], valuesOwed[j] = valuesOwed[j], valuesOwed[i]
// 		})

// 		billingListValues[email] = valuesOwed[i]
// 	}

// 	return billingListValues
// }

func (i ItemList) distributeValuesOwed(emails EmailList, valuesOwed []int) map[string]int {
	billingListValues := make(map[string]int)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(valuesOwed), func(i, j int) {
		valuesOwed[i], valuesOwed[j] = valuesOwed[j], valuesOwed[i]
	})

	for i, email := range emails {
		billingListValues[email] = valuesOwed[i]
	}
	return billingListValues
}

func main() {
	var barCheck = ItemListString{
		{"Cerveja", "11", "42"},
		{"Petisco", "50", "1"},
	}
	var emailList = EmailList{
		"a@email.com",
		"b@email.com",
		"c@email.com",
	}
	barCheckParsed := make(ItemList, 0)

	// var errInput = errors.New("invalid input")

	for _, item := range barCheck {
		priceInCents, err := brlParser.RealToCents(item.UnitPrice)

		if err != nil {
			fmt.Println("verifique os valores informados e tente novamente!")
			return
		}

		amountInt, err := strconv.Atoi(item.Amount)

		if err != nil {
			fmt.Println("verifique os valores informados e tente novamente!")
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
		fmt.Println("verifique os valores informados e tente novamente!")
	}
	for email, amount := range resultingMap {
		parsedValue := brlParser.CentsToReal(amount)
		fmt.Printf("Email: %s. Valor a pagar: %s\n", email, parsedValue)
	}
}
