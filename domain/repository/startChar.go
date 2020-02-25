package repository

import (
	"github.com/ogady/find_the_right_answer/domain/model"
)

type StartCharRepository interface {
	FindRandom() model.StartChar
}
