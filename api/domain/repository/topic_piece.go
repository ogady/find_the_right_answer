package repository

import (
	"github.com/ogady/find_the_right_answer/api/domain/model"
)

type TopicPieceRepository interface {
	Save(*model.TopicPiece) error
	FindAll() ([]model.TopicPiece, error)
	Find(string) (model.TopicPiece, error)
	FindRandom() (model.TopicPiece, error)
}
