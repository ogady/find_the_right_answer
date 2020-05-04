package repository

import (
	"github.com/ogady/find_the_right_answer/api/domain/model"
)

type TopicRepository interface {
	Save(*model.Topic) error
	FindAll() ([]model.Topic, error)
	Find(string) (model.Topic, error)
}
