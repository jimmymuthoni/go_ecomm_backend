package api

import (
	"errors"
	"fmt"
	"os"

	"github.com/anthdm/weavebox"
	"github.com/golang-jwt/jwt/v4"
)

var ErrUnAuthenticated = errors.New("unAuthenticated")

type AdminAuthMiddleware struct{}

func (mw *AdminAuthMiddleware) Authenticate(ctx *weavebox.Context) error {
	tokenString := ctx.Header("x-api-token")
	if len(tokenString) == 0{
		return ErrUnAuthenticated
	}
	token,err := parseJWT(tokenString)
	if err != nil {
		return ErrUnAuthenticated
	}
	if !token.Valid{
		return ErrUnAuthenticated
	}
	claims,ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ErrUnAuthenticated
	}
	fmt.Println(claims)
	fmt.Println("Gurding the admin routes")
	return nil
}

func parseJWT(tokenString string) (*jwt.Token, error){
	secret := os.Getenv("JWT_SECRET")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unxtepected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
}