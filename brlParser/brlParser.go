package brlParser

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func RealToCents(input string) (int, error) {
	input = strings.Replace(input, ",", ".", 1)

	if s, err := strconv.ParseFloat(input, 32); err == nil {
		s = math.Round(s * 100)
		return int(s), nil
	}
	return 0, fmt.Errorf("invalid input")
}

func CentsToReal(value int) (string, error) {
	if value < 0 {
		return "", fmt.Errorf("invalid negative")
	}

	if value == 0 {
		return "R$0,00", nil
	}

	valueString := strconv.Itoa(value)

	if value < 10 {
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
