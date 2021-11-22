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

func CentsToReal(value int) string {
	valueString := strconv.Itoa(value)
	wholePart := valueString[:len(valueString)-2]
	decimalPart := valueString[len(valueString)-2:]
	valueString = fmt.Sprintf("R$%s,%s", wholePart, decimalPart)
	return valueString
}
