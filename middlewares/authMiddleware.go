package middlewares

import (
	"context"
	"errors"
	"fmt"
	"github.com/AnkitDhawale/TodoListApp/auth"
	"github.com/AnkitDhawale/TodoListApp/helpers"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
)

type contextKey string

const (
	bearerPrefix string = "Bearer "

	UserIDKey    contextKey = "userId"
	UserEmailKey contextKey = "userEmail"
)

func TokenResolver(userResolver *helpers.UserResolver) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log.Println("in TokenResolver")
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			decodeToken := getBearerTokenFromRequest(r)
			if decodeToken == "" {
				helpers.WriteResponse(w, http.StatusUnauthorized, nil, errors.New("missing bearer token value"))
				return
			}

			token, err := jwt.ParseWithClaims(decodeToken, &auth.AccessTokenClaims{}, func(token *jwt.Token) (any, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
				return []byte(auth.HMAC_SAMPLE_SECRET), nil
			})
			if err != nil {
				log.Println("failed to parse token, err: ", err)
				helpers.WriteResponse(w, http.StatusUnauthorized, nil, err)
				return
			}
			tokenClaims := token.Claims.(*auth.AccessTokenClaims)

			_, err = token.Claims.GetSubject()
			if err != nil {
				helpers.WriteResponse(w, http.StatusUnauthorized, nil, errors.New("missing subject from token"))
				return
			} else {
				// validate user with DB
				_, err := userResolver.Resolve(tokenClaims.UserId, tokenClaims.Email)
				if err != nil {
					helpers.WriteResponse(w, http.StatusUnauthorized, nil, err)
					return
				}

				// store userId & email in the context
				ctx := context.WithValue(r.Context(), UserIDKey, tokenClaims.UserId)
				ctx = context.WithValue(ctx, UserEmailKey, tokenClaims.Email)

				next.ServeHTTP(w, r.WithContext(ctx))
			}
		})
	}
}

// getBearerTokenFromRequest returns the bearer token from an HTTP request.
func getBearerTokenFromRequest(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")

	if strings.HasPrefix(bearerToken, bearerPrefix) {
		return strings.TrimPrefix(bearerToken, bearerPrefix)
	}
	return bearerToken
}
