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
		expectedSuccess map[string]int
		expectedError   error
	}

	// var errInput = errors.New("invalid input")

	testCases := []testCase{
		{
			name: "should return an equal split (in cents) for all emails",
			itemList: []Item{
				{
					"Cerveja",
					10,
					10,
				},
			},
			emailList: []string{
				"a@email.com",
				"b@email.com",
			},
			expectedSuccess: map[string]int{
				"a@email.com": 50,
				"b@email.com": 50,
			},
		},
		// {
		// 	name: "should return an error",
		// 	itemList: []Item{
		// 		{
		// 			"Cerveja",
		// 			"10,00a",
		// 			10,
		// 		},
		// 	},
		// 	emailList: []string{
		// 		"a@email.com",
		// 		"b@email.com",
		// 	},
		// 	expectedSuccess: map[string]int{},
		// 	expectedError:   errInput,
		// },
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.itemList.SplitBill(tt.emailList)
			if tt.expectedError != nil {
				assertError(t, got, err)
				return
			}
			assertSplit(t, got, tt.expectedSuccess)
		})
	}
}

func assertSplit(t testing.TB, got, expected map[string]int) {
	for email := range got {
		if got[email] != expected[email] {
			t.Errorf("got %v expected %v", got, expected)
		}
	}
}

func assertError(t testing.TB, got map[string]int, err error) {
	if err == nil {
		t.Errorf("got %v expected an error", got)
	}
}
