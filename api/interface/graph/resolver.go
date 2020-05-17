// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
package graph

import (
	"context"

	model "github.com/ogady/find_the_right_answer/api/domain/model"
	"github.com/ogady/find_the_right_answer/api/interface/graph/generated"
	"github.com/ogady/find_the_right_answer/api/usecase"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type Resolver struct{}

func (r *mutationResolver) AddTopicPiece(ctx context.Context, input model.TopicPiece) (*model.TopicPiece, error) {
	//操作のタイミングを追跡するために子スパンを作成します。
	addTopicPieceSpan, _ := tracer.StartSpanFromContext(ctx, "AddTopicPiece")
	u := usecase.NewAddTopicPieceUsecase()
	err := u.AddTopicPiece(&input)
	if err != nil {
		return nil, err
	}
	addTopicPieceSpan.Finish()
	return nil, nil
}

func (r *mutationResolver) LikeTopic(ctx context.Context, input model.Topic) (*model.Topic, error) {
	//操作のタイミングを追跡するために子スパンを作成します。
	likeTopicSpan, _ := tracer.StartSpanFromContext(ctx, "LikeTopic")
	u := usecase.NewLikeTopicUsecase()
	topic, err := u.LikeTopic(input)
	if err != nil {
		return &topic, err
	}
	likeTopicSpan.Finish()
	return &topic, nil
}

func (r *queryResolver) Topic(ctx context.Context) (*model.Topic, error) {
	//操作のタイミングを追跡するために子スパンを作成します。
	createTopicSpan, _ := tracer.StartSpanFromContext(ctx, "createTopic")

	u := usecase.NewCreateTopicUsecase()
	topic, err := u.CreateTopic()
	if err != nil {
		return nil, err
	}
	createTopicSpan.Finish()
	return &topic, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
