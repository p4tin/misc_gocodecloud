package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

const sharedSecret = "sup3rs3cr3t!!"

func signString(stringToSign []byte, sharedSecret []byte) []byte {
	h := hmac.New(sha512.New, sharedSecret)
	h.Write(stringToSign)

	return h.Sum(nil)
}

func main() {
	payload := make(map[string]string)
	payload["name"] = "joe smith"
	payload["category"] = "people"
	payload["action"] = "transport"
	payload["where"] = "pluto"
	payload["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)

	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	signature := signString(jsonBytes, []byte(sharedSecret))

	fmt.Println(fmt.Sprintf("%s - %s",
		base64.URLEncoding.EncodeToString(jsonBytes),
		base64.URLEncoding.EncodeToString(signature)))
}
