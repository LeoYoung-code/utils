package encrypt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateTOTP(t *testing.T) {
	type args struct {
		secret string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Test generate TOTP",
			args: args{
				secret: "JBSWY3DPEHPK3PXP",
			},
			want:    "235678",
			wantErr: assert.NoError,
		},
		{
			name: "Test generate TOTP with invalid secret",
			args: args{
				secret: "invalid",
			},
			want:    "",
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateTOTP(tt.args.secret)
			if !tt.wantErr(t, err, fmt.Sprintf("generateTOTP(%v)", tt.args.secret)) {
				return
			}
			assert.Equalf(t, tt.want, got, "generateTOTP(%v)", tt.args.secret)
		})
	}
}
