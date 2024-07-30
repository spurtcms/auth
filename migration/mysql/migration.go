package mysql

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type TblRole struct {
	Id          int       `gorm:"primaryKey;auto_increment"`
	Name        string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:varchar(255)"`
	Slug        string    `gorm:"type:varchar(255)"`
	IsActive    int       `gorm:"type:int"`
	IsDeleted   int       `gorm:"type:int"`
	CreatedOn   time.Time `gorm:"type:datetime"`
	CreatedBy   int       `gorm:"type:int"`
	ModifiedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL;type:int"`
	TenantId    int       `gorm:"type:int"`
}

type TblUser struct {
	Id                int       `gorm:"primaryKey"`
	Uuid              string    `gorm:"type:varchar(255)"`
	FirstName         string    `gorm:"type:varchar(255)"`
	LastName          string    `gorm:"type:varchar(255)"`
	RoleId            TblRole   `gorm:"type:int;foreignkey:Id"`
	Email             string    `gorm:"type:varchar(255)"`
	Username          string    `gorm:"type:varchar(255)"`
	Password          string    `gorm:"type:varchar(255)"`
	MobileNo          string    `gorm:"type:varchar(255)"`
	IsActive          int       `gorm:"type:int"`
	ProfileImage      string    `gorm:"type:varchar(255)"`
	ProfileImagePath  string    `gorm:"type:varchar(255)"`
	DataAccess        int       `gorm:"type:int"`
	DefaultLanguageId int       `gorm:"type:int"`
	CreatedOn         time.Time `gorm:"type:datetime"`
	CreatedBy         int       `gorm:"type:int"`
	ModifiedOn        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy        int       `gorm:"DEFAULT:NULL;type:int"`
	LastLogin         time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted         int       `gorm:"type:int"`
	DeletedOn         time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy         int       `gorm:"DEFAULT:NULL;type:int"`
	TenantId          int       `gorm:"type:int"`
}

type TblModule struct {
	Id                   int       `gorm:"primaryKey"`
	ModuleName           string    `gorm:"type:varchar(255)"`
	IsActive             int       `gorm:"type:int"`
	DefaultModule        int       `gorm:"type:int"`
	ParentId             int       `gorm:"type:int"`
	IconPath             string    `gorm:"type:varchar(255)"`
	AssignPermission     int       `gorm:"type:int"`
	Description          string    `gorm:"type:varchar(255)"`
	OrderIndex           int       `gorm:"type:int"`
	CreatedBy            int       `gorm:"type:int"`
	CreatedOn            time.Time `gorm:"type:datetime"`
	MenuType             string    `gorm:"type:varchar(255)"`
	FullAccessPermission int       `gorm:"type:int"`
	GroupFlg             int       `gorm:"type:int"`
	TenantId             int       `gorm:"type:int"`
}

type TblModulePermission struct {
	Id                   int       `gorm:"primaryKey"`
	RouteName            string    `gorm:"type:varchar(255)"`
	DisplayName          string    `gorm:"type:varchar(255)"`
	SlugName             string    `gorm:"type:varchar(255)"`
	Description          string    `gorm:"type:varchar(255)"`
	ModuleId             int       `gorm:"type:int"`
	FullAccessPermission int       `gorm:"type:int"`
	ParentId             int       `gorm:"type:int"`
	AssignPermission     int       `gorm:"type:int"`
	BreadcrumbName       string    `gorm:"type:varchar(255)"`
	OrderIndex           int       `gorm:"type:int"`
	CreatedBy            int       `gorm:"type:int"`
	CreatedOn            time.Time `gorm:"type:datetime"`
	ModifiedBy           int       `gorm:"DEFAULT:NULL;type:int"`
	ModifiedOn           time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	TenantId             int       `gorm:"type:int"`
}

type TblRolePermission struct {
	Id           int       `gorm:"primaryKey;auto_increment"`
	RoleId       int       `gorm:"type:int"`
	PermissionId int       `gorm:"type:int"`
	CreatedBy    int       `gorm:"type:int"`
	CreatedOn    time.Time `gorm:"type:datetime"`
	TenantId     int       `gorm:"type:int"`
}

type TblMemberGroup struct {
	Id          int       `gorm:"primaryKey;auto_increment"`
	Name        string    `gorm:"type:varchar(255)"`
	Slug        string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:varchar(255)"`
	IsActive    int       `gorm:"type:int"`
	IsDeleted   int       `gorm:"type:int"`
	CreatedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	CreatedBy   int       `gorm:"type:int"`
	ModifiedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL;type:int"`
	DeletedBy   int       `gorm:"type:int"`
	DeletedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	TenantId    int       `gorm:"type:int"`
}

type TblMember struct {
	Id               int       `gorm:"primaryKey;auto_increment"`
	Uuid             string    `gorm:"type:varchar(255)"`
	FirstName        string    `gorm:"type:varchar(255)"`
	LastName         string    `gorm:"type:varchar(255)"`
	Email            string    `gorm:"type:varchar(255)"`
	MobileNo         string    `gorm:"type:varchar(255)"`
	IsActive         int       `gorm:"type:int"`
	ProfileImage     string    `gorm:"type:varchar(255)"`
	ProfileImagePath string    `gorm:"type:varchar(255)"`
	LastLogin        int       `gorm:"type:int"`
	MemberGroupId    int       `gorm:"type:int"`
	Password         string    `gorm:"type:varchar(255)"`
	Username         string    `gorm:"type:varchar(255)"`
	Otp              int       `gorm:"DEFAULT:NULL;type:int"`
	OtpExpiry        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	LoginTime        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted        int       `gorm:"type:int"`
	DeletedOn        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"DEFAULT:NULL;type:int"`
	CreatedOn        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	CreatedBy        int       `gorm:"type:int"`
	ModifiedOn       time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"DEFAULT:NULL;type:int"`
	TenantId         int       `gorm:"type:int"`
}

type TblMemberProfile struct {
	Id              int               `gorm:"primaryKey;auto_increment"`
	MemberId        int               `gorm:"type:int"`
	ProfilePage     string            `gorm:"type:varchar(255)"`
	ProfileName     string            `gorm:"type:varchar(255)"`
	ProfileSlug     string            `gorm:"type:varchar(255)"`
	CompanyLogo     string            `gorm:"type:varchar(255)"`
	CompanyName     string            `gorm:"type:varchar(255)"`
	CompanyLocation string            `gorm:"type:varchar(255)"`
	About           string            `gorm:"type:varchar(255)"`
	Linkedin        string            `gorm:"type:varchar(255)"`
	Website         string            `gorm:"type:varchar(255)"`
	Twitter         string            `gorm:"type:varchar(255)"`
	SeoTitle        string            `gorm:"type:varchar(255)"`
	SeoDescription  string            `gorm:"type:varchar(255)"`
	SeoKeyword      string            `gorm:"type:varchar(255)"`
	MemberDetails   datatypes.JSONMap `json:"memberDetails" gorm:"column:member_details;type:jsonb"`
	ClaimStatus     int               `gorm:"DEFAULT:0;type:integer"`
	CreatedBy       int               `gorm:"type:int"`
	CreatedOn       time.Time         `gorm:"type:datetime"`
	ModifiedBy      int               `gorm:"DEFAULT:NULL;type:int"`
	ModifiedOn      time.Time         `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted       int               `gorm:"DEFAULT:0"`
	DeletedBy       int               `gorm:"DEFAULT:NULL;type:int"`
	DeletedOn       time.Time         `gorm:"type:datetime;DEFAULT:NULL"`
	TenantId        int               `gorm:"type:int"`
}

func MigrationTables(db *gorm.DB) {

	if err := db.AutoMigrate(
		&TblModule{},
		&TblModulePermission{},
		&TblRolePermission{},
		&TblUser{},
		&TblMemberGroup{},
		&TblMember{},
		&TblMemberProfile{},
	); err != nil {

		panic(err)

	}
}
