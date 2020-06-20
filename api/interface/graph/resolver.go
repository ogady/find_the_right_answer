// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
package graph

import (
	"context"

	model "github.com/ogady/find_the_right_answer/api/domain/model"
	"github.com/ogady/find_the_right_answer/api/interface/adapter"
	"github.com/ogady/find_the_right_answer/api/interface/graph/generated"
	"github.com/ogady/find_the_right_answer/api/usecase"
)

type Resolver struct{}

func (r *mutationResolver) AddTopicPiece(ctx context.Context, input adapter.TopicPiece) (*model.TopicPiece, error) {

	inputTopicPiece := model.TopicPiece{
		TopicPiece: input.TopicPiece,
	}

	//操作のタイミングを追跡するために子スパンを作成します。
	// addTopicPieceSpan, _ := tracer.StartSpanFromContext(ctx, "AddTopicPiece")
	u := usecase.NewAddTopicPieceUsecase()
	topicPiece, err := u.AddTopicPiece(&inputTopicPiece)
	if err != nil {
		return nil, err
	}
	// addTopicPieceSpan.Finish()
	return topicPiece, nil
}

func (r *mutationResolver) LikeTopic(ctx context.Context, input adapter.Topic) (*model.Topic, error) {

	inputTopic := model.Topic{
		StartChar: model.StartChar{
			StartChar: input.StartChar.StartChar,
		},
		TopicPiece: model.TopicPiece{
			TopicPiece: input.TopicPiece.TopicPiece,
		},
		NumOfLikes: model.NumOfLikes{
			NumOfLikes: input.NumOfLikes.NumOfLikes,
		},
	}

	//操作のタイミングを追跡するために子スパンを作成します。
	// likeTopicSpan, _ := tracer.StartSpanFromContext(ctx, "LikeTopic")
	u := usecase.NewLikeTopicUsecase()

	topic, err := u.LikeTopic(inputTopic)
	if err != nil {
		return &topic, err
	}
	// likeTopicSpan.Finish()
	return &topic, nil
}

func (r *queryResolver) Topic(ctx context.Context) (*model.Topic, error) {
	//操作のタイミングを追跡するために子スパンを作成します。
	// createTopicSpan, _ := tracer.StartSpanFromContext(ctx, "createTopic")

	u := usecase.NewCreateTopicUsecase()
	topic, err := u.CreateTopic()
	if err != nil {
		return nil, err
	}
	// createTopicSpan.Finish()
	return &topic, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
