package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/marcoscelio/cli-app/models"
	"net/http"
	"io/ioutil"
)

type ApiClient struct {
	url string
}

func (c *ApiClient) Create(account models.AccountData) error {
	fmt.Println("Creating account")
	data, err := json.Marshal(account)
	if err != nil {
		fmt.Println(err)
		return err
	}
	postBody := fmt.Sprintf(`{"data":%s}: `, string(data))
	fmt.Println("Payload:", postBody)
	payloadData := bytes.NewBuffer([]byte(postBody))
	resp, err := http.Post("http://accountapi:8080/v1/organisation/accounts", "application/json", payloadData)

	defer resp.Body.Close()

	if err != nil {
		fmt.Println("  Error      :", err)
		return err
	}
	// body, err := ioutil.ReadAll(resp.Body)
	return nil
}

func (c *ApiClient) Fetch() (*[]models.AccountData, error) {
	fmt.Println("Fetching accounts")
	resp, err := http.Get("http://accountapi:8080/v1/organisation/accounts")

	if err != nil {
		fmt.Println("  Error: ", err)
		return nil, err
	}
	var data models.FetchDataResponse
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &data)
	return data.Data, nil
}

func (c *ApiClient) Delete(accountId string, version string) error {
	fmt.Println("Deleteing account")
	client := resty.New()
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		EnableTrace().
		Delete(fmt.Sprintf("http://accountapi:8080/v1/organisation/accounts/%s?version=%s", accountId, version))

	if err != nil {
		fmt.Println("  Error      :", err)
		return err
	}
	return nil
}
