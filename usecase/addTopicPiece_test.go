package usecase

import (
	"testing"

	"github.com/ogady/find_the_right_answer/domain/model"
	"github.com/ogady/find_the_right_answer/domain/repository"
	topicPieceInfra "github.com/ogady/find_the_right_answer/infra/topicPiece"
)

func Test_addTopicPieceUsecase_AddTopicPiece(t *testing.T) {
	type fields struct {
		topicPieceRepo repository.TopicPieceRepository
	}
	type args struct {
		newTopicPiece *model.TopicPiece
	}

	mockImpl := topicPieceInfra.NewAddTopicPieceRepoImplMock()
	newTopicPiece := model.TopicPiece{TopicPiece: "正常系は？"}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "normal_test",
			fields:  fields{topicPieceRepo: mockImpl},
			args:    args{newTopicPiece: &newTopicPiece},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &addTopicPieceUsecase{
				topicPieceRepo: tt.fields.topicPieceRepo,
			}
			if err := r.AddTopicPiece(tt.args.newTopicPiece); (err != nil) != tt.wantErr {
				t.Errorf("addTopicPieceUsecase.AddTopicPiece() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
