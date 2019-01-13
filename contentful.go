package main

import (
	"fmt"
	"log"

	"github.com/contentful-labs/contentful-go"
)

func main() {
	token := "CFPAT-12b35de2da84f95bba7b1dca42b5a8c953e30ef781a19835627dfaeb5a0cc9ad" // observe your CMA token from Contentful's web page
	cma := contentful.NewCMA(token)

	space, err := cma.Spaces.Get("jw19dzxsuqrn")
	if err != nil {
		log.Fatal(err)
	}

	collection := cma.ContentTypes.List(space.Sys.ID)
	collection, err = collection.Next()
	if err != nil {
		log.Fatal(err)
	}

	for _, contentType := range collection.ToContentType() {
		fmt.Println(contentType.Name, contentType.Description)
	}
}
