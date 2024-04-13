package auth

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// AuthSetup used initialize auth configruation
func AuthSetup(conf Config) *Auth {

	return &Auth{
		UserId:     conf.UserId,
		ExpiryTime: conf.ExpiryTime,
		SecretKey:  conf.SecretKey,
		DBString:   conf.DB,
	}

}

// Check UserName Password
func (auth *Auth) Checklogin(Username string, Password string) (string, int, error) {

	username := Username

	password := Password

	user, err := CheckLogin(username, password, auth.DBString)

	if err != nil {

		log.Println(err)

	}

	passerr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if passerr != nil || passerr == bcrypt.ErrMismatchedHashAndPassword {

		return "", 0, ErrorPassword

	}

	token, err := auth.CreateToken()

	if err != nil {

		return "", 0, err
	}

	auth.UserId = user.Id

	return token, user.Id, nil
}

// CreateToken creates a token
func (auth *Auth) CreateToken() (string, error) {

	atClaims := jwt.MapClaims{}

	atClaims["user_id"] = auth.UserId

	atClaims["expiry_time"] = time.Now().Add(time.Duration(auth.ExpiryTime) * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	return token.SignedString([]byte(auth.SecretKey))
}

// verify token
func (auth *Auth) VerifyToken(token string, secret string, currentTime int64) (userid int, err error) {

	Claims := jwt.MapClaims{}

	tkn, err := jwt.ParseWithClaims(token, Claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println(err)
			return 0, ErrorToken
		}

		return 0, ErrorToken
	}

	if !tkn.Valid {
		fmt.Println(tkn)
		return 0, ErrorToken
	}

	expiryTime := Claims["expiry_time"]

	if currentTime > int64(expiryTime.(float64)) {

		return 0, ErrorTokenExpiry
	}

	usrid := Claims["user_id"]

	auth.AuthFlg = true

	return int(usrid.(float64)), nil
}
