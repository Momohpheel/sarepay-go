package sarepay

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type TransferService service

type TransferRequest struct {
	CustomerReference string `json:"customer_reference"`
	AccountNumber     string `json:"account_number"`
	BankCode          string `json:"bank_code"`
	Amount            string `json:"amount"`
	Narration         string `json:"narration"`
	RecipientName     string `json:"recipient_name"`
}

type TransferResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Reference              string  `json:"reference"`
		Amount                 string  `json:"amount"`
		Charge                 string  `json:"charge"`
		Status                 string  `json:"status"`
		RecipientName          string  `json:"recipient_name"`
		RecipientBankCode      string  `json:"recipient_bank_code"`
		RecipientAccountNumber string  `json:"recipient_account_number"`
		ProcessorReference     *string `json:"processor_reference"` // Use *string to handle null values
		MerchantReference      string  `json:"merchant_reference"`
	} `json:"data"`
	Message string `json:"message"`
}

// Process Transfer
// For more details see https://documenter.getpostman.com/view/28866628/2s9Y5bRh7W#678924c7-250b-460b-b870-8dbd68c4de16
func (s *TransferService) ProcessTransfer(txn *TransferRequest) (TransferResponse, error) {
	u := baseURL + "/disbursement/transact"
	resp := TransferResponse{}
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

// Verify Transfer
func (s *TransferService) VerifyTransfer(reference string) (TransferResponse, error) {
	u := baseURL + "/disbursement/requery/" + reference
	resp := TransferResponse{}
	client := http.Client{}

	request, err := http.NewRequest("GET", u, nil)
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

type AccountDetails struct {
	AccountNumber string `json:"account_number"`
	BankCode      string `json:"bank_code"`
}

type AccountResponse struct {
	Success bool `json:"success"`
	Data    struct {
		AccountNumber string `json:"account_number"`
		AccountName   string `json:"account_name"`
	} `json:"data"`
	Message string `json:"message"`
}

// Account Lookup
func (s *TransferService) AccountLookup(txn *AccountDetails) (AccountResponse, error) {
	u := baseURL + "/api/disbursement/accounts/validate"
	resp := AccountResponse{}
	client := http.Client{}

	payload, err := json.Marshal(txn)
	if err != nil {
		return resp, err
	}

	request, err := http.NewRequest(http.MethodPost, u, bytes.NewBuffer(payload))
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
