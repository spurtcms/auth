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

	Migration(conf.DB)

	return &Auth{
		UserId:     conf.UserId,
		ExpiryTime: conf.ExpiryTime,
		SecretKey:  conf.SecretKey,
		DB:         conf.DB,
		ExpiryFlg:  conf.ExpiryFlg,
	}

}

// Check UserName Password
func (auth *Auth) Checklogin(Username string, Password string) (string, int, error) {

	username := Username

	password := Password

	user, err := CheckLogin(username, password, auth.DB)

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

	atClaims["expiry_time"] = time.Now().UTC().Add(time.Duration(auth.ExpiryTime) * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	return token.SignedString([]byte(auth.SecretKey))
}

// verify token
func (auth *Auth) VerifyToken(token string, secret string) (userid int, err error) {

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

	if auth.AuthFlg {

		expiryTime := Claims["expiry_time"]

		t, ok := expiryTime.(time.Time)

		if !ok {

			fmt.Println("Could not convert interface to time.Time")

			return 0, ErrorConvertTime
		}

		if t.After(time.Now().UTC()) {

			return 0, ErrorTokenExpiry

		}
	}

	usrid := Claims["user_id"]

	auth.AuthFlg = true

	return int(usrid.(float64)), nil
}

// Check User Permission
func (permission *Auth) IsGranted(modulename string, permisison Action) (bool, error) {

	if permission.RoleId != 1 || permission.RoleName != "Super Admin" { //if not an admin user

		var modid int

		var module TblModule

		var modpermissions TblModulePermission

		if err := permission.DB.Model(TblModule{}).Where("module_name=? and parent_id !=0", modulename).Find(&module).Error; err != nil {

			return false, err
		}

		if err1 := permission.DB.Model(TblModulePermission{}).Where("display_name=?", modulename).Find(&modpermissions).Error; err1 != nil {

			return false, err1
		}

		if module.Id != 0 {

			modid = module.Id

		} else {

			modid = modpermissions.Id
		}

		var modulepermission []TblModulePermission

		if permisison == "CRUD" {

			if err := permission.DB.Model(TblModulePermission{}).Where("id=? and (full_access_permission=1 or display_name='View' or display_name='Update' or  display_name='Create' or display_name='Delete')", modid).Find(&modulepermission).Error; err != nil {

				return false, err
			}

		} else {

			if err := permission.DB.Model(TblModulePermission{}).Where("module_id=? and display_name=?", modid, permisison).Find(&modulepermission).Error; err != nil {

				return false, err
			}

		}

		for _, val := range modulepermission {

			var rolecheck TblRolePermission

			if err := permission.DB.Model(TblRolePermission{}).Where("permission_id=? and role_id=?", val.Id, permission.RoleId).First(&rolecheck).Error; err != nil {

				return false, err
			}

		}

		permission.PermissionFlg = true

	}

	permission.PermissionFlg = true

	return true, nil

}
