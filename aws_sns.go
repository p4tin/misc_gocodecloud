package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {
	//Create a session object to talk to SNS (also make sure you have your key and secret setup in your .aws/credentials file)
	svc := sns.New(session.New())
	// params will be sent to the publish call included here is the bare minimum params to send a message.
	params := &sns.PublishInput{
		Message: aws.String("message"), // This is the message itself (can be XML / JSON / Text - anything you want)
		TopicArn: aws.String("arn:aws:sns:us-east-1:478989820108:MCP_DEV_CATEGORY_EX_TOPIC"),  //Get this from the Topic in the AWS console.
	}

	resp, err := svc.Publish(params)   //Call to puclish the message

	if err != nil {                    //Check for errors
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
