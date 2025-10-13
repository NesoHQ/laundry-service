package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/enghasib/laundry_service/utils"
)

func (m *Middlewares) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Authentication middlewares call.....")
		// header
		AuthenticationHeader := r.Header.Get("Authorization")
		if AuthenticationHeader == "" {
			http.Error(w, "Unauthorized:", http.StatusUnauthorized)
			return
		}

		//split header and grep the token
		headerArr := strings.Split(AuthenticationHeader, " ")
		if len(headerArr) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		jwt_token := headerArr[1]

		// verify token
		isVerified, err := utils.Verify(jwt_token, m.cnf.JwtSecretKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !isVerified {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		user, err := utils.DecodeToken(jwt_token, m.cnf.JwtSecretKey)
		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, "failed to verify")
		}

		ctx := context.WithValue(r.Context(), m.cnf.Auth_ctx_key, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
