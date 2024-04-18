package user

import (
	"net/mail"
	"unicode/utf8"

	"github.com/code-kakitai/go-pkg/ulid"

	errDomain "github/code-kakitai/code-kakitai/domain/error"
)

type User struct {
	id    string
	name  string
	email string
}

// 永続化層から取得したデータをドメインに変換
func Reconstruct(
	id string,
	name string,
	email string,
) (*User, error) {
	return newUser(
		id,
		name,
		email,
	)
}

func NewUser(
	name string,
	email string,
) (*User, error) {
	return newUser(
		ulid.NewULID(),
		name,
		email,
	)
}

func newUser(
	id string,
	name string,
	email string,
) (*User, error) {
	// 名前バリデーション
	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, errDomain.NewError("名前の文字数が不正です。")
	}

	// メールアドレスバリデーション
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, errDomain.NewError("メールアドレスが不正です。")
	}

	return &User{
		id:    id,
		name:  name,
		email: email,
	}, nil
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

const (
	nameLengthMax = 255
	nameLengthMin = 1
)
