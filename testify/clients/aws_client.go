package clients

import (
	"fmt"
	"misc/testify"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// arn:aws:sns:us-east-1:478989820108:A15_DEV_PAUL_DELTA_EVENT_TOPIC
// arn:aws:sns:us-east-1:478989820108:A15_DEV_PAUL_PRODUCT_UPDATE_EVENTS_TOPIC

func GetAwsClient(region string) testify.SnsI {
	return sns.New(session.New(), &aws.Config{Region: aws.String(region)})
}

var (
	GetAwsClientProc = GetAwsClient
)

type SnsClient struct {
	SnsService testify.SnsI
}

func CreateAwsClient() SnsClient {
	snsClient := SnsClient{
		SnsService: GetAwsClientProc("us-east-1"),
	}

	return snsClient
}

func (client SnsClient) ReceiveMessage(interface{}) []string {
	return make([]string, 0)
}

func (client SnsClient) PostMessage(topicArn string, msg string) {
	if client.SnsService == nil {
		client.SnsService = GetAwsClientProc("us-east-1")
	}
	params := &sns.PublishInput{
		Message:  aws.String(msg),      // This is the message itself (can be XML / JSON / Text - anything you want)
		TopicArn: aws.String(topicArn), // Get this from the Topic in the AWS console.
	}

	_, err := client.SnsService.Publish(params) //Call to publish the message

	if err != nil {
		fmt.Println(err)
		return
	}
}

func (client SnsClient) PutMessage(dest string, msg string) {

}
