package service

import (
	m "api/pkg/models"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	er "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

const (
	queueName = "wallet"
	msg       = `{"to":"0xAbC1234567", "value":"0x1bc16d674ec80000"}`
)

// create a aws session, it will be use in tests
func makeSetupSession() (*session.Session, error) {
	awsSession := &session.Session{}
	awsUser := &m.UserConfig{}
	err := awsUser.NewUserConfig()
	if err != nil {
		log.Println(er.Wrap(err, "INFO: var_env_error"))
		return nil, er.New("var_env_error")
	}

	// create an AWS session
	awsSession, err = makeSession(awsUser)
	if err != nil {
		log.Println(er.Wrap(err, "INFO: session_error"))
		return nil, er.New("session_error")
	}

	return awsSession, nil

}

func TestGenerateAwsSessionSucess(t *testing.T) {
	awsSession, err := makeSetupSession()
	assert.Nil(t, err)
	assert.NotNil(t, awsSession)
	assert.NotNil(t, awsSession.Config.Credentials)
	assert.NotNil(t, awsSession.Config)
}

func TestGetQueueURLSucess(t *testing.T) {
	awsSession, err := makeSetupSession()
	assert.Nil(t, err)
	assert.NotNil(t, awsSession)

	queueUrl, err := getQueueURL(awsSession, queueName)
	assert.Nil(t, err)
	assert.NotNil(t, queueUrl)
}

func TestGetQueueURLError(t *testing.T) {
	awsSession, err := makeSetupSession()
	assert.Nil(t, err)
	assert.NotNil(t, awsSession)

	// call getQueueURL with "" to induce error
	queueUrl, err := getQueueURL(awsSession, "")
	assert.NotNil(t, err)
	assert.Nil(t, queueUrl)
}

func TestSendMessageSucess(t *testing.T) {
	awsSession, err := makeSetupSession()
	assert.Nil(t, err)
	assert.NotNil(t, awsSession)

	queueUrl, err := getQueueURL(awsSession, queueName)
	assert.Nil(t, err)
	assert.NotNil(t, queueUrl)

	err = sendMessage(awsSession, *queueUrl.QueueUrl, msg)
	assert.Nil(t, err)
}

func TestSendMessageError(t *testing.T) {
	awsSession, err := makeSetupSession()
	assert.Nil(t, err)
	assert.NotNil(t, awsSession)

	queueUrl, err := getQueueURL(awsSession, queueName)
	assert.Nil(t, err)
	assert.NotNil(t, queueUrl)

	// call sendMessage with "" to induce error
	err = sendMessage(awsSession, *queueUrl.QueueUrl, "")
	assert.NotNil(t, err)
}

// You can use this test to send msg to SQS (that call a lambda func)
func TestSendMsgWithSQSSenderSucess(t *testing.T) {
	err := SQSSender(msg)
	assert.Nil(t, err)
}
