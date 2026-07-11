package middlewares

import (
	"AssetTrack/utils"
	"context"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

type ContextKeys string

const (
	userContext ContextKeys = "userContext"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("x-api-key")
		if apiKey == "" {
			utils.RespondError(w, http.StatusUnauthorized, nil, "token header missing")
			return
		}

		token, err := jwt.Parse(apiKey, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			utils.RespondError(w, http.StatusUnauthorized, err, "invalid token")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.RespondError(w, http.StatusUnauthorized, nil, "invalid token claims")
			return
		}

		userID, ok := claims["userId"].(string)
		if !ok {
			utils.RespondError(w, http.StatusUnauthorized, nil, "invalid token claims")
			return
		}

		ctx := context.WithValue(r.Context(), userContext, userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
func UserContext(r *http.Request) string {
	userID := r.Context().Value(userContext).(string)
	return userID
}
