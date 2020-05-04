package usecase

import (
	"github.com/ogady/find_the_right_answer/api/domain/model"
	"github.com/ogady/find_the_right_answer/api/domain/repository"
	topicPieceInfra "github.com/ogady/find_the_right_answer/api/infra/topicPiece"
)

type addTopicPieceUsecase struct {
	topicPieceRepo repository.TopicPieceRepository
}

// NewAddTopicPieceUsecase -
func NewAddTopicPieceUsecase() addTopicPieceUsecase {
	tpRepo := topicPieceInfra.NewTopicPieceRepoImpl()

	return addTopicPieceUsecase{topicPieceRepo: tpRepo}
}

func (r *addTopicPieceUsecase) AddTopicPiece(newTopicPiece *model.TopicPiece) error {
	var err error

	topicPiece := model.TopicPiece{
		TopicPiece: newTopicPiece.TopicPiece,
	}

	err = r.topicPieceRepo.Save(&topicPiece)

	return err
}
