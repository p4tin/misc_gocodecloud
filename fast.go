package main

import (
	"github.com/ddo/go-fast"
	"time"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"strconv"
	"log"
)

func main() {
	for {
		fastCom := fast.New()
		err := fastCom.Init()
		if err != nil {
			panic(err)
		}

		// get urls
		urls, err := fastCom.GetUrls()
		if err != nil {
			panic(err)
		}

		// measure
		KbpsChan := make(chan float64)

		var lastVal float64
		go func() {
			for Kbps := range KbpsChan {
				lastVal = Kbps / 1000
			}
		}()

		err = fastCom.Measure(urls, KbpsChan)

		if err != nil {
			log.Println(err)
		}
		SendEz("paul@cp-soft.com", "home-internet-speed", lastVal)
		time.Sleep(5 * time.Minute)
	}
}


func SendEz(ezkey string, stat string, value float64) {
	urlstr := "http://api.stathat.com/ez"

	form := url.Values{}
	form.Add("ezkey", ezkey)
	form.Add("stat", stat)
	form.Add("value", strconv.FormatFloat(value, 'f', 6, 64))
	req, err := http.NewRequest("POST", urlstr, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Latest speed:>%f, response Status: %s\n", value, resp.Status)
}
