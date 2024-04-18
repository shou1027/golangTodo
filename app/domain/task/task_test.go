package task

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewTask(t *testing.T) {
	type args struct {
		title  string
		detail string
	}
	tests := []struct {
		name    string
		args    args
		want    *Task
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				title:  "テスト勉強",
				detail: "寝る一時間前まで勉強する",
			},
			want: &Task{
				title:  "テスト勉強",
				detail: "寝る一時間前まで勉強する",
			},
			wantErr: false,
		},
		{
			name: "異常系: タイトルが空文字",
			args: args{
				title:  "",
				detail: "寝る一時間前まで勉強する",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: タスク詳細が空文字",
			args: args{
				title:  "テスト勉強",
				detail: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTask(tt.args.title, tt.args.detail)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(
				got, tt.want,
				cmp.AllowUnexported(Task{}),
				cmpopts.IgnoreFields(Task{}, "id"),
			)
			if diff != "" {
				t.Errorf("NewTask() = %v, want %v. error is %s", got, tt.want, err)
			}
		})
	}
}
