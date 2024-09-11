# Go library for the SarePay API.

sarepay-go is a Go client library for accessing the SarePay API.

<!-- Where possible, the services available on the client groups the API into logical chunks and correspond to the structure of the Paystack API documentation at https://developers.paystack.co/v1.0/reference. -->

## Usage

``` go
    var (
        res map[string]interface{}
        c http.Client
    )

	papiKey := "PUBLIC-test"
	merchantKey := "test"


	client := sarepay.NewClient(papiKey, &c)

	transRequest := sarepay.TransactionInput{
		Key:    papiKey,
		Token:  merchantKey,
		Amount: 10000,
		Customer: sarepay.Customer{
			Name:  "philip",
			Email: "test@test.com",
		},
		Reference: "jekkditestsffarepddddassy",
	}
    
	res, err := client.Transaction.Initialize(&transRequest)
	if err != nil {
		//do something with error
	}

	datas, ok := res["data"].(map[string]interface{})
	if !ok {
		fmt.Println("Error: 'data' is not a map")
		return
	}

	link, ok := datas["link"].(string)
	if !ok {
		fmt.Println("Error: 'link' is not a string")
		return
	}

	// Print the link
	fmt.Println("Link:", link)
```





## CONTRIBUTING
Contributions are of course always welcome. The calling pattern is pretty well established, so adding new methods is relatively straightforward. Please make sure the build succeeds and the test suite passes.
