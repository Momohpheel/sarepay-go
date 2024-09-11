package sarepay

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type VirtualAccountService service

type VirtualAccountRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	OtherName    string `json:"other_name"`
	Bvn          string `json:"bvn"`
	Dob          string `json:"dob"`
	PhoneNumber  string `json:"phone_number"`
	BusinessType string `json:"business_type"`
	Type         string `json:"type"`
	Currency     string `json:"currency"`
}

type VirtualAccountResponse struct {
	Data struct {
		AccountNumber    string `json:"account_number"`
		AccountName      string `json:"account_name"`
		AccountReference string `json:"account_reference"`
		Bank             string `json:"bank"`
		Status           string `json:"status"`
		Type             string `json:"type"`
		ExpiresAt        string `json:"expires_at"`
		ValidityType     string `json:"validity_type"`
	} `json:"data"`
}

// Generates virtual account number
// For more details see https://documenter.getpostman.com/view/28866628/2s9Y5bRh7W#678924c7-250b-460b-b870-8dbd68c4de16
func (s *VirtualAccountService) GeneratePermanentAccount(txn *VirtualAccountRequest) (VirtualAccountResponse, error) {
	u := baseURL + "/virtual-accounts/permanents"
	resp := VirtualAccountResponse{}
	client := http.Client{}

	payload, err := json.Marshal(txn)
	if err != nil {
		return resp, err
	}

	request, err := http.NewRequest("POST", u, bytes.NewBuffer(payload))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("api-key", s.client.secretKey)

	if err != nil {
		return resp, err
	}

	respo, err := client.Do(request)
	if err != nil {
		return resp, err
	}
	defer respo.Body.Close()

	err = json.NewDecoder(respo.Body).Decode(&resp)
	if err != nil {
		return resp, err
	}

	return resp, err
}
