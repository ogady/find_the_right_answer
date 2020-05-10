package usecase

import (
	"testing"

	"github.com/ogady/find_the_right_answer/api/domain/model"
)

func Test_likeTopicUsecase_LikeTopic(t *testing.T) {
	type args struct {
		topic model.Topic
	}

	tests := []struct {
		name    string
		r       *likeTopicUsecase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "normal_test",
			r:    new(likeTopicUsecase),
			args: args{topic: model.Topic{
				StartChar:  model.StartChar{StartChar: "あ"},
				TopicPiece: model.TopicPiece{TopicPiece: "人生を変える本"},
				NumOfLikes: model.NumOfLikes{NumOfLikes: 2},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &likeTopicUsecase{}
			if _, err := r.LikeTopic(tt.args.topic); (err != nil) != tt.wantErr {
				t.Errorf("likeTopicUsecase.LikeTopic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
