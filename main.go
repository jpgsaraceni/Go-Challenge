package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Items struct {
	ProductName string
	UnitPrice   string
	Amount      int
}

type Emails []string

type List []Items

func (l List) SplitBill(emails Emails) (map[string]string, error) {
	billingListValues := make(map[string]int)
	billingListStrings := make(map[string]string)
	numberOfPeople := len(emails)

	var sum int
	var avarage int

	for _, item := range l {
		unitPrice, err := parseInput(item.UnitPrice)

		if err != nil {
			return billingListStrings, fmt.Errorf("invalid input")
		}

		sum += unitPrice * item.Amount
	}

	var rest = sum % numberOfPeople
	var amountToSplit = sum - rest

	avarage = amountToSplit / numberOfPeople

	for _, email := range emails {
		billingListValues[email] = avarage
	}

	rand.Seed(time.Now().UnixNano())
	personToPayRest := emails[rand.Intn(numberOfPeople)]

	billingListValues[personToPayRest] += rest

	for email, amount := range billingListValues {
		billingListStrings[email] = parseOutput(amount)
	}

	return billingListStrings, nil
}

func parseInput(input string) (int, error) {
	input = strings.Replace(input, ",", ".", 1)

	if s, err := strconv.ParseFloat(input, 32); err == nil {
		s = math.Round(s * 100)
		return int(s), nil
	}
	return 0, fmt.Errorf("invalid input.")
}

func parseOutput(value int) string {
	valueString := strconv.Itoa(value)
	integer := valueString[:len(valueString)-2]
	decimal := valueString[len(valueString)-2:]
	valueString = fmt.Sprintf("R$%s,%s", integer, decimal)
	return valueString
}

func main() {
	var barCheck = List{
		{"Cerveja", "11", 4},
		{"Petisco", "50", 1},
	}
	var emailList = Emails{
		"a@email.com",
		"b@email.com",
		"c@email.com",
	}
	resultingMap, err := barCheck.SplitBill(emailList)
	if err != nil {
		fmt.Println("verifique os valores informados e tente novamente!")
	}
	for email, amount := range resultingMap {
		fmt.Printf("Email: %s. Valor a pagar: %s\n", email, amount)
	}
}
