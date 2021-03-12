package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/robfig/cron"
)

type Payload struct {
	Refinements []struct {
		Type           string `json:"type"`
		Value          string `json:"value"`
		NavigationName string `json:"navigationName"`
	} `json:"refinements"`
}

func main() {
	waiter := make(chan struct{})

	c := cron.New()
	c.AddFunc("*/10 * * * *", NewSearchService().CallSearchService)
	c.Start()

	<-waiter
}

type SearchService struct {
	Data Payload
}

func NewSearchService() *SearchService {
	return &SearchService{
		Data: Payload{
			Refinements: []struct {
				Type           string `json:"type"`
				Value          string `json:"value"`
				NavigationName string `json:"navigationName"`
			}{
				{Type: "Value", Value: "FP Collection", NavigationName: "tile.product.brand"},
			},
		},
	}
}

func (ss *SearchService) CallSearchService() {
	payloadBytes, err := json.Marshal(ss.Data)
	if err != nil {
		panic(err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://staging-search.urbncloud.com/api/catalog-search-service/v0/fp-us/tiles?pool=US_DIRECT&from=0&pageSize=86&navigationSlug=shoes", body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Urbn-Channel", "web")
	req.Header.Set("X-Urbn-Currency", "USD")
	req.Header.Set("X-Urbn-Language", "en-US")

	var netTransport = &http.Transport{
		MaxIdleConns:        0,
		MaxIdleConnsPerHost: 0,
		DialContext: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	var netClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}

	resp, err := netClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Print("StatusCode: ", resp.StatusCode, " - ")

	addr, err := net.LookupIP("staging-search.urbncloud.com")
	if err != nil {
		fmt.Println("Unknown host")
	} else {
		fmt.Println("IP address: ", addr)
	}
}
