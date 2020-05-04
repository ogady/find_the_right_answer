package usecase

import "github.com/ogady/find_the_right_answer/api/domain/model"

type likeTopicUsecase struct {
}

// NewLikeTopicUsecase -
func NewLikeTopicUsecase() createTopicUsecase {

	return createTopicUsecase{}
}

func (r *likeTopicUsecase) LikeTopic(topic model.Topic) error {
	var err error
	return err
}
