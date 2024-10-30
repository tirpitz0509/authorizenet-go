package main

import (
	"fmt"
	"github.com/tirpitz0509/authorizenet-go"
)

var newTransactionId string

func main() {

	apiName := "25FnYp3Q2"
	apiKey := "68J3Uhw52tfQZx8C"
	client := authorizenet.New(apiName, apiKey, true)

	connected, err := client.IsConnected()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if connected {
		fmt.Println("Connected to Authorize.net!")
	}

	ChargeCustomer(client)
	//client.VoidTransaction()
}

func ChargeCustomer(client *authorizenet.Client) {

	newTransaction := authorizenet.NewTransaction{
		Amount: "14.75",
		CreditCard: authorizenet.CreditCard{
			CardNumber:     "4012888818888",
			ExpirationDate: "08/25",
			CardCode:       "393",
		},
		BillTo: &authorizenet.BillTo{
			FirstName:   "Timmy",
			LastName:    "Jimmy",
			Address:     "1111 green ct",
			City:        "los angeles",
			State:       "CA",
			Zip:         "43534",
			Country:     "USA",
			PhoneNumber: "8885555555",
		},
		RefId:     "123123",
		InvoiceId: "123123",
	}
	res, err := newTransaction.Charge(*client)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if res.Approved() {
		newTransactionId = res.TransactionID()
		fmt.Println("Transaction was Approved! #", res.TransactionID())
	}
}

func VoidTransaction(client *authorizenet.Client) {

	newTransaction := authorizenet.PreviousTransaction{
		RefId: newTransactionId,
	}
	res, _ := newTransaction.Void(*client)
	if res.Approved() {
		fmt.Println("Transaction was Voided!")
	}

}
