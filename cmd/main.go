package main

import (
	"example/sample/global"
	"example/sample/initialize"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func main() {
	initialize.InitComponents()
	engine := initialize.AdminRouters()
	global.Logger.Info().Msgf("hello my logger")
	//validToken, err := GenerateJWT()
	//if err != nil {
	//	fmt.Println("Failed to generate token")
	//}
	//println(validToken)
	if err := engine.Run(":" + global.Config.System.Addr); err != nil {
		println("Error starting server...")
	}
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
