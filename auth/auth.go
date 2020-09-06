package auth

import (
	"net/http"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

// GetTokenHandler get token
var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)

	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = "54546557354"
	claims["name"] = "taro"
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// 電子署名
	tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
	tokenString,_ = token.SignedString([byte("u8PzVT8fiumhXc9Ng7iXdh6Qdl4ow1c2QmEeYfDn6IXVsKY4U51N_DEPCyEiFrlpLS1t4pM9E-W1cYxRp7Zq4-3trBu9a9HczPSnJ9h69JiIedg76WYWbVzZS_Pe4zsak1UN-jYe-yQ581FLU1wkhUwV-0K--IsH00bieBI3as8")])

	// JWTを返却
	w.Write([]byte(tokenString))
})

// JwtMiddleware check token
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte("u8PzVT8fiumhXc9Ng7iXdh6Qdl4ow1c2QmEeYfDn6IXVsKY4U51N_DEPCyEiFrlpLS1t4pM9E-W1cYxRp7Zq4-3trBu9a9HczPSnJ9h69JiIedg76WYWbVzZS_Pe4zsak1UN-jYe-yQ581FLU1wkhUwV-0K--IsH00bieBI3as8"), nil
		//return []byte(os.Getenv("SIGNINGKEY")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
