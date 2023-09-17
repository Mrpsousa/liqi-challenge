package service

import (
	"log"

	m "api/pkg/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	er "github.com/pkg/errors"
)

// retrieves the URL of an SQS queue based on the queue name.
func getQueueURL(sess *session.Session, queue string) (*sqs.GetQueueUrlOutput, error) {
	// create an SQS client with the provided session.
	sqsClient := sqs.New(sess)

	// call the GetQueueUrl function to get the queue URL with the specified name.
	url, err := sqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queue,
	})

	if err != nil {
		return nil, err
	}

	return url, nil
}

// sends a message to an SQS queue using the queue URL
func sendMessage(sess *session.Session, queueUrl string, messageBody string) error {
	// create an SQS client with the provided session.
	sqsClient := sqs.New(sess)

	// send the message to the queue.
	_, err := sqsClient.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    &queueUrl,
		MessageBody: aws.String(messageBody),
	})

	return err
}
func makeSession(user *m.UserConfig) (*session.Session, error) {
	// Create a new AWS session with the provided user configuration
	session, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String(user.MyRegion),
			Credentials: credentials.NewStaticCredentials(
				user.AccessKeyID,
				user.SecretAccessKey,
				"",
			),
		},
	})

	if err != nil {
		// Handle session creation error
		return nil, err
	}
	return session, nil
}

func SQSSender(msg string) error {
	var awsSession = &session.Session{}
	queueName := "wallet"

	// create a new AWS user configuration
	awsUser := &m.UserConfig{}
	err := awsUser.NewUserConfig()
	if err != nil {

		log.Println(er.Wrap(err, "INFO: var_env_error"))
		return er.New("var_env_error")
	}

	// create an AWS session
	awsSession, err = makeSession(awsUser)
	if err != nil {
		log.Println(er.Wrap(err, "INFO: session_error"))
		return er.New("session_error")
	}

	// get the URL of the specified SQS queue
	queueUrl, err := getQueueURL(awsSession, queueName)
	if err != nil {
		log.Println(er.Wrap(err, "INFO: queue_url_error"))
		return er.New("queue_url_error")
	}

	// send the provided message to the SQS queue
	err = sendMessage(awsSession, *queueUrl.QueueUrl, msg)
	if err != nil {
		log.Println(er.Wrap(err, "INFO: erro_send_message_to_queue"))
		return er.New("erro_send_message_to_queue")
	}

	// return nil if the message was sent successfully
	return nil
}
