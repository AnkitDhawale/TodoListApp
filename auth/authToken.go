package auth

import (
	"errors"
	"github.com/AnkitDhawale/TodoListApp/dto"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

const (
	ACCESS_TOKEN_DURATION = time.Hour
	HMAC_SAMPLE_SECRET    = "hmacSampleSecret"
)

type AccessTokenClaims struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type AuthToken struct {
	token *jwt.Token
}

func NewAccessToken(claims *AccessTokenClaims) (*dto.LoginResponse, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(HMAC_SAMPLE_SECRET))
	if err != nil {
		log.Println("failed to sign token: ", err.Error())

		return nil, errors.New("failed to generate access token")
	}

	return &dto.LoginResponse{AccessToken: signedToken}, nil
}
