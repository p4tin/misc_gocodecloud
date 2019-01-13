package main

import (
	"net/http"
	"strings"
	"os"
	"io"
	"log"
)

func main() {

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	body := strings.NewReader(`
{
  "from": 0,
  "size": 1000,
  "query": {
	"match": {
		"product.keywords": "Basketball"
	}
  }
}`)
	req, err := http.NewRequest("GET", "http://search-catalog-urbn-stg-wbv3ad3tfm5nnd63xjiati542e.us-east-1.es.amazonaws.com/catalog/product/_search", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
