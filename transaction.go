package sarepay

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// TransactionService handles operations related to transactions
// For more details see
type TransactionService service

type Customer struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type Metadata struct {
	TaxID      string `json:"taxIs,omitempty"`
	CustomerID string `json:"customerId,omitempty"`
}

type InitializeTransactionInput struct {
	Key                  string   `json:"key,omitempty"`
	Token                string   `json:"token,omitempty"`
	Amount               int      `json:"amount,omitempty"`
	Currency             string   `json:"currency,omitempty"`
	FeeBearer            string   `json:"feeBearer,omitempty"`
	DefaultPaymentMethod string   `json:"defaultPaymentMethod"`
	PaymentMethods       []string `json:"paymentMethods,omitempty"`
	Customer             Customer `json:"customer,omitempty"`
	ContainerID          string   `json:"containerId,omitempty"`
	Metadata             Metadata `json:"metadata,omitempty"`
	Reference            string   `json:"reference,omitempty"`
}

// Initialize initiates a transaction process
// For more details see https://documenter.getpostman.com/view/28866628/2s9Y5bRh7W#678924c7-250b-460b-b870-8dbd68c4de16
func (s *TransactionService) Initialize(txn *InitializeTransactionInput) (Response, error) {
	u := baseURL + "/payments/initialize"
	resp := Response{}
	client := http.Client{}

	payload, err := json.Marshal(txn)
	if err != nil {
		return resp, err
	}

	request, err := http.NewRequest("POST", u, bytes.NewBuffer(payload))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("X-Pub-Key", s.client.publicKey)

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