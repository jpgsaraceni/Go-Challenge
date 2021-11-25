package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/jpgsaraceni/Go-Challenge/brlParser"
)

type Item struct {
	ProductName string
	UnitPrice   string
	Amount      int
}

type EmailList []string

type ItemList []Item

func (i ItemList) SplitBill(emails EmailList) (map[string]int, error) {

	billingList := make(map[string]int)
	numberOfPeople := len(emails)

	sum, err := i.sumItems()

	if err != nil {
		return billingList, err
	}

	valuesOwed := i.getValuesOwed(sum, numberOfPeople)
	billingList = i.distributeValuesOwed(emails, valuesOwed)

	return billingList, nil
}

var errInput = errors.New("invalid input")

func (i ItemList) sumItems() (int, error) {

	var sum int

	for _, item := range i {
		unitPrice, err := brlParser.RealToCents(item.UnitPrice)

		if err != nil {
			return 0, errInput
		}

		sum += unitPrice * item.Amount
	}
	return sum, nil
}

func (i ItemList) getValuesOwed(sum, numberOfPeople int) []int {
	var valuesOwed = make([]int, numberOfPeople)
	var rest int = sum % numberOfPeople

	for i := 0; i < numberOfPeople; i++ {
		valuesOwed[i] = (sum - rest) / numberOfPeople

		if i < rest {
			valuesOwed[i] += 1
		}
	}

	return valuesOwed
}

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
	var barCheck = ItemList{
		{"Cerveja", "11", 42},
		{"Petisco", "50", 1},
	}
	var emailList = EmailList{
		"a@email.com",
		"b@email.com",
		"c@email.com",
	}
	resultingMap, err := barCheck.SplitBill(emailList)
	if err != nil {
		fmt.Println("verifique os valores informados e tente novamente!")
	}
	for email, amount := range resultingMap {
		parsedValue := brlParser.CentsToReal(amount)
		fmt.Printf("Email: %s. Valor a pagar: %s\n", email, parsedValue)
	}
}
