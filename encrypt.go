package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

var key []byte

func init() {
	content, err := ioutil.ReadFile("./aes.key")
	if err != nil {
		log.Fatalln("Could not read key file (aes.key)")
	}
	key, err = base64.StdEncoding.DecodeString(string(content))
	if err != nil {
		log.Fatalln("Could not decode base64 key from file (aes.key)")
	}
}

func main() {
	plaintext := []byte("Some really really really long plaintext")
	fmt.Printf("%s\n", plaintext)
	ciphertext, err := encrypt(key, plaintext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0x - %d\n", ciphertext, len(ciphertext))
	result, err := decrypt(key, ciphertext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", result)
}

func encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}
