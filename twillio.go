package main

import (
	"net/http"
	"net/url"
	"fmt"
	"strings"
)

func main() {
	// Set initial variables
	accountSid := "AC63960c1e63b8bd55f1a2c5e0d4732eda"
	authToken := "0159385bd9249c694a84bcde36ffed42"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// Build out the data for our message
	v := url.Values{}
	v.Set("To","12679084110")
	v.Set("From","14073783890")
	v.Set("Body","Alert Test Message")
	rb := *strings.NewReader(v.Encode())

	// Create Client
	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)
	fmt.Println(resp.Status)
}