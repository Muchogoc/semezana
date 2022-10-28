package auth

import (
	"testing"

	"github.com/google/uuid"
)

func TestVerifyToken(t *testing.T) {
	token, err := CreateToken(uuid.NewString())
	if err != nil {
		t.Errorf("failed to create token")
	}

	type args struct {
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
				tokenString: token,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := VerifyToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetUIDFromToken(t *testing.T) {
	uid := uuid.NewString()
	token, err := CreateToken(uid)
	if err != nil {
		t.Errorf("failed to create token")
	}

	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "happy case: get uid",
			args: args{
				tokenString: token,
			},
			want:    uid,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUIDFromToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUIDFromToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUIDFromToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
