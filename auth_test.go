package auth

import (
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBSetup() (*gorm.DB, error) {

	dbConfig := map[string]string{
		"username": "postgres",
		"password": "123",
		"host":     "localhost",
		"port":     "5432",
		"dbname":   "spurt-cms",
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "user=" + dbConfig["username"] + " password=" + dbConfig["password"] +
			" dbname=" + dbConfig["dbname"] + " host=" + dbConfig["host"] +
			" port=" + dbConfig["port"] + " sslmode=disable TimeZone=Asia/Kolkata",
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestCreateToken(t *testing.T) {

	db, _ := DBSetup()

	config := Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "test@123",
		DB:         db,
	}

	log.Println("config", config)

	auth := AuthSetup(config)

	token, _ := auth.CreateToken()

	if token == "" {

		t.Errorf("Error: Empty token")

	} else {

		fmt.Println("token", token)
	}

}

func TestLogin(t *testing.T) {

	db, _ := DBSetup()

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
		{
			Username: "",
			Password: "Admin@123",
		},
	}

	config := Config{
		ExpiryTime: 2, // It should be in hours not minutes or seconds
		SecretKey:  "test@123",
		DB:         db,
	}

	auth := AuthSetup(config)

	for _, val := range test {

		t.Run("checklogin", func(t *testing.T) {

			_, _, err := auth.Checklogin(val.Username, val.Password, 1)

			log.Println(err)

		})
	}

}

func TestVerifyToken(t *testing.T) {

	db, _ := DBSetup()

	config := Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "test@123",
		DB:         db,
	}

	auth := AuthSetup(config)

	token, _ := auth.CreateToken()

	_, _, err1 := auth.VerifyToken(token, "test@123")

	log.Println(err1)
}