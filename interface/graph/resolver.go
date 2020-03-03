// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
package graph

import (
	"context"

	model "github.com/ogady/find_the_right_answer/domain/model"
	"github.com/ogady/find_the_right_answer/interface/graph/generated"
	"github.com/ogady/find_the_right_answer/usecase"
)

type Resolver struct{}

func (r *mutationResolver) AddTopicPiece(ctx context.Context, input model.TopicPiece) (*model.TopicPiece, error) {
	u := usecase.NewAddTopicPieceUsecase()
	err := u.AddTopicPiece(&input)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *queryResolver) Topic(ctx context.Context) (*model.Topic, error) {
	u := usecase.NewCreateTopicUsecase()
	topic, err := u.CreateTopic()
	if err != nil {
		return nil, err
	}

	return &topic, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
