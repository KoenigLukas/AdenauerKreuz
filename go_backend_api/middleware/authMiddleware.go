package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/koeniglukas/config"
	"net/http"
	"strconv"
	"strings"
)


type Claims struct {
	UserID int `json:"userID"`
	jwt.StandardClaims
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization")," ")
		if len(authHeader)!=2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		} else {
		jwtToken := authHeader[1]

		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Get("JWT_SECRET")), nil
		})

			s := strconv.FormatInt(int64(claims.UserID), 10)
		if claims.UserID != 0 && tkn.Valid{
			r.Header.Set("userID",s)
			next.ServeHTTP(w, r)
		} else {
			fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
	}
	})
}

