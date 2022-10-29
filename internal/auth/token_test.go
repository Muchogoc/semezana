package auth

import (
	"context"
	"testing"

	"github.com/google/uuid"
)

func TestVerifyToken(t *testing.T) {
	token, err := CreateToken(uuid.NewString())
	if err != nil {
		t.Errorf("failed to create token")
	}

	type args struct {
		ctx         context.Context
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "happy case: valid token",
			args: args{
				ctx:         context.Background(),
				tokenString: token,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseToken(tt.args.ctx, tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
