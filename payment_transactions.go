package authorizenet

import (
	"encoding/json"
)

func (tranx NewTransaction) Charge(c Client) (*TransactionResponse, error) {
	var new TransactionRequest
	new = TransactionRequest{
		TransactionType: "authCaptureTransaction",
		Amount:          tranx.Amount,
		Payment: &Payment{
			CreditCard: tranx.CreditCard,
		},
		BillTo:   tranx.BillTo,
		AuthCode: tranx.AuthCode,
		Order: &Order{
			InvoiceNumber: tranx.InvoiceId,
		},
		CustomerData: tranx.CustomerData,
		CustomerIP:   tranx.CustomerIP,
		UserFields:   tranx.UserFields,
	}
	if tranx.CreateProfile {
		new.Profile = &Profile{
			ProfileCreate: true,
		}
	}
	res, err := c.SendTransactionRequest(new)
	return res, err
}

func (tranx NewTransaction) ChargeProfile(profile Customer, c Client) (*TransactionResponse, error) {
	var new TransactionRequest
	new = TransactionRequest{
		TransactionType: "authCaptureTransaction",
		Amount:          tranx.Amount,
		Profile: &Profile{
			CustomerProfileId: profile.ID,
			PaymentProfile: &PaymentProfile{
				PaymentProfileId: profile.PaymentID,
			},
		},
		Order: &Order{
			InvoiceNumber: tranx.InvoiceId,
		},
	}
	if tranx.Authorize {
		new.TransactionType = "authOnlyTransaction"
	}
	res, err := c.SendTransactionRequest(new)
	return res, err
}

func (tranx NewTransaction) AuthOnly(c Client) (*TransactionResponse, error) {
	var new TransactionRequest
	new = TransactionRequest{
		TransactionType: "authOnlyTransaction",
		Amount:          tranx.Amount,
		Payment: &Payment{
			CreditCard: tranx.CreditCard,
		},
		BillTo: tranx.BillTo,
		Order: &Order{
			InvoiceNumber: tranx.InvoiceId,
		},
		CustomerData: tranx.CustomerData,
		CustomerIP:   tranx.CustomerIP,
		UserFields:   tranx.UserFields,
	}
	if tranx.CreateProfile {
		new.Profile = &Profile{
			ProfileCreate: true,
		}
	}
	res, err := c.SendTransactionRequest(new)
	return res, err
}

func (tranx NewTransaction) Refund(c Client) (*TransactionResponse, error) {
	var new TransactionRequest
	new = TransactionRequest{
		TransactionType: "refundTransaction",
		Amount:          tranx.Amount,
		RefTransId:      tranx.RefTransId,
		Payment: &Payment{
			CreditCard: tranx.CreditCard,
		},
		Order: &Order{
			InvoiceNumber: tranx.InvoiceId,
		},
	}
	res, err := c.SendTransactionRequest(new)
	return res, err
}

func (tranx PreviousTransaction) Void(c Client) (*TransactionResponse, error) {
	var new TransactionRequest
	new = TransactionRequest{
		TransactionType: "voidTransaction",
		RefTransId:      tranx.RefId,
	}
	res, err := c.SendTransactionRequest(new)
	return res, err
}

func (tranx PreviousTransaction) Capture(c Client) (*TransactionResponse, error) {
	var new TransactionRequest
	new = TransactionRequest{
		TransactionType: "priorAuthCaptureTransaction",
		RefTransId:      tranx.RefId,
	}
	res, err := c.SendTransactionRequest(new)
	return res, err
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

func (c Client) SendTransactionRequest(input TransactionRequest) (*TransactionResponse, error) {
	action := CreatePayment{
		CreateTransactionRequest: CreateTransactionRequest{
			MerchantAuthentication: c.GetAuthentication(),
			TransactionRequest:     input,
		},
	}
	req, err := json.Marshal(action)
	if err != nil {
		return nil, err
	}
	res, err := c.SendRequest(req)
	var dat TransactionResponse
	err = json.Unmarshal(res, &dat)
	if err != nil {
		return nil, err
	}
	return &dat, err
}

type NewTransaction struct {
	Amount        string        `json:"amount,omitempty"`
	InvoiceId     string        `json:"invoiceId,omitempty"`
	RefTransId    string        `json:"refTransId,omitempty"`
	CreditCard    CreditCard    `json:"payment,omitempty"`
	AuthCode      string        `json:"authCode,omitempty"`
	BillTo        *BillTo       `json:"billTo,omitempty"`
	RefId         string        `json:"refId,omitempty"`
	CustomerData  *CustomerData `json:"customerData,omitempty"`
	CustomerIP    string        `json:"customerIP,omitempty"`
	UserFields    *UserFields   `json:"userFields,omitempty"`
	CreateProfile bool          `json:"createProfile,omitempty"`
	Authorize     bool          `json:"authorize,omitempty"`
}

type PreviousTransaction struct {
	RefId  string `json:"refTransId,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type TransactionResponse struct {
	Response TranxResponse `json:"transactionResponse"`
	MessagesResponse
}

type TranxResponse struct {
	ResponseCode   string `json:"resCode"`
	AuthCode       string `json:"authCode"`
	AvsResultCode  string `json:"avsResultCode"`
	CvvResultCode  string `json:"cvvResultCode"`
	CavvResultCode string `json:"cavvResultCode"`
	TransID        string `json:"transId"`
	RefTransID     string `json:"refTransID"`
	TransHash      string `json:"transHash"`
	TestRequest    string `json:"testRequest"`
	AccountNumber  string `json:"accountNumber"`
	AccountType    string `json:"accountType"`
	Errors         []struct {
		ErrorCode string `json:"errorCode"`
		ErrorText string `json:"errorText"`
	} `json:"errors"`
	TransactionMessages
	TransHashSha2   string `json:"transHashSha2"`
	Message         TransactionMessages
	ProfileResponse CustomProfileResponse `json:"profileResponse"`
}

type TransactionMessages struct {
	Message []struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"messages"`
}

type Message struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

type MerchantAuthentication struct {
	Name           string `json:"name,omitempty"`
	TransactionKey string `json:"transactionKey,omitempty"`
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
	ID          string `json:"id,omitempty"`
	Email       string `json:"email,omitempty"`
	PaymentID   string `json:"paymentId,omitempty"`
	ShippingID  string `json:"shippingId,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
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

type Order struct {
	InvoiceNumber string `json:"invoiceNumber,omitempty"`
	Description   string `json:"description,omitempty"`
}

type TransactionRequest struct {
	TransactionType string        `json:"transactionType,omitempty"` //
	Amount          string        `json:"amount,omitempty"`          //
	Payment         *Payment      `json:"payment,omitempty"`         //
	RefTransId      string        `json:"refTransId,omitempty"`      //
	AuthCode        string        `json:"authCode,omitempty"`        //
	Profile         *Profile      `json:"profile,omitempty"`         //
	LineItems       *LineItems    `json:"lineItems,omitempty"`       //
	Order           *Order        `json:"order,omitempty"`           //
	CustomerData    *CustomerData `json:"customer,omitempty"`        //
	BillTo          *BillTo       `json:"billTo,omitempty"`          //
	ShipTo          *Address      `json:"shipTo,omitempty"`          //
	CustomerIP      string        `json:"customerIP,omitempty"`      //
	//Tax                 Tax                 `json:"tax,omitempty"` //
	//Duty                Duty                `json:"duty,omitempty"` //
	//Shipping            Shipping            `json:"shipping,omitempty"` //
	//PoNumber            string              `json:"poNumber,omitempty"` //
	//TransactionSettings TransactionSettings `json:"transactionSettings,omitempty"` //
	UserFields *UserFields `json:"userFields,omitempty"` //

}

type Address struct {
	FirstName         string `json:"firstName,omitempty"`
	LastName          string `json:"lastName,omitempty"`
	Company           string `json:"company,omitempty"`
	Address           string `json:"address,omitempty"`
	City              string `json:"city,omitempty"`
	State             string `json:"state,omitempty"`
	Zip               string `json:"zip,omitempty"`
	Country           string `json:"country,omitempty"`
	PhoneNumber       string `json:"phoneNumber,omitempty"`
	FaxNumber         string `json:"faxNumber,omitempty"`
	CustomerAddressID string `json:"customerAddressId,omitempty"`
}

type CustomerData struct {
	CustomerType string `json:"type,omitempty"`
	ID           string `json:"id,omitempty"`
	Email        string `json:"email,omitempty"`
}
