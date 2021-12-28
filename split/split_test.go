package main

import (
	"errors"
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
		shuffleMock     SpyShuffler
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
			shuffleMock: SpyShuffler{
				ShuffleFunc: func(emails EmailList) EmailList {
					return EmailList{
						"a@email.com",
						"b@email.com",
					}
				},
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
			shuffleMock: SpyShuffler{
				ShuffleFunc: func(emails EmailList) EmailList {
					return EmailList{
						"a@email.com",
						"b@email.com",
						"c@email.com",
						"d@email.com",
						"e@email.com",
					}
				},
			},
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
			shuffleMock: SpyShuffler{
				ShuffleFunc: func(emails EmailList) EmailList {
					return EmailList{
						"a@email.com",
						"a@email.com",
						"c@email.com",
						"d@email.com",
						"e@email.com",
					}
				},
			},
		},
		{
			name: "should return one unit more for some emails",
			itemList: []Item{
				{"Cerveja", 1, 10},
				{"Petisco", 7, 1},
			},
			emailList: []string{
				"a@email.com",
				"b@email.com",
				"c@email.com",
				"d@email.com",
				"e@email.com",
			},
			expectedSuccess: map[string]int{
				"b@email.com": 4,
				"a@email.com": 4,
				"d@email.com": 3,
				"c@email.com": 3,
				"e@email.com": 3,
			},
			shuffleMock: SpyShuffler{
				ShuffleFunc: func(emails EmailList) EmailList {
					return EmailList{
						"b@email.com",
						"a@email.com",
						"d@email.com",
						"c@email.com",
						"e@email.com",
					}
				},
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.emailList.checkForRepeated()
			if tt.expectedError != nil {
				assertError(t, err, tt.expectedError)

				return
			}

			got, err := tt.itemList.SplitBill(tt.emailList, &tt.shuffleMock)
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
	if !errors.Is(gotError, expectedError) {
		t.Errorf("got %v error expected %v error", gotError, expectedError)
	}
}

type SpyShuffler struct {
	ShuffleFunc func(emails EmailList) EmailList
}

func (s *SpyShuffler) Shuffle(emails EmailList) EmailList {
	return s.ShuffleFunc(emails)
}
