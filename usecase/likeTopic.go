package usecase

type likeTopicUsecase struct {
}

// NewLikeTopicUsecase -
func NewLikeTopicUsecase() createTopicUsecase {

	return createTopicUsecase{}
}

func (r *likeTopicUsecase) LikeTopic(id string) error {
	var err error
	return err
}
