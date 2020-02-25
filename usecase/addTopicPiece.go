package usecase

import (
	"github.com/ogady/find_the_right_answer/domain/model"
	inModel "github.com/ogady/find_the_right_answer/interface/graph/model"
	topicPieceInfra "github.com/ogady/find_the_right_answer/infra/topicPiece"
)

func AddTopicPieceUsecase(newTopicPiece *inModel.NewTopicPiece) error {
	var err error

	topicPiece := model.TopicPiece{
		TopicPiece: newTopicPiece.TopicPiece,
	}

	tpRepo := topicPieceInfra.NewTopicPieceRepoImpl()
	err = tpRepo.Save(&topicPiece)

	return err
}
