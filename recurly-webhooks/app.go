package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/clbanning/mxj"
)

func handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("ioutil - Can't read body"))
	}

	mv, err := mxj.NewMapXml(body)
	mv["request_method"] = r.Method
	js, err := mv.JsonIndent("", "    ", true)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("MXJ - Json Parse Error"))
	}
	fmt.Printf("%s\n", string(js))

	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

/*
	- new_account_notification			: create subscription hosted pages recurly
	- paid_charge_invoice_notification	: create subscription hosted pages recurly
	- new_charge_invoice_notification	: create subscription hosted pages recurly
	- successful_payment_notification	: create subscription hosted pages recurly
	- billing_info_updated_notification	: create subscription hosted pages recurly
	- new_subscription_notification		: create subscription hosted pages recurly

	- canceled_account_notification		: archive account in recurly (account_code)

	- billing_info_updated_notification	: changed from hosted pages

	- expired_subscription_notification	: cancelled plan
		-- account.account_code: "email_address"
		-- subscription.uuid:  "subscription_id"

*/
