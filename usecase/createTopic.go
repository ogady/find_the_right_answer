package usecase

import (
	"github.com/ogady/find_the_right_answer/domain/model"
	startCharInfra "github.com/ogady/find_the_right_answer/infra/startChar"
	topicInfra "github.com/ogady/find_the_right_answer/infra/topic"
	topicPieceInfra "github.com/ogady/find_the_right_answer/infra/topicPiece"
)

func CreateTopicUsecase() (model.Topic, error) {
	var topic model.Topic
	var err error
	stRepo := startCharInfra.NewStartCharRepoImpl()
	tpRepo := topicPieceInfra.NewTopicPieceRepoImpl()
	tRepo := topicInfra.NewTopicRepoImpl()

	topicPiece, err := tpRepo.FindRandom()
	if err != nil {
		return topic, err
	}

	topic = model.Topic{
		StartChar:  stRepo.FindRandom(),
		TopicPiece: topicPiece,
	}

	err = tRepo.Save(&topic)

	if err != nil {
		return topic, err
	}

	return topic, err
}
