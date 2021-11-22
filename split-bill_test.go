package main

import (
	"testing"
)

var billingLists = []struct {
	itemList  ItemList
	emailList EmailList
	want      map[string]string
}{
	{itemList: []Item{
		{
			ProductName: "Cerveja",
			UnitPrice:   "10,00",
			Amount:      12,
		}},
		emailList: []string{
			"a",
			"b",
		},
		want: map[string]string{"a": "R$60,00", "b": "R$60,00"},
	},
}

func TestSplitBill(t *testing.T) {
	for _, tt := range billingLists {
		got, err := tt.itemList.SplitBill(tt.emailList)
		for email := range got {
			if got[email] != tt.want[email] {
				t.Errorf("got %v want %v", got, tt.want)

			}
		}
		if err == nil {

		}
	}
}
