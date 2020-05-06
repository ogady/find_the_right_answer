package usecase

import (
	"strings"

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
		switch {
		// 重複エラーは問題なくレスポンスする
		case strings.Contains(err.Error(), "ConditionalCheckFailedException"):
			// TODO breakする処理に変える
			return topic, nil

		default:
			return topic, err
		}
	}

	// TODO TopicのNumOfLikesのみを取ってくる処理を追加
	// 取ってきたTopic := TopicのNumOfLikesのみを取ってくる()
	// topic.NumOfLikes = 取ってきたTopic.NumOfLikes

	return topic, nil
}
