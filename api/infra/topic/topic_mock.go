package infra

import (
	"github.com/ogady/find_the_right_answer/api/domain/model"
	"github.com/ogady/find_the_right_answer/api/domain/repository"
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
	var topics []model.Topic
	var err error

	return topics, err
}

func (r *TopicRepoImplMock) Find(topicID string) (model.Topic, error) {
	var topic model.Topic
	var err error

	return topic, err
}

func (r *TopicRepoImplMock) FetchOnlyNumOfLikeByTopic(topic model.Topic) (int, error) {
	//var topic model.Topic
	var numOfLikes int
	var err error

	return numOfLikes, err
}

func (r *TopicRepoImplMock) UpdateTopicNumOfLike(topic model.Topic) (model.Topic, error) {
	var resultTopic model.Topic
	var err error

	return resultTopic, err
}
