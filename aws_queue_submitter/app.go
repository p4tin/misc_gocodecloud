package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/remeh/sizedwaitgroup"
)

type WorkMsg struct {
	QueueName string
	Message   string
}

var sqsClient *sqs.SQS

func main() {
	fmt.Println(time.Now())
	defer fmt.Println(time.Now())
	args := os.Args[1:]
	if len(args) < 2 {
		panic("Not enough arguments, please provide <filename> <list of queues command delimited>")
	}
	filename := args[0]
	queues := strings.Split(args[1], ",")
	fmt.Println(queues)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	awsConfig := &aws.Config{Region: aws.String("us-east-1")}
	sqsClient = sqs.New(session.New(), awsConfig)

	scanner := bufio.NewScanner(file)
	var messages []string
	for scanner.Scan() {
		messages = append(messages, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	swg := sizedwaitgroup.New(100)

	queueIndex := 0
	for x := 0; x < len(messages); x++ {
		swg.Add()
		go func(x, queueIndex int) {
			defer swg.Done()
			sendMessage(queues[queueIndex], messages[x])
		}(x, queueIndex)
		queueIndex++
		if queueIndex >= len(queues) {
			queueIndex = 0
		}
	}

	swg.Wait()
}

func sendMessage(queueName, message string) {
	sendMessageInput := &sqs.SendMessageInput{
		MessageBody: aws.String(message),
		QueueUrl:    aws.String(fmt.Sprintf("https://sqs.us-east-1.amazonaws.com/478989820108/%s", queueName)),
	}
	_, err := sqsClient.SendMessage(sendMessageInput)
	if err != nil {
		fmt.Println(message, err)
	}
}
