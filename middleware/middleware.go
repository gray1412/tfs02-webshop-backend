package middleware

import (
	"net/http"

	jwt "project-backend/util/jwt"
	response "project-backend/util/response"
)

func TokenAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := jwt.ExtractToken(r)
		claims, err := jwt.VerifyToken(token)
		if err != nil {
			response.RespondWithJSON(w, 401, 0, err.Error(), nil)
			return
		}
		r.Header.Set("email", claims.Email)
		next(w, r)
	}
}
