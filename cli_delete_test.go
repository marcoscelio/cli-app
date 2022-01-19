package main

import (
	"fmt"
	"testing"
)

func TestDelete(t *testing.T) {
	err := apiClient.Delete(fmt.Sprintf(`%s`, AccountID), "0")
	if err != nil {
		t.Errorf("Test failed for deleting accounts")
	}
}
