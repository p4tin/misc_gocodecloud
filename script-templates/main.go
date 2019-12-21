package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"
)

type Output struct {
	Prefix  string
	Text    string
	Postfix string
}

func main() {
	mainLineInfo := &Output{
		Prefix:  "{{.Count}}",
		Text:    " items are made of ",
		Postfix: "{{.Material}}",
	}
	fmt.Printf("%+v\n", mainLineInfo)

	datas := make([]map[string]interface{}, 0)
	datas = append(datas, map[string]interface{}{"Material": "wool", "Count": 17})
	datas = append(datas, map[string]interface{}{"Material": "cotton", "Count": 23})

	for _, data := range datas {
		info := applyTemplate(mainLineInfo, data)

		fmt.Printf("%+v\n", info)
	}
}

func applyTemplate(lineInfo *Output, data map[string]interface{}) *Output {
	bites, err := json.Marshal(lineInfo)
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("test").Parse(string(bites))
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, data)
	if err != nil {
		panic(err)
	}

	var output Output
	err = json.Unmarshal(tpl.Bytes(), &output)
	if err != nil {
		panic(err)
	}

	return &output
}
