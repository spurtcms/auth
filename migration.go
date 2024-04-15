package auth

import (
	"time"

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
}

type TblRolePermission struct {
	Id           int       `gorm:"primaryKey;type:serial"`
	RoleId       int       `gorm:"type:integer"`
	PermissionId int       `gorm:"type:integer"`
	CreatedBy    int       `gorm:"type:integer"`
	CreatedOn    time.Time `gorm:"type:timestamp without time zone"`
}

func Migration(db *gorm.DB) {

	if err := db.AutoMigrate(
		&TblModule{},
		&TblModulePermission{},
		&TblRolePermission{},
		&TblUser{},

	); err != nil {

		panic(err)

	}
}
