package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type response struct {
	Kind        string
	ID          string
	OutputLabel string
	OutputMulti []struct {
		Label string
		Score string
	}
	OutputValue float64
}

var labels = map[string]string{
	"English": "🇬🇧 English",
	"Spanish": "🇪🇸 Spanish",
	"French":  "🇫🇷 French",
}

func do(query string) (*response, error) {
	values := url.Values{
		"model":  []string{"Language Detection"},
		"Phrase": []string{query},
	}
	res, err := http.PostForm("http://try-prediction.appspot.com/predict", values)
	if err != nil {
		return nil, err
	}
	var result response
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for s.Scan() {
		res, err := do(s.Text())
		if err != nil {
			fmt.Println("failed:", err)
			continue
		}
		label, ok := labels[res.OutputLabel]
		if !ok {
			label = res.OutputLabel
		}
		fmt.Printf("  ^^ That looks like %s to me\n", label)
		fmt.Print("> ")
	}
}
