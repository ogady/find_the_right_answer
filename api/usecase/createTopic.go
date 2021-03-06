package usecase

import (
	"github.com/ogady/find_the_right_answer/api/domain/model"
	"github.com/ogady/find_the_right_answer/api/domain/repository"
	startCharInfra "github.com/ogady/find_the_right_answer/api/infra/startChar"
	topicInfra "github.com/ogady/find_the_right_answer/api/infra/topic"
	topicPieceInfra "github.com/ogady/find_the_right_answer/api/infra/topicPiece"
)

type createTopicUsecase struct {
	topicPieceRepo repository.TopicPieceRepository
	startChartRepo repository.StartCharRepository
	topicRepo      repository.TopicRepository
}

// NewCreateTopicUsecase -
func NewCreateTopicUsecase() createTopicUsecase {
	tpRepo := topicPieceInfra.NewTopicPieceRepoImpl()
	stRepo := startCharInfra.NewStartCharRepoImpl()
	tRepo := topicInfra.NewTopicRepoImpl()
	ctUsecase := createTopicUsecase{
		topicPieceRepo: tpRepo,
		startChartRepo: stRepo,
		topicRepo:      tRepo,
	}
	return ctUsecase
}

func (r *createTopicUsecase) CreateTopic() (model.Topic, error) {
	var topic model.Topic
	var err error

	startChar := r.startChartRepo.FindRandom()

	topicPiece, err := r.topicPieceRepo.FindRandom()
	if err != nil {
		return topic, err
	}

	topic = model.Topic{
		StartChar:  startChar,
		TopicPiece: topicPiece,
	}

	err = r.topicRepo.Save(&topic)
	if err != nil {
		return topic, err
	}

	numOfLikes, err := r.topicRepo.FetchOnlyNumOfLikeByTopic(topic)
	if err != nil {
		return topic, err
	}

	topic.NumOfLikes = numOfLikes

	return topic, nil
}
