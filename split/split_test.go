package main

import (
	"testing"
)

func TestSplitBill(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name            string
		itemList        ItemList
		emailList       EmailList
		expectedSuccess map[string]string
		expectedError   bool // change to a variable
	}

	testCases := []testCase{
		{
			name: "should return an equal split for all emails",
			itemList: []Item{
				{
					"Cerveja",
					"10,00",
					10,
				},
			},
			emailList: []string{
				"a@email.com",
				"b@email.com",
			},
			expectedSuccess: map[string]string{
				"a@email.com": "R$50,00",
				"b@email.com": "R$50,00",
			},
		},
		{
			name: "should return an error",
			itemList: []Item{
				{
					"Cerveja",
					"10,00a",
					10,
				},
			},
			emailList: []string{
				"a@email.com",
				"b@email.com",
			},
			expectedSuccess: map[string]string{},
			expectedError:   true,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.itemList.SplitBill(tt.emailList)
			if tt.expectedError {
				assertError(t, got, err)
				return
			}
			assertSplit(t, got, tt.expectedSuccess)
		})
	}
}

func assertSplit(t testing.TB, got, expected map[string]string) {
	for email := range got {
		if got[email] != expected[email] {
			t.Errorf("got %v expected %v", got, expected)
		}
	}
}

func assertError(t testing.TB, got map[string]string, err error) {
	if err == nil {
		t.Errorf("got %v expected an error", got)
	}
}
