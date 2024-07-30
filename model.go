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

type MemberLoginCheck struct {
	Username             string
	Email                string
	Password             string
	OTP                  int
	UsernameWithOTP      bool
	EmailWithOTP         bool
	UsernameWithPassword bool
	EmailwithPassword    bool
}

type TblMember struct {
	Id               int
	Uuid             string
	FirstName        string
	LastName         string
	Email            string
	MobileNo         string
	IsActive         int
	ProfileImage     string
	ProfileImagePath string
	LastLogin        int
	MemberGroupId    int
	Password         string
	Username         string
	Otp              int
	OtpExpiry        time.Time
	LoginTime        time.Time
	IsDeleted        int
	DeletedOn        time.Time
	DeletedBy        int
	CreatedOn        time.Time
	CreatedBy        int
	ModifiedOn       time.Time
	ModifiedBy       int
	TenantId         int
}

type TblModule struct {
	Id               int
	ModuleName       string
	IsActive         int
	DefaultModule    int
	ParentId         int
	IconPath         string
	AssignPermission int
	Description      string
	OrderIndex       int
	CreatedBy        int
	CreatedOn        time.Time
	MenuType         string
}

type TblModulePermission struct {
	Id                   int
	RouteName            string
	DisplayName          string
	SlugName             string
	Description          string
	ModuleId             int
	FullAccessPermission int
	ParentId             int
	AssignPermission     int
	BreadcrumbName       string
	OrderIndex           int
	CreatedBy            int
	CreatedOn            time.Time
	ModifiedBy           int
	ModifiedOn           time.Time
}

type TblRolePermission struct {
	Id           int
	RoleId       int
	PermissionId int
	CreatedBy    int
	CreatedOn    time.Time
}

type OTP struct {
	Length   int
	Duration time.Duration //minutes only
	MemberId int
}

type authmodel struct{}

var Authmodel authmodel

// soft delete check
func IsDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("is_deleted = 0")
}

// check db userlogin
func (auth authmodel) CheckLogin(username string, Password string, db *gorm.DB, tenantid int) (user tbluser, err error) {

	if err := db.Table("tbl_users").Where("username = ? and tenant_id=? and is_deleted = 0", username, tenantid).First(&user).Error; err != nil {

		return tbluser{}, err

	}

	return user, nil
}

// check email with password
func (auth authmodel) CheckMemberLoginWithEmail(email string, username string, DB *gorm.DB, tenantid int) (member TblMember, err error) {

	if email != "" {

		if err := DB.Model(TblMember{}).Where("is_deleted =0 and email=? and tenant_id=?", email, tenantid).First(&member).Error; err != nil {

			return TblMember{}, err
		}

	} else if username != "" {

		if err := DB.Model(TblMember{}).Where("is_deleted =0 and username=? and tenant_id=? ", username, tenantid).First(&member).Error; err != nil {

			return TblMember{}, err
		}

	}

	return member, nil
}

func (auth authmodel) CheckEmailWithOtp(email string, DB *gorm.DB, tenantid int) (member TblMember, err error) {

	if err := DB.Model(TblMember{}).Where("is_deleted = 0 and email = ? and tenant_id=?", email, tenantid).First(&member).Error; err != nil {

		return TblMember{}, err
	}

	return member, nil
}

func (auth authmodel) CheckUsernameWithOtp(username string, DB *gorm.DB, tenantid int) (member TblMember, err error) {

	if err := DB.Model(TblMember{}).Where("is_deleted = 0 and username = ? and tenant_id=?", username, tenantid).First(&member).Error; err != nil {

		return TblMember{}, err
	}

	return member, nil
}

func (auth authmodel) UpdateMemberOtp(id int, otp int, otpExpiry string, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_members").Where("id=? and tenant_id=?", id, tenantid).Updates(map[string]interface{}{
		"otp": otp, "otp_expiry": otpExpiry,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (auth authmodel) GetMemberDetailsByMemberId(MemberDetails *TblMember, memberId int, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_members").Where("is_deleted=0 and id = ? and tenant_id=?", memberId, tenantid).First(&MemberDetails).Error; err != nil {
		return err
	}

	return nil
}
