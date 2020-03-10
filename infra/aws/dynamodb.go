package infra

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/ogady/find_the_right_answer/config"
)

func NewDynamoDBConn() *dynamo.DB {

	var db *dynamo.DB

	env := os.Getenv("ENV")
	if env == "" {
		panic("failed to get application mode, check whether APP_MODE is set.")
	}

	if env == "prod" {
		db = dynamo.New(session.New(), &aws.Config{
			Region: aws.String(config.DynamoDB.Region),
		})
	} else {
		db = dynamo.New(session.New(), &aws.Config{
			Region:   aws.String(config.DynamoDB.Region),
			Endpoint: aws.String(config.DynamoDB.Endpoint),
			Credentials: credentials.NewStaticCredentials(
				config.DynamoDB.AccessKey,
				config.DynamoDB.SecretKey,
				config.DynamoDB.SessionToken,
			),
		})
	}

	return db
}
