package testify

import "github.com/aws/aws-sdk-go/service/sns"

type SnsI interface {
	Publish(*sns.PublishInput) (*sns.PublishOutput, error)
}
