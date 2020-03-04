package usecase

import (
	"reflect"
	"testing"

	"github.com/ogady/find_the_right_answer/domain/model"
	"github.com/ogady/find_the_right_answer/domain/repository"
	startCharInfra "github.com/ogady/find_the_right_answer/infra/startChar"
	topicInfra "github.com/ogady/find_the_right_answer/infra/topic"
	topicPieceInfra "github.com/ogady/find_the_right_answer/infra/topicPiece"
)

func Test_createTopicUsecase_CreateTopic(t *testing.T) {
	type fields struct {
		topicPieceRepo repository.TopicPieceRepository
		startChartRepo repository.StartCharRepository
		topicRepo      repository.TopicRepository
	}

	topicPieceMockImpl := topicPieceInfra.NewAddTopicPieceRepoImplMock()
	topicMockImpl := topicInfra.NewTopicRepoImplMock()
	startCharMockImpl := startCharInfra.NewStartCharRepoImplMock()

	tests := []struct {
		name    string
		fields  fields
		want    model.Topic
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "normal_test",
			fields: fields{
				topicPieceRepo: topicPieceMockImpl,
				startChartRepo: startCharMockImpl,
				topicRepo:      topicMockImpl,
			},
			want: model.Topic{
				StartChar:  model.StartChar{StartChar: "あ"},
				TopicPiece: model.TopicPiece{TopicPiece: "人生を変える本は？"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &createTopicUsecase{
				topicPieceRepo: tt.fields.topicPieceRepo,
				startChartRepo: tt.fields.startChartRepo,
				topicRepo:      tt.fields.topicRepo,
			}
			got, err := r.CreateTopic()
			if (err != nil) != tt.wantErr {
				t.Errorf("createTopicUsecase.CreateTopic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createTopicUsecase.CreateTopic() = %v, want %v", got, tt.want)
			}
		})
	}
}
