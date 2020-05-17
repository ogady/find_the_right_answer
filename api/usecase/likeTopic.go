package usecase

import (
	"github.com/ogady/find_the_right_answer/api/domain/model"
	"github.com/ogady/find_the_right_answer/api/domain/repository"
	topicInfra "github.com/ogady/find_the_right_answer/api/infra/topic"
)

type likeTopicUsecase struct {
	topicRepo repository.TopicRepository
}

// NewLikeTopicUsecase -
func NewLikeTopicUsecase() likeTopicUsecase {
	tRepo := topicInfra.NewTopicRepoImpl()
	ltUsecase := likeTopicUsecase{
		topicRepo: tRepo,
	}
	return ltUsecase

}

// TODO 元のNumOfLikeを取得
func (r *likeTopicUsecase) LikeTopic(topic model.Topic) (model.Topic, error) {
	var err error

	incrementedTopic := model.Topic{
		StartChar:  topic.StartChar,
		TopicPiece: topic.TopicPiece,
		NumOfLikes: model.NumOfLikes{
			NumOfLikes: topic.NumOfLikes.NumOfLikes + 1,
		},
	}

	topic, err = r.topicRepo.UpdateTopicNumOfLike(incrementedTopic)
	if err != nil {
		return topic, err
	}

	return topic, err
}
