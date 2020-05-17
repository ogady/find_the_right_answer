package usecase

import (
	"reflect"
	"testing"

	"github.com/ogady/find_the_right_answer/api/domain/model"
	"github.com/ogady/find_the_right_answer/api/domain/repository"
	topicInfra "github.com/ogady/find_the_right_answer/api/infra/topic"
)

func Test_likeTopicUsecase_LikeTopic(t *testing.T) {
	type fields struct {
		topicRepo repository.TopicRepository
	}
	type args struct {
		topic model.Topic
	}

	topicMockImpl := topicInfra.NewTopicRepoImplMock()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Topic
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "normal_test",
			fields: fields{
				topicRepo: topicMockImpl,
			},
			args: args{topic: model.Topic{
				StartChar:  model.StartChar{StartChar: "あ"},
				TopicPiece: model.TopicPiece{TopicPiece: "人生を変える本"},
				NumOfLikes: model.NumOfLikes{NumOfLikes: 2},
			}},
			want: model.Topic{
				topic: model.Topic{
					StartChar:  model.StartChar{StartChar: "あ"},
					TopicPiece: model.TopicPiece{TopicPiece: "人生を変える本"},
					NumOfLikes: model.NumOfLikes{NumOfLikes: 3},
				}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &likeTopicUsecase{
				topicRepo: tt.fields.topicRepo,
			}
			got, err := r.LikeTopic(tt.args.topic)
			if (err != nil) != tt.wantErr {
				t.Errorf("likeTopicUsecase.LikeTopic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("likeTopicUsecase.LikeTopic() = %v, want %v", got, tt.want)
			}
		})
	}
}
