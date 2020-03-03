package usecase

import (
	"fmt"
	"strings"

	"github.com/ogady/find_the_right_answer/domain/model"
	"github.com/ogady/find_the_right_answer/domain/repository"
	startCharInfra "github.com/ogady/find_the_right_answer/infra/startChar"
	topicInfra "github.com/ogady/find_the_right_answer/infra/topic"
	topicPieceInfra "github.com/ogady/find_the_right_answer/infra/topicPiece"
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
	fmt.Println(startChar)

	topicPiece, err := r.topicPieceRepo.FindRandom()
	if err != nil {
		return topic, err
	}

	fmt.Println(topicPiece)

	topic = model.Topic{
		StartChar:  startChar,
		TopicPiece: topicPiece,
	}

	err = r.topicRepo.Save(&topic)

	if err != nil {
		switch {
		// 重複エラーは問題なくレスポンスする
		case strings.Contains(err.Error(), "ConditionalCheckFailedException"):
			return topic, nil
		default:
			return topic, err
		}
	}
	return topic, nil
}
