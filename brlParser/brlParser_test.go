package brlParser

import (
	"fmt"
	"testing"
)

func TestRealToCents(t *testing.T) {
	type testCase struct {
		name            string
		input           string
		expectedSuccess int
		expectedError   error
	}

	testCases := []testCase{
		{
			name:            "should return value in cents, receives comma separated 2 decimal places",
			input:           "1,01",
			expectedSuccess: 101,
			expectedError:   nil,
		},
		{
			name:            "should return value in cents, receives comma separated 2 decimal places",
			input:           "100,00",
			expectedSuccess: 10000,
			expectedError:   nil,
		},
		{
			name:            "should return value in cents, receives comma separated 2 decimal places",
			input:           "23,13",
			expectedSuccess: 2313,
			expectedError:   nil,
		},
		{
			name:            "should return value in cents, receives comma separated 1 decimal place",
			input:           "10,0",
			expectedSuccess: 1000,
			expectedError:   nil,
		},
		{
			name:            "should return value in cents, receives comma separated 1 decimal place",
			input:           "23,4",
			expectedSuccess: 2340,
			expectedError:   nil,
		},
		{
			name:            "should return value in cents, receives integer",
			input:           "2",
			expectedSuccess: 200,
			expectedError:   nil,
		},
		{
			name:            "should return value in cents, receives dot separated 2 decimal places",
			input:           "74.01",
			expectedSuccess: 7401,
			expectedError:   nil,
		},
		{
			name:            "should return value in cents, receives dot separated 2 decimal places",
			input:           "100.00",
			expectedSuccess: 10000,
			expectedError:   nil,
		},
		{
			name:            "should return value in cents, receives dot separated 2 decimal places",
			input:           "123.13",
			expectedSuccess: 12313,
			expectedError:   nil,
		},
		{
			name:            "should return value in cents, receives dot separated 1 decimal place",
			input:           "10.0",
			expectedSuccess: 1000,
			expectedError:   nil,
		},
		{
			name:            "should return value in cents, receives dot separated 1 decimal place",
			input:           "90.3",
			expectedSuccess: 9030,
			expectedError:   nil,
		},
		{
			name:            "should return value in cents, rounding exceeding decimal places",
			input:           "1.036",
			expectedSuccess: 104,
			expectedError:   nil,
		},
		{
			name:            "should return value in cents, rounding exceeding decimal places",
			input:           "1.034",
			expectedSuccess: 103,
			expectedError:   nil,
		},
		{
			name:            "should return value in cents, rounding exceeding decimal places",
			input:           "20,129",
			expectedSuccess: 2013,
			expectedError:   nil,
		},
		{
			name:            "should return value in cents, rounding exceeding decimal places",
			input:           "3,010",
			expectedSuccess: 301,
			expectedError:   nil,
		},
		{
			name:            "should return error when not a number.",
			input:           "3,a010",
			expectedSuccess: 0,
			expectedError:   fmt.Errorf("invalid input"),
		},
		{
			name:            "should return error when not a number.",
			input:           "abc",
			expectedSuccess: 0,
			expectedError:   fmt.Errorf("invalid input"),
		},
		{
			name:            "should return error when more than one comma.",
			input:           "1,20,1",
			expectedSuccess: 0,
			expectedError:   fmt.Errorf("invalid input"),
		},
		{
			name:            "should return error when more than one dot.",
			input:           "1.20.2",
			expectedSuccess: 0,
			expectedError:   fmt.Errorf("invalid input"),
		},
		{
			name:            "should return error when a dot with no number.",
			input:           ".",
			expectedSuccess: 0,
			expectedError:   fmt.Errorf("invalid input"),
		},
		{
			name:            "should return error when a comma with no number.",
			input:           ",",
			expectedSuccess: 0,
			expectedError:   fmt.Errorf("invalid input"),
		},
		{
			name:            "should return error when comma and dot.",
			input:           "1,2.3",
			expectedSuccess: 0,
			expectedError:   fmt.Errorf("invalid input"),
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := RealToCents(tt.input)

			if tt.expectedError != nil && err == nil {
				t.Errorf("got %d expected an error", got)
				return
			}

			if got != tt.expectedSuccess {
				t.Errorf("got %d expected %d", got, tt.expectedSuccess)
			}
		})
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

func TestCentsToReal(t *testing.T) {
	type testCase struct {
		name            string
		input           int
		expectedSuccess string
		expectedError   error
	}
	testCases := []testCase{
		{
			name:            "should return value in R$X,XX format, receives in cents",
			input:           101,
			expectedSuccess: "R$1,01",
			expectedError:   nil,
		},
		{
			name:            "should return value in R$X,XX format, receives less than 100 in cents",
			input:           10,
			expectedSuccess: "R$0,10",
			expectedError:   nil,
		},
		{
			name:            "should return value in R$X,XX format, receives less than 100 in cents",
			input:           1,
			expectedSuccess: "R$0,01",
			expectedError:   nil,
		},
		{
			name:            "should return R$0,00 when receives 0",
			input:           0,
			expectedSuccess: "R$0,00",
			expectedError:   nil,
		},
		{
			name:            "should return error when receives negative",
			input:           -1,
			expectedSuccess: "",
			expectedError:   fmt.Errorf("invalid input"),
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := CentsToReal(tt.input)

			if tt.expectedError != nil && err == nil {
				t.Errorf("got %s expected an error", got)
				return
			}

			if got != tt.expectedSuccess {
				t.Errorf("got %s expected %s", got, tt.expectedSuccess)
			}
		})
	}
}
