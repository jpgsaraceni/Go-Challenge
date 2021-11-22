package brlParser

import "testing"

var inputs = []struct {
	input string
	want  int
}{
	{"1,00", 100},
	{"1,01", 101},
	{"1,0", 100},
	{"1", 100},
	{"1.0", 100},
	{"1.00", 100},
}

func TestAtoi(t *testing.T) {
	for _, tt := range inputs {
		got, _ := RealToCents(tt.input)
		if got != tt.want {
			t.Errorf("got %d want %d", got, tt.want)
		}
	}
}

var outputs = []struct {
	input int
	want  string
}{
	{100, "R$1,00"},
	{110, "R$1,10"},
	{111, "R$1,11"},
}

func TestItoa(t *testing.T) {
	for _, tt := range outputs {
		got := CentsToReal(tt.input)
		if got != tt.want {
			t.Errorf("got %s want %s", got, tt.want)
		}
	}
}
