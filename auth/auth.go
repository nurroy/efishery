package auth

import (
	"belajar/efishery/models"
	utils "belajar/efishery/utils/response"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"

	"net/http"
	"strings"
	"time"
)

func CreateToken(authD models.Auth) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userid"] = authD.UserID
	claims["username"] = authD.Username
	claims["role"] = authD.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET")))
}

// Token Validation
func TokenValid(r *http.Request, auth string) (jwt.MapClaims,error) {
	token, err := VerifyToken(r,auth)
	if err != nil {
		return nil,err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	if r != nil {
		if strings.Contains(r.URL.Path, "aggregate") {
			if claims["role"] != "admin" {
				return nil, utils.NewErr("You need to be an admin to perform this api call")
			}
		}
	}

	return claims,nil
}

// Verify Token
func VerifyToken(r *http.Request,auth string) (*jwt.Token, error) {
	var tokenString string
	if r != nil{
		tokenString = ExtractToken(r)
	}else{
		tokenString = auth
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,errors.New("Signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil,errors.New("Signing method invalid")
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

//get the token from the request body
func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

