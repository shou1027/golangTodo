package task

import (
	"unicode/utf8"

	errDomain "github/code-kakitai/code-kakitai/domain/error"

	"github.com/code-kakitai/go-pkg/ulid"
)

type Task struct {
	id     string
	title  string
	detail string
}

// 永続化層から取得したデータをドメインに変換
func Reconstruct(
	id string,
	title string,
	detail string,
) (*Task, error) {
	return newTask(
		id,
		title,
		detail,
	)
}

func NewTask(
	title string,
	detail string,
) (*Task, error) {
	return newTask(
		ulid.NewULID(),
		title,
		detail,
	)
}

func newTask(
	id string,
	title string,
	detail string,
) (*Task, error) {
	//タイトルバリデーション
	if utf8.RuneCountInString(title) > titleLengthMax || utf8.RuneCountInString(title) < titleLengthMin {
		return nil, errDomain.NewError("タイトルの文字数が不正です。")
	}

	//タスク内容バリデーション
	if utf8.RuneCountInString(detail) > detailLengthMax || utf8.RuneCountInString(detail) < detailLengthMin {
		return nil, errDomain.NewError("内容の文字数が不正です。")
	}

	return &Task{
		id:     id,
		title:  title,
		detail: detail,
	}, nil
}

func (t *Task) ID() string {
	return t.id
}

func (t *Task) Title() string {
	return t.title
}

func (t *Task) Detail() string {
	return t.detail
}

const (
	titleLengthMax  = 20
	titleLengthMin  = 1
	detailLengthMax = 180
	detailLengthMin = 1
)
