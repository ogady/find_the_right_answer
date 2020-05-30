package usecase

import (
	"reflect"
	"testing"

	"github.com/ogady/find_the_right_answer/api/domain/model"
	"github.com/ogady/find_the_right_answer/api/domain/repository"
	topicPieceInfra "github.com/ogady/find_the_right_answer/api/infra/topicPiece"
)

func Test_addTopicPieceUsecase_AddTopicPiece(t *testing.T) {
	type fields struct {
		topicPieceRepo repository.TopicPieceRepository
	}
	type args struct {
		topicPiece *model.TopicPiece
	}

	mockImpl := topicPieceInfra.NewAddTopicPieceRepoImplMock()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.TopicPiece
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "normal_test",
			fields:  fields{topicPieceRepo: mockImpl},
			args:    args{topicPiece: &model.TopicPiece{TopicPiece: "正常系テスト"}},
			want:    &model.TopicPiece{TopicPiece: "正常系テスト"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &addTopicPieceUsecase{
				topicPieceRepo: tt.fields.topicPieceRepo,
			}
			got, err := r.AddTopicPiece(tt.args.topicPiece)
			if (err != nil) != tt.wantErr {
				t.Errorf("addTopicPieceUsecase.AddTopicPiece() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTopicPieceUsecase.AddTopicPiece() = %v, want %v", got, tt.want)
			}
		})
	}
}
