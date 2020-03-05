package infra

import (
	"github.com/ogady/find_the_right_answer/domain/model"
	"github.com/ogady/find_the_right_answer/domain/repository"
)

type StartCharRepoImplMock struct {
	charList []string
}

func NewStartCharRepoImplMock() repository.StartCharRepository {

	startCharRepoImpl := StartCharRepoImplMock{charList: create()}

	return &startCharRepoImpl
}
func (r *StartCharRepoImplMock) FindRandom() model.StartChar {

	var sc model.StartChar

	sc = model.StartChar{StartChar: "あ"}

	return sc
}
