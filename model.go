package auth

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Tbluser struct {
	Id                   int       `gorm:"column:id"`
	Uuid                 string    `gorm:"column:uuid"`
	FirstName            string    `gorm:"column:first_name"`
	LastName             string    `gorm:"column:last_name"`
	RoleId               int       `gorm:"column:role_id"`
	Email                string    `gorm:"column:email"`
	Username             string    `gorm:"column:username"`
	Password             string    `gorm:"column:password"`
	MobileNo             string    `gorm:"column:mobile_no"`
	IsActive             int       `gorm:"column:is_active"`
	ProfileImage         string    `gorm:"column:profile_image"`
	ProfileImagePath     string    `gorm:"column:profile_image_path"`
	StorageType          string    `gorm:"column:storage_type"`
	DataAccess           int       `gorm:"column:data_access"`
	CreatedOn            time.Time `gorm:"column:created_on"`
	CreatedBy            int       `gorm:"column:created_by"`
	ModifiedOn           time.Time `gorm:"column:modified_on;DEFAULT:NULL"`
	ModifiedBy           int       `gorm:"column:modified_by;DEFAULT:NULL"`
	LastLogin            time.Time `gorm:"column:last_login;DEFAULT:NULL"`
	IsDeleted            int       `gorm:"column:is_deleted"`
	DeletedOn            time.Time `gorm:"column:deleted_on;DEFAULT:NULL"`
	DeletedBy            int       `gorm:"column:deleted_by;DEFAULT:NULL"`
	ModuleName           string    `gorm:"-"`
	RouteName            string    `gorm:"-:migration;<-:false"`
	DisplayName          string    `gorm:"-:migration;<-:false"`
	Description          string    `gorm:"-"`
	ModuleId             int       `gorm:"-:migration;<-:false"`
	PermissionId         int       `gorm:"-"`
	FullAccessPermission int       `gorm:"-:migration;<-:false"`
	RoleName             string    `gorm:"-:migration;<-:false"`
	DefaultLanguageId    int       `gorm:"column:default_language_id"`
	NameString           string    `gorm:"-"`
	TenantId             int
	Otp                  int       `gorm:"column:otp"`
	OtpExpiry            time.Time `gorm:"column:otp_expiry"`
}

type SocialLogin struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	GivenName string `json:"given_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type TblMstrTenant struct {
	Id        int       `gorm:"primaryKey;auto_increment;type:serial"`
	TenantId  int       `gorm:"type:integer"`
	DeletedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted int       `gorm:"type:integer;DEFAULT:0"`
}
type Tblrole struct {
	Id          int       `gorm:"column:id"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	Slug        string    `gorm:"column:slug"`
	IsActive    int       `gorm:"column:is_active"`
	IsDeleted   int       `gorm:"column:is_deleted"`
	CreatedOn   time.Time `gorm:"column:created_on"`
	CreatedBy   int       `gorm:"column:created_by"`
	ModifiedOn  time.Time `gorm:"column:modified_on;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"column:modified_by;DEFAULT:NULL"`
	CreatedDate string    `gorm:"-:migration;<-:false"`
	User        []Tbluser `gorm:"-"`
	TenantId    int       `gorm:"column:tenant_id;DEFAULT:NULL"`
}
type TblGraphqlSettings struct {
	Id          int
	TokenName   string
	Description string
	Duration    string
	CreatedBy   int `gorm:"DEFAULT:NULL"`
	CreatedOn   time.Time
	ModifiedBy  int       `gorm:"DEFAULT:NULL"`
	ModifiedOn  time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy   int       `gorm:"DEFAULT:NULL"`
	DeletedOn   time.Time `gorm:"DEFAULT:NULL"`
	IsDeleted   int       `gorm:"DEFAULT:0"`
	Token       string
	IsDefault   int       `gorm:"DEFAULT:0"`
	ExpiryTime  time.Time `gorm:"DEFAULT:NULL"`
	TenantId    int
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

type TblStorageType struct {
	Id           int
	Local        string
	Aws          datatypes.JSONMap `gorm:"type:jsonb"`
	Azure        datatypes.JSONMap `gorm:"type:jsonb"`
	Drive        datatypes.JSONMap `gorm:"type:jsonb"`
	SelectedType string
}

type authmodel struct{}

var Authmodel authmodel

// soft delete check
func IsDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("is_deleted = 0")
}

// check db userlogin
func (auth authmodel) CheckLogin(username string, Password string, db *gorm.DB) (user Tbluser, err error) {

	if err := db.Table("tbl_users").Where("username = ? and  is_deleted = 0", username).First(&user).Error; err != nil {

		return Tbluser{}, err

	}

	return user, nil
}

// check email with password
func (auth authmodel) CheckMemberLoginWithEmail(email string, username string, DB *gorm.DB, tenantid int) (member TblMember, err error) {

	if email != "" {

		if err := DB.Model(TblMember{}).Where("is_deleted =0 and email=? and (tenant_id is NULL or tenant_id=?)", email, tenantid).First(&member).Error; err != nil {

			return TblMember{}, err
		}

	} else if username != "" {

		if err := DB.Model(TblMember{}).Where("is_deleted =0 and username=? and (tenant_id is NULL or tenant_id=?)", username, tenantid).First(&member).Error; err != nil {

			return TblMember{}, err
		}

	}

	return member, nil
}

func (auth authmodel) CheckEmailWithOtp(email string, DB *gorm.DB, tenantid int) (member TblMember, err error) {

	if err := DB.Model(TblMember{}).Where("is_deleted = 0 and email = ? and (tenant_id is NULL or tenant_id=?)", email, tenantid).First(&member).Error; err != nil {

		return TblMember{}, err
	}

	return member, nil
}

func (auth authmodel) CheckUsernameWithOtp(username string, DB *gorm.DB, tenantid int) (member TblMember, err error) {

	if err := DB.Model(TblMember{}).Where("is_deleted = 0 and username = ? and (tenant_id is NULL or tenant_id=?)", username, tenantid).First(&member).Error; err != nil {

		return TblMember{}, err
	}

	return member, nil
}

func (auth authmodel) UpdateMemberOtp(id int, otp int, otpExpiry string, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_members").Where("id=? and (tenant_id is NULL or tenant_id=?)", id, tenantid).Updates(map[string]interface{}{
		"otp": otp, "otp_expiry": otpExpiry,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (auth authmodel) GetMemberDetailsByMemberId(MemberDetails *TblMember, memberId int, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_members").Where("is_deleted=0 and id = ? and (tenant_id is NULL or tenant_id=?)", memberId, tenantid).First(&MemberDetails).Error; err != nil {
		return err
	}

	return nil
}
func (auth authmodel) GetUserByEmail(email string, DB *gorm.DB, tenantid int) (user Tbluser, err error) {

	query := DB.Debug().Table("tbl_users").Where("is_deleted = 0 AND email = ?", email)

	if tenantid != -1 {
		query = query.Where("(tenant_id IS NULL OR tenant_id = ?)", tenantid)
	}

	if err := query.First(&user).Error; err != nil {
		return Tbluser{}, err
	}

	return user, nil
}
func (auth authmodel) UpdateUserOtp(user Tbluser, DB *gorm.DB) error {

	result := DB.Table("tbl_users").Where("id = ? and (tenant_id is NULL or tenant_id=?)", user.Id, user.TenantId).UpdateColumns(map[string]interface{}{"modified_on": user.ModifiedOn, "modified_by": user.Id, "otp": user.Otp, "otp_expiry": user.OtpExpiry})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (auth authmodel) CheckRoleByName(rolename string, DB *gorm.DB) (role Tblrole, err error) {

	if err := DB.Debug().Table("tbl_roles").Where("slug =? and is_deleted=0 ", rolename).Find(&role).Error; err != nil {

		return Tblrole{}, err
	}

	return role, err
}

func (auth authmodel) CreateRole(role Tblrole, DB *gorm.DB) (roledetails Tblrole, err error) {

	if err := DB.Debug().Table("tbl_roles").Create(&role).Error; err != nil {

		return Tblrole{}, err
	}

	return role, nil
}
func (auth authmodel) CreateUser(user *Tbluser, DB *gorm.DB) (team Tbluser, terr error) {

	if err := DB.Debug().Table("tbl_users").Create(&user).Error; err != nil {

		return Tbluser{}, err

	}

	return *user, nil
}

func (auth authmodel) GetRoleById(roleid int, DB *gorm.DB) (role Tblrole, err error) {

	if err := DB.Table("tbl_roles").Where("id=?", roleid).First(&role).Error; err != nil {

		return Tblrole{}, err

	}

	return role, nil
}
func (auth authmodel) CreateTenantid(user *TblMstrTenant, DB *gorm.DB) (int, error) {
	result := DB.Table("tbl_mstr_tenants").Create(user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.Id, nil
}

func (auth authmodel) UpdateTenantId(UserId int, Tenantid int, DB *gorm.DB) error {

	result := DB.Table("tbl_users").Where("id = ?", UserId).Update("tenant_id", Tenantid)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
func (auth authmodel) CreateTenantApiToken(DB *gorm.DB, tokenDetails *TblGraphqlSettings) error {
	if err := DB.Debug().Create(&tokenDetails).Error; err != nil {
		return err
	}
	return nil
}

func (auth authmodel) GetStorageValue(DB *gorm.DB) (tblstorgetype TblStorageType, err error) {

	if err := DB.Model(TblStorageType{}).Where("id = 1 ").First(&tblstorgetype).Error; err != nil {

		return TblStorageType{}, err
	}

	return tblstorgetype, nil
}

func (auth authmodel) UpdateS3FolderName(tenantId, userId int, s3FolderPath string, DB *gorm.DB) error {

	result := DB.Table("tbl_users").Where("id = ?", userId).Update("s3_folder_name", s3FolderPath)
	if result.Error != nil {
		return result.Error
	}

	result = DB.Table("tbl_mstr_tenants").Where("id = ?", tenantId).Update("s3_storage_path", s3FolderPath)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (auth authmodel) CreateTenantDefaultData(userId, tenantId int, db *gorm.DB) error {

	file, err := os.Open("tenant-defaults.sql")

	if err != nil {

		return err
	}

	scanner := bufio.NewScanner(file)

	var (
		defChildCatId, defParentCatId, tagId int
		defBlockIds                          []int
		blockDynamicRetrieve                 bool
	)

	for scanner.Scan() {

		query := scanner.Text()

		if len(query) > 0 && !strings.HasPrefix(query, "--") {

			currentTime, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

			parts := strings.Fields(fmt.Sprint(currentTime))

			lowerQuery := strings.ToLower(query)

			switch {

			case strings.Contains(lowerQuery, "tbl_categories") && strings.Contains(lowerQuery, "insert into"):

				if err := db.Table("tbl_categories").Where("is_deleted=0 and tenant_id=? and parent_id=0", tenantId).Pluck("id", &defParentCatId).Error; err != nil {

					return err
				}

				if strings.Contains(lowerQuery, "pid") {

					query = strings.ReplaceAll(query, "pid", strconv.Itoa(defParentCatId))

				}

			case strings.Contains(lowerQuery, "tbl_channel_categories") && strings.Contains(lowerQuery, "insert into"):

				if err := db.Table("tbl_categories").Where("is_deleted=0 and tenant_id=? and parent_id=?", tenantId, defParentCatId).Pluck("id", &defChildCatId).Error; err != nil {

					return err
				}

				if strings.Contains(lowerQuery, "mapcat") {

					mapCategories := fmt.Sprintf("%v,%v", defParentCatId, defChildCatId)

					query = strings.ReplaceAll(query, "mapcat", mapCategories)

				}

			case (strings.Contains(lowerQuery, "tbl_block_tags") || strings.Contains(lowerQuery, "tbl_block_collections")) && strings.Contains(lowerQuery, "insert into") && !blockDynamicRetrieve:

				if err := db.Table("tbl_block_mstr_tags").Select("id").Where("tenant_id=?", tenantId).Pluck("id", &tagId).Error; err != nil {

					return err
				}

				if err := db.Table("tbl_blocks").Where("is_deleted=0 and tenant_id=?", tenantId).Find(&defBlockIds).Error; err != nil {

					return err
				}

				blockDynamicRetrieve = true

			}

			timeStamp := fmt.Sprintf("%s %s", parts[0], parts[1])

			replacer := strings.NewReplacer("uid", strconv.Itoa(userId), "tid", strconv.Itoa(tenantId), "time", timeStamp)

			finalQuery := replacer.Replace(query)

			switch {

			default:

				if err := db.Exec(finalQuery).Error; err != nil {

					return err
				}

			case strings.Contains(lowerQuery, "tbl_block_tags") && strings.Contains(lowerQuery, "insert into"):

				for _, v := range defBlockIds {

					replacer := strings.NewReplacer("blid", strconv.Itoa(v), "tagid", strconv.Itoa(tagId))

					modQuery := replacer.Replace(finalQuery)

					if err := db.Exec(modQuery).Error; err != nil {

						return err
					}
				}

			case strings.Contains(lowerQuery, "tbl_block_collections") && strings.Contains(lowerQuery, "insert into"):

				for _, v := range defBlockIds {

					replacer := strings.NewReplacer("blid", strconv.Itoa(v), "tagid", strconv.Itoa(tagId))

					modQuery := replacer.Replace(finalQuery)

					if err := db.Exec(modQuery).Error; err != nil {

						return err
					}
				}
			}

		}

	}

	return nil

}
