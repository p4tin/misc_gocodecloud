package main

import (
	"encoding/json"
	"flag"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Environment struct {
	Awsegion    string
	OrdersQueue string
	ErrorQueue  string
}

var configs = map[string]Environment{
	"dev": {
		Awsegion:    "us-east-1",
		OrdersQueue: "https://sqs.us-east-1.amazonaws.com/478989820108/A15-DEV-ORDER-FULFILLMENT",
		ErrorQueue:  "https://sqs.us-east-1.amazonaws.com/478989820108/A15-DEV-ORDER-FULFILLMENT-ERRORS",
	},
	"staging": {
		Awsegion:    "us-east-1",
		ErrorQueue:  "https://sqs.us-east-1.amazonaws.com/478989820108/A15-STAGE2-ORDER-FULFILLMENT",
		OrdersQueue: "https://sqs.us-east-1.amazonaws.com/478989820108/A15-STAGE2-ORDER-FULFILLMENT-ERRORS",
	},
	"prod": {
		Awsegion:    "us-east-1",
		OrdersQueue: "https://sqs.us-east-1.amazonaws.com/478989820108/A15-PROD1-ORDER-FULFILLMENT",
		ErrorQueue:  "https://sqs.us-east-1.amazonaws.com/478989820108/A15-PROD1-ORDER-FULFILLMENT-ERRORS",
	},
}

func main() {
	var env string
	flag.StringVar(&env, "env", "", "environment to resubmit orders")
	flag.Parse()
	if env == "" {
		log.Fatal("environment was not supplied, please enter one of: 'dev', 'staging' or 'prod'")
	}

	src := configs[env].ErrorQueue
	dest := configs[env].OrdersQueue
	region := configs[env].Awsegion

	log.Printf("environment : %v", env)
	log.Printf("aws region : %v", region)
	log.Printf("source queue : %v", src)
	log.Printf("destination queue : %v", dest)

	config := &aws.Config{
		Region: &region,
	}

	client := sqs.New(session.New(), config)

	maxMessages := int64(10)
	waitTime := int64(0)
	messageAttributeNames := aws.StringSlice([]string{"All"})

	rmin := &sqs.ReceiveMessageInput{
		QueueUrl:              &src,
		MaxNumberOfMessages:   &maxMessages,
		WaitTimeSeconds:       &waitTime,
		MessageAttributeNames: messageAttributeNames,
	}

	// loop as long as there are messages on the queue
	for {
		resp, err := client.ReceiveMessage(rmin)

		if err != nil {
			panic(err)
		}

		if len(resp.Messages) == 0 {
			log.Printf("done")
			return
		}

		log.Printf("received %v messages...", len(resp.Messages))

		var wg sync.WaitGroup
		wg.Add(len(resp.Messages))

		for _, m := range resp.Messages {
			go func(m *sqs.Message) {
				defer wg.Done()

				// m.Body to map[string]string
				var orderSubmitMessage map[string]interface{}
				err := json.Unmarshal([]byte(*m.Body), &orderSubmitMessage)
				if err != nil {
					log.Printf("ERROR Unmarshaling message: %v", err)
					return
				}
				log.Printf("Processing Cart %s, %s, %s\n", orderSubmitMessage["cartId"], orderSubmitMessage["profileId"], orderSubmitMessage["dataCenterId"])
				// Remove "retries"
				delete(orderSubmitMessage, "retries")

				// map[string]string to json
				jsonBytes, err := json.Marshal(orderSubmitMessage)
				if err != nil {
					log.Printf("ERROR Unmarshaling message: %v", err)
					return
				}
				jsonStr := string(jsonBytes)

				// write the message to the destination queue
				smi := sqs.SendMessageInput{
					MessageAttributes: m.MessageAttributes,
					MessageBody:       &jsonStr,
					QueueUrl:          &dest,
				}

				_, err = client.SendMessage(&smi)

				if err != nil {
					log.Printf("ERROR sending message to destination %v", err)
					return
				}

				// message was sent, dequeue from source queue
				dmi := &sqs.DeleteMessageInput{
					QueueUrl:      &src,
					ReceiptHandle: m.ReceiptHandle,
				}

				if _, err := client.DeleteMessage(dmi); err != nil {
					log.Printf("ERROR dequeueing message ID %v : %v",
						*m.ReceiptHandle,
						err)
				}
			}(m)
		}

		// wait for all jobs from this batch...
		wg.Wait()
	}
}
