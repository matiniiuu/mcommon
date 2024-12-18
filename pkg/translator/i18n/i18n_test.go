package i18n

import (
	"testing"

	"github.com/matiniiuu/mcommon/pkg/translator"
	"github.com/matiniiuu/mcommon/pkg/translator/messages"
)

func TestMessageBundle_Translate(t *testing.T) {
	type args struct {
		message  string
		language translator.Language
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "translate farsi",
			args: args{message: messages.GeneralError, language: "fa"},
			want: "خطایی وجود دارد",
		},
		{
			name: "translate english",
			args: args{message: messages.UserNotFound, language: "en"},
			want: "user not found",
		},
		{
			name: "message key not found",
			args: args{message: "NoKeyFound", language: "en"},
			want: "NoKeyFound",
		},
	}

	translator, err := New("testdata")
	if err != nil {
		t.Errorf("New() error : %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := translator.Translate(tt.args.message, tt.args.language)
			if got != tt.want {
				t.Errorf("Translate() got = %v, want %v", got, tt.want)
			}

		})
	}
}
