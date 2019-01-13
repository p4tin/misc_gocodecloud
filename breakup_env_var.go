package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//export POLLING_PERIOD=10
//export QUEUES_MONITORED="A15-PROD1-AN-GROUPBY-TILE-EVENTS:0,\
//						   A15-PROD1-FP-PRODUCT-UPDATE:0,\
//						   A15-PROD1-AN-GROUPBY-PRODUCT-EVENTS:0,\
//						   A15-PROD1-FP-GROUPBY-PRODUCT-EVENTS:0,\
//						   A15-PROD1-FP-GROUPBY-TILE-EVENTS:0,\
//						   A15-PROD1-UOEU-PRODUCT-UPDATE:0,\
//						   A15-PROD1-UOUS-PRODUCT-UPDATE:0
//						   A15-PROD1-ANUS-PRODUCT-UPDATE:0
//						   A15-PROD1-UO-GROUPBY-TILE-EVENTS:0,\
//						   A15-PROD1-ANEU-PRODUCT-UPDATE:0,\
//						   A15-PROD1-UO-GROUPBY-PRODUCT-EVENTS:0"

type Queue struct {
	Name      string
	MaxDepth  int
	LastDepth int
}

var queueList []Queue

func main() {
	queueList = make([]Queue, 0)
	queues := os.Getenv("QUEUES_MONITORED")
	result := strings.Split(queues, ",")
	for _, res := range result {
		tmpQ := strings.Split(res, ":")
		max, _ := strconv.Atoi(tmpQ[1])
		queue := Queue{
			Name:     tmpQ[0],
			MaxDepth: max,
		}
		queueList = append(queueList, queue)
	}
	fmt.Printf("%+v", queueList)
}
