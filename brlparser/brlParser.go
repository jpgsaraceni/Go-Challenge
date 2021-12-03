package brlparser

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	errInputPrice    = errors.New("preços contêm caracteres inválidos")
	errNegativeValue = errors.New("valor negativo")
)

func RealToCents(input string) (int, error) {
	input = strings.Replace(input, ",", ".", 1)

	const floatBitSize int = 32

	if s, err := strconv.ParseFloat(input, floatBitSize); err == nil {
		s = math.Round(s * 100)

		return int(s), nil
	}

	return 0, errInputPrice
}

func CentsToReal(value int) (string, error) {
	if value < 0 {
		return "", errNegativeValue
	}

	if value == 0 {
		return "R$0,00", nil
	}

	valueString := strconv.Itoa(value)

	if digitLimit := 10; value < digitLimit {
		return "R$0,0" + valueString, nil
	}

	wholePrefix := ""
	if value < 100 {
		wholePrefix = "0"
	}
	wholePart := valueString[:len(valueString)-2]
	decimalPart := valueString[len(valueString)-2:]
	valueString = fmt.Sprintf("R$%s%s,%s", wholePrefix, wholePart, decimalPart)

	return valueString, nil
}
