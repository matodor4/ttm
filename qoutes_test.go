package main

import "testing"

func TestGetQuotes(t *testing.T) {
	var quotes []string
	expectedLen := 5
	got, _ := GetQuotes(QuotesPath, quotes)

	if len(got) != expectedLen {
		t.Errorf("wrong fetching^ want %d but got %d", expectedLen, len(got))
	}
}
