package middlewares

import (
	"AssetTrack/utils"
	"context"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

type ContextKey string

const UserContextKey ContextKey = "user"

func Authenticate(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("x-api-key")

		if tokenString == "" {
			utils.RespondError(w, http.StatusUnauthorized, nil, "token missing")
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}

			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			utils.RespondError(w, http.StatusUnauthorized, err, "invalid token")
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		user := models.UserContext{
			UserID: claims["userId"].(string),
			Role:   claims["role"].(string),
		}

		ctx := context.WithValue(r.Context(), UserContextKey, user)

<<<<<<< Updated upstream
		ctx := context.WithValue(r.Context(), userContext, userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
func UserContext(r *http.Request) (string, bool) {
	userID, ok := r.Context().Value(userContext).(string)
	return userID, ok
=======
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserContext(r *http.Request) models.UserContext {

	return r.Context().Value(UserContextKey).(models.UserContext)
}

func RequireRoles(next http.Handler, roles ...string) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user := GetUserContext(r)

		for _, role := range roles {

			if user.Role == role {
				next.ServeHTTP(w, r)
				return
			}
		}

		utils.RespondError(w, http.StatusForbidden, nil, "access denied")
	})
>>>>>>> Stashed changes
}
