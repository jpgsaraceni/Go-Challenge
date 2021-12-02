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

	testCases := []testCase{
		{
			name: "should return an equal split for all emails",
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
		{
			name: "should return an equal split for all emails",
			itemList: []Item{
				{"Cerveja", 10, 10},
				{"Petisco", 40, 2},
			},
			emailList: []string{
				"a@email.com",
				"b@email.com",
				"c@email.com",
				"d@email.com",
				"e@email.com",
			},
			expectedSuccess: map[string]int{
				"a@email.com": 36,
				"b@email.com": 36,
				"c@email.com": 36,
				"d@email.com": 36,
				"e@email.com": 36,
			},
			expectedError: nil,
		},
		{
			name: "should return a repeated email error",
			itemList: []Item{
				{"Cerveja", 10, 10},
				{"Petisco", 40, 2},
			},
			emailList: []string{
				"a@email.com",
				"a@email.com",
				"c@email.com",
				"d@email.com",
				"e@email.com",
			},
			expectedSuccess: map[string]int{},
			expectedError:   ErrRepeatedEmails,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.itemList.SplitBill(tt.emailList)
			if tt.expectedError != nil {
				assertError(t, err, tt.expectedError)
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

func assertError(t testing.TB, gotError, expectedError error) {
	if gotError != expectedError {
		t.Errorf("got %v error expected %v error", gotError, expectedError)
	}
}

type SpyShuffler struct {
	// how do I mock this?
}

func (s *SpyShuffler) Shuffle() {

}
