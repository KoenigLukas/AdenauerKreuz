package middleware

import (
	"../config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)


//type Claims struct {
//	UserID int `json:"userID"`
//}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization")," ")
		if len(authHeader)!=2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		} else {
		jwtToken := authHeader[1]
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.Get("JWT_SECRET")), nil
		})


		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			r.Header.Set("userID",(claims)["userID"].(string))
			next.ServeHTTP(w, r)
		} else {
			fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
	}

		secret := config.Get("JWT_SECRET")
		fmt.Println(secret)
		next.ServeHTTP(w, r)
	})
}

