package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/marcoscelio/cli-app/models"
	"testing"
)


func TestCreate(t *testing.T) {
	ID := AccountID
	OrganisationID := uuid.New()
	Name := "Test name"
	Country := "BR"
	BaseCurrency := "GBP"
	BankID := "400300"
	BankIDCode := "GBDSC"
	Bic := "NWBKGB22"
	data := fmt.Sprintf(`{
                            "type": "accounts",
                            "id": "%s",
                            "organisation_id": "%s",
                            "attributes": {
                                "name": ["%s"],
                                "country": "%s",
                                "base_currency": "%s",
                                "bank_id": "%s",
                                "bank_id_code": "%s",
                                "bic": "%s"
                            }
                        }`, ID, OrganisationID, Name, Country, BaseCurrency, BankID, BankIDCode, Bic)
	var accountData models.AccountData
	err := json.Unmarshal([]byte(data), &accountData)
	if err != nil {
		t.Errorf("Test failed for parsing account")
	}
	err = apiClient.Create(accountData)
	if err != nil {
		t.Errorf("Test failed for creating new account")
	}
}
