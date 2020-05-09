package repository

import (
	"github.com/ogady/find_the_right_answer/api/domain/model"
)

type TopicRepository interface {
	// Save - Topicを保存する。
	Save(*model.Topic) error
	FindAll() ([]model.Topic, error)
	Find(string) (model.Topic, error)
	// FetchOnlyNumOfLikeByTopic - TopicのNumofLikeのみ取得する。
	FetchOnlyNumOfLikeByTopic(model.Topic) (int, error)
	UpdateTopicNumOfLike(model.Topic) (model.Topic, error)
}
