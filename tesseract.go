package main

import (
	"fmt"
	"github.com/otiai10/gosseract"
)

func main() {
	// This is the simlest way :)
	out := gosseract.Must(gosseract.Params{
		Src:       "/Users/Paul/Desktop/rp.png",
		Languages: "eng",
	})
	fmt.Println(out)

	//Using client
	client, _ := gosseract.NewClient()
	out, _ = client.Src("/Users/Paul/Desktop/rp.png").Out()
	fmt.Println(out)
}
