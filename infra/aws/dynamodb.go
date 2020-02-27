package infra

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/ogady/find_the_right_answer/config"
)

func NewDynamoDBConn() *dynamo.DB {
	var conf *config.DynamoDBConf
	db := dynamo.New(session.New(), &aws.Config{
		Region:      aws.String(conf.Region),
		Endpoint:    aws.String(conf.Endpoint),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	})
	return db
}
