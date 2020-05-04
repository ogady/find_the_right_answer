package infra

import (
	"math/rand"
	"time"

	"github.com/ogady/find_the_right_answer/api/domain/model"
	"github.com/ogady/find_the_right_answer/api/domain/repository"
)

type StartCharRepoImpl struct {
	charList []string
}

func NewStartCharRepoImpl() repository.StartCharRepository {

	startCharRepoImpl := StartCharRepoImpl{charList: create()}

	return &startCharRepoImpl
}
func (r *StartCharRepoImpl) FindRandom() model.StartChar {

	var sc model.StartChar

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(44)

	sc = model.StartChar{StartChar: (r.charList[i])}

	return sc
}

func create() []string {
	chars := []string{
		"あ",
		"い",
		"う",
		"え",
		"お",
		"か",
		"き",
		"く",
		"け",
		"こ",
		"さ",
		"し",
		"す",
		"せ",
		"そ",
		"た",
		"ち",
		"つ",
		"て",
		"と",
		"な",
		"に",
		"ぬ",
		"ね",
		"の",
		"は",
		"ひ",
		"ふ",
		"へ",
		"ほ",
		"ま",
		"み",
		"む",
		"め",
		"も",
		"や",
		"ゆ",
		"よ",
		"ら",
		"り",
		"る",
		"れ",
		"ろ",
		"わ",
	}
	return chars
}
