package usecase

import "github.com/ogady/find_the_right_answer/api/domain/model"

type likeTopicUsecase struct {
}

// NewLikeTopicUsecase -
func NewLikeTopicUsecase() createTopicUsecase {

	return createTopicUsecase{}
}

func (r *likeTopicUsecase) LikeTopic(topic model.Topic) (model.Topic, error) {
	var err error
	var incrementedTopic model.Topic

	return incrementedTopic, err
}
