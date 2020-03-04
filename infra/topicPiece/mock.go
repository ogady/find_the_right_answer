package infra

import (
	"github.com/ogady/find_the_right_answer/domain/model"
	"github.com/ogady/find_the_right_answer/domain/repository"
)

type addTopicPieceRepoImplMock struct {
}

func NewAddTopicPieceRepoImplMock() repository.TopicPieceRepository {

	mock := addTopicPieceRepoImplMock{}

	return &mock
}

func (r *addTopicPieceRepoImplMock) Save(topicPiece *model.TopicPiece) error {
	var err error
	return err
}

func (r *addTopicPieceRepoImplMock) FindAll() ([]model.TopicPiece, error) {
	var topicPieces []model.TopicPiece
	var err error

	return topicPieces, err
}

func (r *addTopicPieceRepoImplMock) Find(str string) (model.TopicPiece, error) {
	var topicPiece model.TopicPiece
	var err error

	return topicPiece, err
}

func (r *addTopicPieceRepoImplMock) FindRandom() (model.TopicPiece, error) {
	var topicPiece model.TopicPiece
	var err error
	topicPiece = model.TopicPiece{TopicPiece: "人生を変える本は？"}
	return topicPiece, err
}
