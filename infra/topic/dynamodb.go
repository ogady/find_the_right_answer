package infra

import (
	"github.com/guregu/dynamo"
	"github.com/ogady/find_the_right_answer/domain/model"
	"github.com/ogady/find_the_right_answer/domain/repository"
	awsInfra "github.com/ogady/find_the_right_answer/infra/aws"
)

type TopicRepoImpl struct {
	dynamoDB *dynamo.DB
	table    *dynamo.Table
}

func NewTopicRepoImpl() repository.TopicRepository {

	db := awsInfra.NewDynamoDBConn()

	table := db.Table("topic")
	topicRepoImpl := &TopicRepoImpl{
		dynamoDB: db,
		table:    &table,
	}

	return topicRepoImpl
}

func (r *TopicRepoImpl) Save(topic *model.Topic) error {

	var err error
	t := model.Topic{
		StartChar:  topic.StartChar,
		TopicPiece: topic.TopicPiece,
	}
	err = r.table.Put(t).If("attribute_not_exists(StartChar)").Run()
	return err
}

func (r *TopicRepoImpl) FindAll() ([]model.Topic, error) {
	var topics []model.Topic
	var err error

	err = r.table.Scan().All(&topics)
	if err != nil {
		return nil, err
	}
	return topics, err
}

func (r *TopicRepoImpl) Find(topicID string) (model.Topic, error) {
	var topic model.Topic
	var err error

	return topic, err
}
