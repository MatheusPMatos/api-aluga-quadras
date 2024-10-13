package utils

import (
	"strconv"
	"testing"

	"github.com/MatheusPMatos/api-aluga-quadras/config"
	"github.com/stretchr/testify/assert"
)

func TestJwtDecodeToken(t *testing.T) {

	tests := []struct {
		name        string
		username    string
		userID      uint
		expectedErr error
	}{
		{
			name:   "Success",
			userID: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serv := NewJwt(config.Environments{TokenSecret: "123123123"})

			tk, err := serv.CreateAccesstoken(tt.userID)
			if err != nil {

				t.Fatalf("error: %s", err.Error())
			}
			claim, err := serv.DecodeAccessToken(tk)
			if err != nil {

				t.Fatalf("error: %s", err.Error())
			}
			sub, err := claim.GetSubject()
			if err != nil {
				t.Fatalf("error: %s", err.Error())
			}
			id, err := strconv.Atoi(sub)
			if err != nil {

				t.Fatalf("error: %s", err.Error())
			}
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.userID, uint(id))

		})
	}
}
