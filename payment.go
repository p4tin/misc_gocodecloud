package main

type MerchantAuthenticationData struct {
	Name           string `json:"name"`
	TransactionKey string `json:"transactionKey"`
}

type AddressData struct {
	Address   string `json:"address"`
	City      string `json:"city"`
	Company   string `json:"company"`
	Country   string `json:"country"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
}

type CustomerData struct {
	ID string `json:"id"`
}
type DutyData struct {
	Amount      string `json:"amount"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type LineItemData struct {
	Description string `json:"description"`
	ItemID      string `json:"itemId"`
	Name        string `json:"name"`
	Quantity    string `json:"quantity"`
	UnitPrice   string `json:"unitPrice"`
}

type CreditCardData struct {
	CardCode       string `json:"cardCode"`
	CardNumber     string `json:"cardNumber"`
	ExpirationDate string `json:"expirationDate"`
}

type AuthRequest struct {
	CreateTransactionRequest struct {
		MerchantAuthentication MerchantAuthenticationData `json:"merchantAuthentication"`
		RefID                  string                     `json:"refId"`
		TransactionRequest     struct {
			Amount     string       `json:"amount"`
			BillTo     AddressData  `json:"billTo"`
			Customer   CustomerData `json:"customer"`
			CustomerIP string       `json:"customerIP"`
			Duty       DutyData     `json:"duty"`
			LineItems  struct {
				LineItem LineItemData `json:"lineItem"`
			} `json:"lineItems"`
			Payment struct {
				CreditCard CreditCardData `json:"creditCard"`
			} `json:"payment"`
			PoNumber string      `json:"poNumber"`
			ShipTo   AddressData `json:"shipTo"`
			Shipping struct {
				Amount      string `json:"amount"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"shipping"`
			Tax struct {
				Amount      string `json:"amount"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"tax"`
			TransactionSettings struct {
				Setting struct {
					SettingName  string `json:"settingName"`
					SettingValue string `json:"settingValue"`
				} `json:"setting"`
			} `json:"transactionSettings"`
			TransactionType string `json:"transactionType"`
			UserFields      struct {
				UserField []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"userField"`
			} `json:"userFields"`
		} `json:"transactionRequest"`
	} `json:"createTransactionRequest"`
}

func main() {
}
