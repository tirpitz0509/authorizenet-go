package AuthorizeCIM

import (
	"encoding/json"
)

func (tranx NewTransaction) Charge() TransactionResponse {
	var new TransactionRequest
	new = TransactionRequest{
		TransactionType: "authCaptureTransaction",
		Amount:          tranx.Amount,
		Payment: &Payment{
			CreditCard: tranx.CreditCard,
		},
		AuthCode: tranx.AuthCode,
	}
	response, _ := SendTransactionRequest(new)
	return response
}

func (tranx NewTransaction) AuthOnly() TransactionResponse {
	var new TransactionRequest
	new = TransactionRequest{
		TransactionType: "authOnlyTransaction",
		Amount:          tranx.Amount,
		Payment: &Payment{
			CreditCard: tranx.CreditCard,
		},
	}
	response, _ := SendTransactionRequest(new)
	return response
}

func (tranx NewTransaction) Refund() TransactionResponse {
	var new TransactionRequest
	new = TransactionRequest{
		TransactionType: "refundTransaction",
		Amount:          tranx.Amount,
		RefTransId:      tranx.RefTransId,
	}
	response, _ := SendTransactionRequest(new)
	return response
}

func (tranx PreviousTransaction) Void() TransactionResponse {
	var new TransactionRequest
	new = TransactionRequest{
		TransactionType: "voidTransaction",
		RefTransId:      tranx.RefId,
	}
	response, _ := SendTransactionRequest(new)
	return response
}

func (tranx PreviousTransaction) Capture() TransactionResponse {
	var new TransactionRequest
	new = TransactionRequest{
		TransactionType: "priorAuthCaptureTransaction",
		RefTransId:      tranx.RefId,
	}
	response, _ := SendTransactionRequest(new)
	return response
}

func UpdateSplitTenderGround() {

}

func DebitBankAccount() {

}

func CreditBankAccount() {

}

func ChargeTokenCard() {

}

func CreditAcceptPaymentTransaction() {

}

func GetAccessPaymentPage() {

}

func GetHostedPaymentPage() {

}

func SendTransactionRequest(input TransactionRequest) (TransactionResponse, interface{}) {
	action := CreatePayment{
		CreateTransactionRequest: CreateTransactionRequest{
			MerchantAuthentication: GetAuthentication(),
			TransactionRequest:     input,
		},
	}
	jsoned, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}
	response := SendRequest(jsoned)
	var dat TransactionResponse
	err = json.Unmarshal(response, &dat)
	if err != nil {
		panic(err)
	}
	return dat, err
}

type NewTransaction struct {
	Amount     string     `json:"amount,omitempty"`
	InvoiceId  string     `json:"invoiceId,omitempty"`
	RefTransId string     `json:"refTransId,omitempty"`
	CreditCard CreditCard `json:"payment,omitempty"`
	AuthCode   string     `json:"authCode,omitempty"`
}

type PreviousTransaction struct {
	RefId  string `json:"refTransId,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type TransactionResponse struct {
	Response struct {
		ResponseCode   string          `json:"responseCode"`
		AuthCode       string          `json:"authCode"`
		AvsResultCode  string          `json:"avsResultCode"`
		CvvResultCode  string          `json:"cvvResultCode"`
		CavvResultCode string          `json:"cavvResultCode"`
		TransID        string          `json:"transId"`
		RefTransID     string          `json:"refTransID"`
		TransHash      string          `json:"transHash"`
		TestRequest    string          `json:"testRequest"`
		AccountNumber  string          `json:"accountNumber"`
		AccountType    string          `json:"accountType"`
		Errors         []AuthNetErrors `json:"errors"`
		TransHashSha2  string          `json:"transHashSha2"`
	} `json:"transactionResponse"`
	Messages struct {
		ResultCode string    `json:"resultCode"`
		Message    []Message `json:"message"`
	} `json:"messages"`
}

type AuthNetErrors struct {
	ErrorCode string `json:"errorCode"`
	ErrorText string `json:"errorText"`
}

type Message struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

type MerchantAuthentication struct {
	Name           *string `json:"name,omitempty"`
	TransactionKey *string `json:"transactionKey,omitempty"`
}

type CreatePayment struct {
	CreateTransactionRequest CreateTransactionRequest `json:"createTransactionRequest,omitempty"`
}

type CreateTransactionRequest struct {
	MerchantAuthentication MerchantAuthentication `json:"merchantAuthentication,omitempty"`
	RefID                  string                 `json:"refId,omitempty"`
	TransactionRequest     TransactionRequest     `json:"transactionRequest,omitempty"`
}

type Payment struct {
	CreditCard CreditCard `json:"creditCard,omitempty"`
}

type CreditCard struct {
	CardNumber     string `json:"cardNumber,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty"`
	CardCode       string `json:"cardCode,omitempty"`
}

type LineItems struct {
	LineItem []LineItem `json:"lineItem,omitempty"`
}

type LineItem struct {
	ItemID      string `json:"itemId,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Quantity    string `json:"quantity,omitempty"`
	UnitPrice   string `json:"unitPrice,omitempty"`
}

type Shipping struct {
	Amount      string `json:"amount,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Duty struct {
	Amount      string `json:"amount,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Tax struct {
	Amount      string `json:"amount,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Customer struct {
	ID string `json:"id,omitempty"`
	PaymentID string `json:"id,omitempty"`
}

type TransactionSettings struct {
	Setting Setting `json:"setting,omitempty"`
}

type Setting struct {
	SettingName  string `json:"settingName,omitempty"`
	SettingValue string `json:"settingValue,omitempty"`
}

type UserFields struct {
	UserField []UserField `json:"userField,omitempty"`
}

type UserField struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type TransactionRequest struct {
	TransactionType string   `json:"transactionType,omitempty"`
	Amount          string   `json:"amount,omitempty"`
	Payment         *Payment `json:"payment,omitempty"`
	RefTransId      string   `json:"refTransId,omitempty"`
	AuthCode        string   `json:"authCode,omitempty"`
	//LineItems           LineItems           `json:"lineItems,omitempty"`
	//Tax                 Tax                 `json:"tax,omitempty"`
	//Duty                Duty                `json:"duty,omitempty"`
	//Shipping            Shipping            `json:"shipping,omitempty"`
	//PoNumber            string              `json:"poNumber,omitempty"`
	//Customer            Customer            `json:"customer,omitempty"`
	//BillTo              Address             `json:"billTo,omitempty"`
	//ShipTo              Address             `json:"shipTo,omitempty"`
	//CustomerIP          string              `json:"customerIP,omitempty"`
	//TransactionSettings TransactionSettings `json:"transactionSettings,omitempty"`
	//UserFields          UserFields          `json:"userFields,omitempty"`
}

type Address struct {
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	Address     string `json:"address,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	Zip         string `json:"zip,omitempty"`
	Country     string `json:"country,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
}