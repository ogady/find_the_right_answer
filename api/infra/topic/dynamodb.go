package infra

import (
	"strings"

	"github.com/guregu/dynamo"
	"github.com/ogady/find_the_right_answer/api/domain/model"
	"github.com/ogady/find_the_right_answer/api/domain/repository"
	awsInfra "github.com/ogady/find_the_right_answer/api/infra/aws"
)

type TopicRepoImpl struct {
	dynamoDB *dynamo.DB
	table    *dynamo.Table
}

// TODO Connection破棄
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

// Save - TopicをDynamoDBのTopicテーブルに保存する。
func (r *TopicRepoImpl) Save(topic *model.Topic) error {

	var err error
	dt := dynamoTopic{
		StartChar:  topic.StartChar.StartChar,
		TopicPiece: topic.TopicPiece.TopicPiece,
		NumOfLikes: topic.NumOfLikes.NumOfLikes,
	}

	err = r.table.Put(&dt).If("attribute_not_exists(StartChar)").Run()
	if err != nil {
		switch {
		// 重複エラーは問題なくレスポンスする
		case strings.Contains(err.Error(), "ConditionalCheckFailedException"):
			return nil

		default:
			return err
		}
	}

	return nil
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
			NumOfLikes: model.NumOfLikes{NumOfLikes: v.NumOfLikes},
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

func (r *TopicRepoImpl) FetchOnlyNumOfLikeByTopic(topic model.Topic) (model.NumOfLikes, error) {
	var numOfLikes model.NumOfLikes
	var err error

	err = r.table.Get("StartChar", topic.StartChar.StartChar).Range("TopicPiece", dynamo.Equal, topic.TopicPiece.TopicPiece).One(&numOfLikes)

	if err != nil {
		return numOfLikes, err
	}

	return numOfLikes, nil
}

func (r *TopicRepoImpl) UpdateTopicNumOfLike(topic model.Topic) (model.Topic, error) {
	var resultTopic dynamoTopic
	var err error

	err = r.table.Update("StartChar", topic.StartChar.StartChar).Range("TopicPiece", topic.TopicPiece.TopicPiece).Set("NumOfLikes", &topic.NumOfLikes.NumOfLikes).Value(&resultTopic)
	if err != nil {
		return topic, err
	}

	topic.StartChar.StartChar = resultTopic.StartChar
	topic.TopicPiece.TopicPiece = resultTopic.TopicPiece
	topic.NumOfLikes.NumOfLikes = resultTopic.NumOfLikes

	return topic, nil
}
