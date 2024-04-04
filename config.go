package auth

import "gorm.io/gorm"

type Config struct {
	UserId     int
	ExpiryTime int //It should be an hour not a mintues
	SecretKey  string
	DB         *gorm.DB
}

type Auth struct {
	UserId     int
	ExpiryTime int
	SecretKey  string
	DBString   *gorm.DB
}
