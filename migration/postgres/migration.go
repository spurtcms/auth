package postgres

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type TblUser struct {
	Id                int       `gorm:"primaryKey;type:serial"`
	Uuid              string    `gorm:"type:character varying"`
	FirstName         string    `gorm:"type:character varying"`
	LastName          string    `gorm:"type:character varying"`
	RoleId            int       `gorm:"type:integer"`
	Email             string    `gorm:"type:character varying"`
	Username          string    `gorm:"type:character varying"`
	Password          string    `gorm:"type:character varying"`
	MobileNo          string    `gorm:"type:character varying"`
	IsActive          int       `gorm:"type:integer"`
	ProfileImage      string    `gorm:"type:character varying"`
	ProfileImagePath  string    `gorm:"type:character varying"`
	DataAccess        int       `gorm:"type:integer"`
	DefaultLanguageId int       `gorm:"type:integer"`
	CreatedOn         time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy         int       `gorm:"type:integer"`
	ModifiedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy        int       `gorm:"DEFAULT:NULL;type:integer"`
	LastLogin         time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted         int       `gorm:"type:integer"`
	DeletedOn         time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy         int       `gorm:"DEFAULT:NULL"`
	TenantId          int       `gorm:"type:integer"`
}

type TblModule struct {
	Id               int       `gorm:"primaryKey;type:serial"`
	ModuleName       string    `gorm:"type:character varying"`
	IsActive         int       `gorm:"type:integer"`
	DefaultModule    int       `gorm:"type:integer"`
	ParentId         int       `gorm:"type:integer"`
	IconPath         string    `gorm:"type:character varying"`
	AssignPermission int       `gorm:"type:integer"`
	Description      string    `gorm:"type:character varying"`
	OrderIndex       int       `gorm:"type:integer"`
	CreatedBy        int       `gorm:"type:integer"`
	CreatedOn        time.Time `gorm:"type:timestamp without time zone"`
	MenuType         string    `gorm:"type:character varying"`
	GroupFlg         int       `gorm:"type:integer"`
	TenantId         int       `gorm:"type:integer"`
}

type TblModulePermission struct {
	Id                   int       `gorm:"primaryKey;type:serial"`
	RouteName            string    `gorm:"type:character varying;unique"`
	DisplayName          string    `gorm:"type:character varying"`
	SlugName             string    `gorm:"type:character varying"`
	Description          string    `gorm:"type:character varying"`
	ModuleId             int       `gorm:"type:integer"`
	FullAccessPermission int       `gorm:"type:integer"`
	ParentId             int       `gorm:"type:integer"`
	AssignPermission     int       `gorm:"type:integer"`
	BreadcrumbName       string    `gorm:"type:character varying"`
	OrderIndex           int       `gorm:"type:integer"`
	CreatedBy            int       `gorm:"type:integer"`
	CreatedOn            time.Time `gorm:"type:timestamp without time zone"`
	ModifiedBy           int       `gorm:"DEFAULT:NULL"`
	ModifiedOn           time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId             int       `gorm:"type:integer"`
}

type TblRolePermission struct {
	Id           int       `gorm:"primaryKey;type:serial"`
	RoleId       int       `gorm:"type:integer"`
	PermissionId int       `gorm:"type:integer"`
	CreatedBy    int       `gorm:"type:integer"`
	CreatedOn    time.Time `gorm:"type:timestamp without time zone"`
	TenantId     int       `gorm:"type:integer"`
}

type TblMemberGroup struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	Name        string    `gorm:"type:character varying"`
	Slug        string    `gorm:"type:character varying"`
	Description string    `gorm:"type:character varying"`
	IsActive    int       `gorm:"type:integer"`
	IsDeleted   int       `gorm:"type:integer"`
	CreatedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy   int       `gorm:"type:integer"`
	ModifiedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL"`
	TenantId    int       `gorm:"type:integer"`
}

type TblMember struct {
	Id               int       `gorm:"primaryKey;auto_increment;type:serial"`
	Uuid             string    `gorm:"type:character varying"`
	FirstName        string    `gorm:"type:character varying"`
	LastName         string    `gorm:"type:character varying"`
	Email            string    `gorm:"type:character varying"`
	MobileNo         string    `gorm:"type:character varying"`
	IsActive         int       `gorm:"type:integer"`
	ProfileImage     string    `gorm:"type:character varying"`
	ProfileImagePath string    `gorm:"type:character varying"`
	LastLogin        int       `gorm:"type:integer"`
	MemberGroupId    int       `gorm:"type:integer"`
	Password         string    `gorm:"type:character varying"`
	Username         string    `gorm:"DEFAULT:NULL"`
	Otp              int       `gorm:"DEFAULT:NULL"`
	OtpExpiry        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	LoginTime        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted        int       `gorm:"type:integer"`
	DeletedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"DEFAULT:NULL"`
	CreatedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy        int       `gorm:"type:integer"`
	ModifiedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"DEFAULT:NULL"`
	TenantId         int       `gorm:"type:integer"`
}

type TblMemberProfile struct {
	Id              int               `gorm:"primaryKey;auto_increment;type:serial"`
	MemberId        int               `gorm:"type:integer"`
	ProfilePage     string            `gorm:"type:character varying"`
	ProfileName     string            `gorm:"type:character varying"`
	ProfileSlug     string            `gorm:"type:character varying"`
	CompanyLogo     string            `gorm:"type:character varying"`
	CompanyName     string            `gorm:"type:character varying"`
	CompanyLocation string            `gorm:"type:character varying"`
	About           string            `gorm:"type:character varying"`
	Linkedin        string            `gorm:"type:character varying"`
	Website         string            `gorm:"type:character varying"`
	Twitter         string            `gorm:"type:character varying"`
	SeoTitle        string            `gorm:"type:character varying"`
	SeoDescription  string            `gorm:"type:character varying"`
	SeoKeyword      string            `gorm:"type:character varying"`
	MemberDetails   datatypes.JSONMap `json:"memberDetails" gorm:"column:member_details;type:jsonb"`
	ClaimStatus     int               `gorm:"DEFAULT:0;type:integer"`
	CreatedBy       int               `gorm:"type:integer"`
	CreatedOn       time.Time         `gorm:"type:timestamp without time zone"`
	ModifiedBy      int               `gorm:"DEFAULT:NULL"`
	ModifiedOn      time.Time         `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	IsDeleted       int               `gorm:"DEFAULT:0"`
	DeletedBy       int               `gorm:"DEFAULT:NULL"`
	DeletedOn       time.Time         `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	TenantId        int               `gorm:"type:integer"`
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