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

type dynamoTopic struct {
	StartChar  string
	TopicPiece string
	NumOfLikes int
}

func (r *TopicRepoImpl) Save(topic *model.Topic) error {

	var err error
	dt := dynamoTopic{
		StartChar:  topic.StartChar.StartChar,
		TopicPiece: topic.TopicPiece.TopicPiece,
		NumOfLikes: topic.NumOfLikes,
	}

	err = r.table.Put(&dt).If("attribute_not_exists(StartChar)").Run()

	return err
}

func (r *TopicRepoImpl) FindAll() ([]model.Topic, error) {
	var dts []dynamoTopic
	var ts []model.Topic
	var err error

	err = r.table.Scan().All(&dts)
	if err != nil {
		return nil, err
	}

	for _, v := range dts {
		t := model.Topic{
			StartChar:  model.StartChar{StartChar: v.StartChar},
			TopicPiece: model.TopicPiece{TopicPiece: v.TopicPiece},
			NumOfLikes: v.NumOfLikes,
		}
		ts = append(ts, t)
	}
	return ts, err
}

func (r *TopicRepoImpl) Find(topicID string) (model.Topic, error) {
	var topic model.Topic
	var err error

	return topic, err
}
