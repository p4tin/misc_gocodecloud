package clients

import (
	"misc/testify"
	"testing"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/stretchr/testify/mock"
)

type MockSnsClient struct {
	mock.Mock
}

func (msc *MockSnsClient) Publish(input *sns.PublishInput) (output *sns.PublishOutput, err error) {
	args := msc.Called(input)
	return args.Get(0).(*sns.PublishOutput), args.Error(1)
}

func TestName(t *testing.T) {

	// create an instance of our test object
	testObj := &MockSnsClient{}

	// setup expectations
	id := "123"
	testObj.On("Publish", mock.Anything).Return(&sns.PublishOutput{
		MessageId: &id,
	}, nil)

	GetAwsClientProc = func(region string) testify.SnsI {
		return testObj
	}

	ac := CreateAwsClient()
	ac.PostMessage("arn", "Hello World!!!")

	//testObj.AssertCalled(t, "Publish", mock.Anything)
	testObj.AssertCalled(t, "Publish", mock.Anything)
}
