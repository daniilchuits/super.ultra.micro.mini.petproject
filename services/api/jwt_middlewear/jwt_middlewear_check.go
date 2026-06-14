package jwtmiddlewear

import (
	"api/domain"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type secretKey struct {
	secret []byte
}

func NewSecretKey(secret []byte) secretKey {
	return secretKey{secret: secret}
}

func (scrt secretKey) JWTCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")

		if auth == "" {
			http.Error(w, domain.ErrEmptyAuthorization.Error(), 400)
			return
		}

		if ok := strings.HasPrefix(auth, "Bearer "); !ok {
			http.Error(w, domain.ErrNoBearerInAuthorization.Error(), 400)
			return
		}

		tokenString := strings.TrimPrefix(auth, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected singing method: %v\n", t.Header["alg"])
			}

			return scrt.secret, nil
		})

		if err != nil {
			log.Println(err)
			http.Error(w, domain.ErrParsingToken.Error(), 400)
			return
		}

		if !token.Valid {
			http.Error(w, domain.ErrTokenIsNotValid.Error(), 400)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Claims type assertion error", 400)
			return
		}
		userIdAny := claims["user_id"]
		userIdFloat, ok := userIdAny.(float64)
		if !ok {
			http.Error(w, "'user_id' in claims is not float", 400)
			return
		}

		userId := int(userIdFloat)

		r.Header.Del("user_id")
		r.Header.Set("user_id", strconv.Itoa(userId))
		next.ServeHTTP(w, r)
	})
}
