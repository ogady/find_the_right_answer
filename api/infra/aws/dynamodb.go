package infra

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/ogady/find_the_right_answer/api/config"
	awstrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/aws/aws-sdk-go/aws"
)

func NewDynamoDBConn() *dynamo.DB {

	var db *dynamo.DB

	var sess *session.Session

	env := os.Getenv("ENV")
	if env == "" {
		panic("failed to get application mode, check whether APP_MODE is set.")
	}

	if env == "prod" {
		sess = session.Must(session.NewSession(&aws.Config{
			Region: aws.String(config.DynamoDB.Region),
		}))
		sess = awstrace.WrapSession(sess)
		db = dynamo.New(sess)

	} else {

		sess = session.Must(session.NewSession(&aws.Config{
			Region:   aws.String(config.DynamoDB.Region),
			Endpoint: aws.String(config.DynamoDB.Endpoint),
			Credentials: credentials.NewStaticCredentials(
				config.DynamoDB.AccessKey,
				config.DynamoDB.SecretKey,
				config.DynamoDB.SessionToken,
			),
		}))
		sess = awstrace.WrapSession(sess)
		db = dynamo.New(sess)

	}

	return db
}
