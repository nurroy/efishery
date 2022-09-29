package auth

import (
	"belajar/efishery/configs"
	"belajar/efishery/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"

	"net/http"
	"strings"
	"time"
)

func CreateToken(authD models.Auth,config *configs.Config) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userid"] = authD.UserID
	claims["username"] = authD.Username
	claims["role"] = authD.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Secret.Key))
}

// Token Validation
func TokenValid(r *http.Request, auth string,config *configs.Config) (jwt.MapClaims,error) {
	token, err := VerifyToken(r,auth,config)
	if err != nil {
		return nil,err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims,nil
}

// Verify Token
func VerifyToken(r *http.Request,auth string,config *configs.Config) (*jwt.Token, error) {
	var tokenString string
	if r != nil{
		tokenString = ExtractToken(r)
	}else{
		tokenString = auth
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//does this token conform to "SigningMethodHMAC" ?
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Secret.Key), nil
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

