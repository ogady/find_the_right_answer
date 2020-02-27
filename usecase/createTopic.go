package usecase

import (
	"fmt"
	"strings"

	"github.com/ogady/find_the_right_answer/domain/model"
	startCharInfra "github.com/ogady/find_the_right_answer/infra/startChar"
	topicInfra "github.com/ogady/find_the_right_answer/infra/topic"
	topicPieceInfra "github.com/ogady/find_the_right_answer/infra/topicPiece"
)

func CreateTopicUsecase() (model.Topic, error) {
	var topic model.Topic
	var err error
	stRepo := startCharInfra.NewStartCharRepoImpl()
	tpRepo := topicPieceInfra.NewTopicPieceRepoImpl()
	tRepo := topicInfra.NewTopicRepoImpl()

	startChar := stRepo.FindRandom()
	fmt.Println(startChar)

	topicPiece, err := tpRepo.FindRandom()
	if err != nil {
		return topic, err
	}

	fmt.Println(topicPiece)

	topic = model.Topic{
		StartChar:  startChar,
		TopicPiece: topicPiece,
	}

	err = tRepo.Save(&topic)

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
