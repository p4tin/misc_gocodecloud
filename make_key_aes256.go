package main

import (
	"crypto/rand"
	"fmt"
	b64 "encoding/base64"
	"encoding/hex"
	"io/ioutil"
)

func main() {
    key := make([]byte, 32)

    _, err := rand.Read(key)
    if err != nil {
        // handle error here
    }


    fmt.Println("=================")
    fmt.Println(key)
    fmt.Println("=================")
    sEnc := b64.StdEncoding.EncodeToString(key)
    fmt.Println(sEnc)
	fmt.Println("=================")
	sEnc2 := hex.EncodeToString(key)
	fmt.Println(sEnc2)
	fmt.Println("=================")
    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(sDec)
    fmt.Println("=================")
	sDec2, _ := hex.DecodeString(sEnc2)
	fmt.Println(sDec2)

	d1 := []byte(sEnc)
	_ = ioutil.WriteFile("./aes.key", d1, 0644)

	content, err := ioutil.ReadFile("./aes.key")
	if err != nil {
		//Do something
	}
	fmt.Println(string(content))
	fmt.Println("=================")
	sDec, _ = b64.StdEncoding.DecodeString(string(content))
	fmt.Println(sDec)
}

