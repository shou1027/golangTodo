package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewUser(t *testing.T) {
	type args struct {
		name  string
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				name:  "テスト　太郎",
				email: "test@example.com",
			},
			want: &User{
				name:  "テスト　太郎",
				email: "test@example.com",
			},
			wantErr: false,
		},
		{
			name: "異常系: 名前が空文字",
			args: args{
				name:  "",
				email: "test@example.com",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: メールアドレスが不正",
			args: args{
				name:  "",
				email: "testcom",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.name, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(
				got, tt.want,
				cmp.AllowUnexported(User{}),
				cmpopts.IgnoreFields(User{}, "id"),
			)
			if diff != "" {
				t.Errorf("NewUser() = %v, want %v. error is %s", got, tt.want, err)
			}
		})
	}
}
