package infra

import (
	"math/rand"
	"time"

	"github.com/guregu/dynamo"
	"github.com/ogady/find_the_right_answer/api/domain/model"
	"github.com/ogady/find_the_right_answer/api/domain/repository"
	awsInfra "github.com/ogady/find_the_right_answer/api/infra/aws"
)

type TopicPieceRepoImpl struct {
	dynamoDB *dynamo.DB
	table    *dynamo.Table
}

func NewTopicPieceRepoImpl() repository.TopicPieceRepository {

	db := awsInfra.NewDynamoDBConn()

	table := db.Table("topic_piece")
	topicPieceRepoImpl := &TopicPieceRepoImpl{
		dynamoDB: db,
		table:    &table,
	}

	return topicPieceRepoImpl
}

func (r *TopicPieceRepoImpl) Save(topicPiece *model.TopicPiece) error {
	var err error
	tp := model.TopicPiece{TopicPiece: topicPiece.TopicPiece}
	err = r.table.Put(tp).Run()
	return err
}

func (r *TopicPieceRepoImpl) FindAll() ([]model.TopicPiece, error) {
	var topicPieces []model.TopicPiece
	var err error

	err = r.table.Scan().All(&topicPieces)
	if err != nil {
		return nil, err
	}
	return topicPieces, err
}

func (r *TopicPieceRepoImpl) Find(topicID string) (model.TopicPiece, error) {
	var topicPiece model.TopicPiece
	var err error

	return topicPiece, err
}

func (r *TopicPieceRepoImpl) FindRandom() (model.TopicPiece, error) {
	var topicPieces []model.TopicPiece
	var topicPiece model.TopicPiece
	var err error
	err = r.table.Scan().All(&topicPieces)
	if err != nil {
		return topicPiece, err
	}

	cnt := len(topicPieces)
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(cnt)

	return topicPieces[i], err
}
