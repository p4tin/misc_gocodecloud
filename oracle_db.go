package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	_ "github.com/mattn/go-oci8"
	"time"
)

type Tile struct {
	Create_date   string `json:"create_date"`
	Publish_state int    `json:"publish_state"`
	Tile_id       string `json:"tile_id"`
	Visible       int    `json:"visible"`
	Category_id   string `json:"category_id"`
	Product_id    string `json:"product_id"`
}

func main() {
	n := time.Now().UnixNano()
	db, err := sql.Open("oci8", "MCPDEV/MCPDEV@10.9.8.10:1531/ATGUTL")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	fmt.Printf("1. %d\n", (time.Now().UnixNano() - n))

	if err = db.Ping(); err != nil {
		fmt.Printf("Error connecting to the database: %s\n", err)
		return
	}
	n = time.Now().UnixNano()
	rows, err := db.Query("select CREATE_DATE, PUBLISH_STATE, TILE_ID, VISIBLE, CATEGORY_ID, PRODUCT_ID from UC_TILE where PUBLISH_STATE = 3")
	if err != nil {
		fmt.Println("Error fetching Tiles")
		fmt.Println(err)
		return
	}
	defer rows.Close()
	fmt.Printf("2. %d\n", (time.Now().UnixNano() - n))
	n = time.Now().UnixNano()
	results := []Tile{}
	for rows.Next() {
		t := Tile{}
		rows.Scan(&t.Create_date, &t.Publish_state, &t.Tile_id, &t.Visible, &t.Category_id, &t.Product_id)
		//fmt.Printf("%+v\n", t)
		publish(t)
		results = append(results, t)
	}
	fmt.Printf("3. %d\n", (time.Now().UnixNano() - n))
	fmt.Printf("%d tiles selected...\n", len(results))
}

func publish(t Tile) {
	msg, err := json.Marshal(&t)
	if err != nil {
		fmt.Println("Marshal Err:", err.Error())
		return
	}
	svc := sns.New(session.New(), &aws.Config{Endpoint: aws.String("http://localhost:4100"), Region: aws.String("local")})
	params := &sns.PublishInput{
		Message:  aws.String(string(msg)),
		TopicArn: aws.String("arn:aws:sns:local:000000000000:A15_DEV_TILE_FACT"),
	}

	_, err = svc.Publish(params)
	if err != nil {
		fmt.Println("Publish Err:", err.Error())
		return
	}
}
