package main

import (
	"testing"
)


func TestFetch(t *testing.T) {
	_, err := apiClient.Fetch()
	if err != nil {
		t.Errorf("Test failed for fetching accounts")
	}
}
