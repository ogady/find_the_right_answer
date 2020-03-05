package infra

import (
	"github.com/ogady/find_the_right_answer/domain/model"
	"github.com/ogady/find_the_right_answer/domain/repository"
)

type TopicRepoImplMock struct {
}

func NewTopicRepoImplMock() repository.TopicRepository {
	topicRepoImplMock := &TopicRepoImplMock{}

	return topicRepoImplMock
}

type dynamoTopicMock struct {
	StartChar  string
	TopicPiece string
	NumOfLikes int
}

func (r *TopicRepoImplMock) Save(topic *model.Topic) error {
	var err error

	return err
}

func (r *TopicRepoImplMock) FindAll() ([]model.Topic, error) {
	var ts []model.Topic
	var err error

	return ts, err
}

func (r *TopicRepoImplMock) Find(topicID string) (model.Topic, error) {
	var topic model.Topic
	var err error

	return topic, err
}
