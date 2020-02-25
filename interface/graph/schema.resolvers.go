// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"

	model1 "github.com/ogady/find_the_right_answer/domain/model"
	"github.com/ogady/find_the_right_answer/interface/graph/generated"
	"github.com/ogady/find_the_right_answer/interface/graph/model"
	"github.com/ogady/find_the_right_answer/usecase"
)

func (r *mutationResolver) AddTopicPiece(ctx context.Context, input model.NewTopicPiece) (*model1.TopicPiece, error) {
	err := usecase.AddTopicPieceUsecase(&input)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *queryResolver) Topic(ctx context.Context) (*model1.Topic, error) {
	topic, err := usecase.CreateTopicUsecase()
	if err != nil {
		return nil, err
	}

	return &topic, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
