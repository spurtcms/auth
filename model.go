package auth

import (
	"time"

	"gorm.io/gorm"
)

type tbluser struct {
	Id               int `gorm:"primaryKey;auto_increment"`
	Uuid             string
	FirstName        string
	LastName         string
	RoleId           int
	Email            string
	Username         string
	Password         string
	MobileNo         string
	IsActive         int
	ProfileImage     string
	ProfileImagePath string
	DataAccess       int
	CreatedOn        time.Time
	CreatedBy        int
	ModifiedOn       time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"DEFAULT:NULL"`
	LastLogin        time.Time `gorm:"DEFAULT:NULL"`
	IsDeleted        int
	DeletedOn        time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy        int       `gorm:"DEFAULT:NULL"`
}

// soft delete check
func IsDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("is_deleted = 0")
}

// check db userlogin
func CheckLogin(username string, Password string, db *gorm.DB) (user tbluser, err error) {

	if err := db.Table("tbl_users").Scopes(IsDeleted).Where("username = ?", username).First(&user).Error; err != nil {

		return tbluser{}, err

	}

	return user, nil
}
