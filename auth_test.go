package auth

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
)

func TestCreateToken(t *testing.T) {

	config := Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "test@123",
	}

	auth := AuthSetup(config)

	token, _ := auth.CreateToken()

	if token == "" {

		t.Errorf("Error: Empty token")

	} else {

		fmt.Println(token)
	}

}

func TestLogin(t *testing.T) {

	type TestCase struct {
		Username string
		Password string
	}

	test := []TestCase{
		{
			Username: "Admin",
			Password: "Admin@123",
		},
		{
			Username: "Admin",
			Password: "",
		},
	}

	config := Config{
		ExpiryTime: 2, // It should be in hours not minutes or seconds
		SecretKey:  "test@123",
		DB:         &gorm.DB{},
	}

	auth := AuthSetup(config)

	for _, val := range test {

		t.Run("checklogin", func(t *testing.T) {

			auth.Checklogin(val.Username, val.Password)

		})
	}

}
