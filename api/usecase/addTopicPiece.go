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

func (r *addTopicPieceUsecase) AddTopicPiece(topicPiece *model.TopicPiece) (*model.TopicPiece, error) {
	var err error

	err = r.topicPieceRepo.Save(topicPiece)
	if err != nil {
		return nil, err
	}

	return topicPiece, err
}
