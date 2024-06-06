package auth

import "gorm.io/gorm"

type Type string

const ( //for permission check
	Postgres Type = "postgres"
	Mysql    Type = "mysql"
)

type Config struct {
	UserId       int      //(optional) if you use login function this userid no need
	ExpiryTime   int      //It should be an hour not a mintues, UTC time only
	ExpiryFlg    bool     //if you want to check token expiry time enable expiryflg true otherwise expirytime not check
	SecretKey    string   //jwt secretkey
	DB           *gorm.DB //database connection
	DataBaseType Type
	RoleId       int
	RoleName     string
	OTPAttempt   int //how many times otp attempt
}

type Authentication struct {
	Token     string
	SecretKey string
}

type Auth struct {
	UserId        int
	ExpiryTime    int
	ExpiryFlg     bool
	SecretKey     string
	DB            *gorm.DB
	AuthFlg       bool
	PermissionFlg bool
	RoleId        int
	RoleName      string
	DataAccess    int
}
