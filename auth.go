package auth

import (
	cr "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/spurtcms/auth/migration"
	"golang.org/x/crypto/bcrypt"
)

// AuthSetup used initialize auth configruation
func AuthSetup(conf Config) *Auth {

	migration.AutoMigration(conf.DB, conf.DataBaseType)

	return &Auth{
		UserId:     conf.UserId,
		ExpiryTime: conf.ExpiryTime,
		SecretKey:  conf.SecretKey,
		DB:         conf.DB,
		ExpiryFlg:  conf.ExpiryFlg,
		RoleId:     conf.RoleId,
	}

}

// Check UserName Password - userlogin
func (auth *Auth) Checklogin(Username string, Password string, tenantid string) (string, int, error) {

	username := Username

	password := Password

	user, err := Authmodel.CheckLogin(username, password, auth.DB)

	if err != nil {

		return "", 0, err

	}

	if user.IsActive == 0 {

		return "", 0, ErrorInactive
	}

	passerr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if passerr != nil || passerr == bcrypt.ErrMismatchedHashAndPassword {

		return "", 0, ErrorPassword

	}

	auth.UserId = user.Id

	auth.RoleId = user.RoleId

	token, err := auth.CreateToken()

	if err != nil {

		return "", 0, err
	}

	return token, user.Id, nil
}

// CreateToken creates a token
func (auth *Auth) CreateToken() (string, error) {

	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = auth.UserId
	atClaims["role_id"] = auth.RoleId
	atClaims["expiry_time"] = time.Now().UTC().Add(3 * 30 * 24 * time.Hour)
	atClaims["login_type"] = ""

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	return token.SignedString([]byte(auth.SecretKey))
}

// verify token
func (auth *Auth) VerifyToken(token string, secret string) (userid int, loginType string, err error) {

	Claims := jwt.MapClaims{}

	tkn, err := jwt.ParseWithClaims(token, Claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println(err)
			return 0, "", ErrorToken
		}

		return 0, "", ErrorToken
	}

	if !tkn.Valid {
		fmt.Println(tkn)
		return 0, "", ErrorToken
	}

	if auth.ExpiryFlg {
		expiryTime := Claims["expiry_time"]
		expiryTimeStr, ok := expiryTime.(string)
		if !ok {
			fmt.Println("Could not convert interface to string")
			return 0, "", ErrorConvertTime
		}
		expiryTimes, err := time.Parse(time.RFC3339, expiryTimeStr)
		if err != nil {
			fmt.Println("Could not parse expiry time:", err)
			return 0, "", err
		}
		fmt.Println("expiryTimes:", expiryTimes)
		fmt.Println("currentTime:", time.Now().UTC())
		if time.Now().UTC().After(expiryTimes) {
			fmt.Println("expired:")
			return 0, "", ErrorTokenExpiry

		}
	}

	usrid := Claims["user_id"]

	var logintypee string
	if Claims["login_type"] != nil {
		logintypee = Claims["login_type"].(string)
	}

	auth.AuthFlg = true
	return int(usrid.(float64)), logintypee, nil
}

func (auth *Auth) CheckMemberLogin(memberlogin MemberLoginCheck, tenantid string) (TblMember, error) {

	var member TblMember

	var err error

	currentTime := time.Now().UTC()

	if memberlogin.EmailwithPassword {

		member, err = Authmodel.CheckMemberLoginWithEmail(memberlogin.Email, memberlogin.Username, auth.DB, tenantid)

		if err != nil {

			return TblMember{}, err
		}

		passerr := bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(memberlogin.Password))

		if passerr != nil && passerr != bcrypt.ErrMismatchedHashAndPassword {
			return TblMember{}, passerr
		}
		if passerr == bcrypt.ErrMismatchedHashAndPassword {
			return TblMember{}, ErrorPassword
		}

	} else if memberlogin.UsernameWithPassword {

		member, err = Authmodel.CheckMemberLoginWithEmail(memberlogin.Email, memberlogin.Username, auth.DB, tenantid)

		if err != nil {

			return TblMember{}, err
		}

		passerr := bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(memberlogin.Password))

		if passerr != nil && passerr != bcrypt.ErrMismatchedHashAndPassword {
			return TblMember{}, passerr
		}
		if passerr == bcrypt.ErrMismatchedHashAndPassword {
			return TblMember{}, ErrorPassword
		}

	} else if memberlogin.EmailWithOTP {

		member, err = Authmodel.CheckEmailWithOtp(memberlogin.Email, auth.DB, tenantid)

		if err != nil {
			return TblMember{}, err
		}

		if member.IsActive == 0 {
			return TblMember{}, ErrorInactiveMember
		}

		if member.Otp != memberlogin.OTP {
			return TblMember{}, ErrorInvalidOTP
		}

		if currentTime.After(member.OtpExpiry) {
			return TblMember{}, ErrorOtpExpiry
		}

	} else if memberlogin.UsernameWithOTP {

		member, err = Authmodel.CheckUsernameWithOtp(memberlogin.Email, auth.DB, tenantid)

		if err != nil {
			return TblMember{}, err
		}

		if member.Otp != memberlogin.OTP {
			return TblMember{}, ErrorInvalidOTP
		}

		if currentTime.After(member.OtpExpiry) {
			return TblMember{}, ErrorOtpExpiry
		}
	}

	return member, nil

}

// member token
func (auth *Auth) GenerateMemberToken(memberid int, loginType string, secretKey string, tenantid string) (token string, err error) {

	var MemberDetails TblMember

	if err := Authmodel.GetMemberDetailsByMemberId(&MemberDetails, memberid, auth.DB, tenantid); err != nil {

		return "", err
	}

	token, tokenerr := CreateMemberToken(MemberDetails.Id, MemberDetails.MemberGroupId, secretKey, loginType)
	if tokenerr != nil {

		return "", err
	}

	return token, nil
}

/*Create meber token*/
func CreateMemberToken(userid, roleId int, secretKey string, loginType string) (string, error) {

	atClaims := jwt.MapClaims{}
	atClaims["member_id"] = userid
	atClaims["group_id"] = roleId
	atClaims["expiry_time"] = time.Now().Add(168 * time.Hour).Unix()
	atClaims["login_type"] = loginType
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	return token.SignedString([]byte(secretKey))
}

// verify token
func (auth *Auth) MemberVerifyToken(token string, secret string) (memberid int, groupid int, loginType string, err error) {

	Claims := jwt.MapClaims{}
	tkn, err := jwt.ParseWithClaims(token, Claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println(err)
			return 0, 0, "", ErrorToken
		}

		return 0, 0, "", ErrorToken
	}

	if !tkn.Valid {
		fmt.Println(tkn)
		return 0, 0, "", ErrorToken
	}

	if auth.ExpiryFlg {
		expiryTime := Claims["expiry_time"]
		t, ok := expiryTime.(time.Time)
		if !ok {
			fmt.Println("Could not convert interface to time.Time")
			return 0, 0, "", ErrorConvertTime
		}
		if t.After(time.Now().UTC()) {
			return 0, 0, "", ErrorTokenExpiry
		}
	}

	usrid := Claims["member_id"]
	grpid := Claims["group_id"]

	var logintypee string
	if Claims["login_type"] != nil {
		logintypee = Claims["login_type"].(string)
	}

	auth.AuthFlg = true
	return int(usrid.(float64)), int(grpid.(float64)), logintypee, nil
}

// update otp
func (auth *Auth) UpdateMemberOTP(otp OTP, tenantid string) (int, time.Time, error) {

	//generate otp
	generateotp := func() int {
		const digits = "0123456789"
		var otpp string
		for i := 0; i < otp.Length; i++ {
			randomInt, err := cr.Int(cr.Reader, big.NewInt(int64(len(digits))))
			if err == nil {
				otpp += string(digits[randomInt.Int64()])
			}
		}
		otpint, _ := strconv.Atoi(otpp)
		return otpint
	}

	genOtp := generateotp()

	otp_expiry := time.Now().UTC().Add(otp.Duration)

	otp_expiry_time := otp_expiry.Format("2006-01-02 15:04:05")

	if err := Authmodel.UpdateMemberOtp(otp.MemberId, genOtp, otp_expiry_time, auth.DB, tenantid); err != nil {
		return 0, time.Time{}, err
	}

	return genOtp, otp_expiry, nil
}

func (auth *Auth) OtpLoginVerification(otp int, email string, tenantid string) (Tbluser, string, bool, error) {

	userdet, err := Authmodel.GetUserByEmail(email, auth.DB, tenantid)

	var isNewUser bool

	if err != nil {

		return Tbluser{}, "", false, fmt.Errorf("")
	}

	currentTime := time.Now().UTC()

	if userdet.Otp != otp {

		return Tbluser{}, "", false, ErrorInvalidOTP
	}

	if !userdet.OtpExpiry.After(currentTime) {

		return Tbluser{}, "", false, ErrorOtpExpiry
	}

	if tenantid == "" {

		isNewUser = true

		// tenantID, err := Authmodel.CreateTenantid(&TblMstrTenant{TenantId: "1"}, auth.DB)

		// if err != nil {
		// 	fmt.Println("Tenant ID not created:", err)
		// 	return Tbluser{}, "", false, nil
		// }

		uuid := (uuid.New()).String()

		arr := strings.Split(uuid, "-")

		UniqueId := arr[len(arr)-1]

		err = Authmodel.UpdateTenantId(userdet.Id, UniqueId, auth.DB)

		if err != nil {

			return Tbluser{}, "", false, nil
		}

		userdet.TenantId = UniqueId

	}
	var subdomain string

	if userdet.Subdomain == "" {

		subdomain = strings.ToLower(userdet.FirstName) + strconv.Itoa(userdet.Id)

		err1 := Authmodel.UpdateSubDomain(userdet.Id, subdomain, auth.DB)

		if err1 != nil {

			return Tbluser{}, "", false, nil
		}
		userdet.Subdomain = subdomain
	}

	auth.UserId = userdet.Id

	auth.RoleId = userdet.RoleId

	token, _ := auth.CreateToken()

	return userdet, token, isNewUser, nil

}

func (auth *Auth) UpdateUserOTP(user Tbluser) (Tbluser, error) {

	ExpirationTime := time.Now().UTC().Add(5 * time.Minute)

	user.OtpExpiry = &ExpirationTime

	user.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Authmodel.UpdateUserOtp(user, auth.DB)

	if err != nil {

		return Tbluser{}, err
	}

	return user, nil

}

func (auth *Auth) CheckWebAuth(login *SocialLogin) (string, Tbluser, bool, error) {

	userinfo, _ := Authmodel.GetUserByEmail(login.Email, auth.DB, "")

	createdon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	var isNewUser bool

	if userinfo.Email != login.Email {

		isNewUser = true

		roledetails, _ := Authmodel.CheckRoleByName("admin", auth.DB)

		uvuid := (uuid.New()).String()

		var newRoleId int

		if roledetails.Id == 0 {

			newrole, _ := Authmodel.CreateRole(Tblrole{Name: "Admin", Description: "Admin role type", IsActive: 1, CreatedOn: createdon, CreatedBy: 1, Slug: "admin"}, auth.DB)

			newRoleId = newrole.Id

		} else {

			newRoleId = roledetails.Id
		}

		newUser := Tbluser{
			FirstName:         login.FirstName,
			LastName:          login.LastName,
			Email:             login.Email,
			Username:          login.GivenName,
			IsActive:          1,
			CreatedOn:         createdon,
			DefaultLanguageId: 1,
			Uuid:              uvuid,
			RoleId:            newRoleId,
			GoTemplateDefault: 1,
		}

		userinfo, _ = Authmodel.CreateUser(&newUser, auth.DB)

		uuid := (uuid.New()).String()

		arr := strings.Split(uuid, "-")

		UniqueId := arr[len(arr)-1]

		err := Authmodel.UpdateTenantId(userinfo.Id, UniqueId, auth.DB)

		if err != nil {

			return "", Tbluser{}, false, nil
		}

		userinfo.TenantId = UniqueId

	}

	var subdomain string

	if userinfo.Subdomain == "" {

		subdomain = strings.ToLower(userinfo.FirstName) + strconv.Itoa(userinfo.Id)

		err1 := Authmodel.UpdateSubDomain(userinfo.Id, subdomain, auth.DB)

		if err1 != nil {

			return "", Tbluser{}, false, nil
		}
		userinfo.Subdomain = subdomain
	}

	if userinfo.IsActive == 0 {

		return "", Tbluser{}, false, ErrorInactive
	}

	auth.UserId = userinfo.Id

	auth.RoleId = userinfo.RoleId

	token, err := auth.CreateToken()

	if err != nil {

		return "", Tbluser{}, false, nil
	}

	return token, userinfo, isNewUser, nil
}

func (auth *Auth) CheckOtpLogin(email string) (Tbluser, error) {

	var userdetails, _ = Authmodel.GetUserByEmail(email, auth.DB, "")

	createdon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	if email != userdetails.Email {

		roledetails, _ := Authmodel.CheckRoleByName("admin", auth.DB)

		uvuid := (uuid.New()).String()

		var newRoleId int

		if roledetails.Id == 0 {

			newrole, _ := Authmodel.CreateRole(Tblrole{Name: "Admin", Description: "Admin role type", IsActive: 1, CreatedOn: createdon, CreatedBy: 1, Slug: "admin"}, auth.DB)

			newRoleId = newrole.Id

		} else {

			newRoleId = roledetails.Id
		}

		FirstName, LastName := strings.Split(email, "@")[0], ""

		newUser := Tbluser{
			FirstName:         FirstName,
			LastName:          LastName,
			Email:             email,
			Username:          FirstName,
			IsActive:          1,
			CreatedOn:         createdon,
			DefaultLanguageId: 1,
			Uuid:              uvuid,
			RoleId:            newRoleId,
			GoTemplateDefault: 1,
		}

		userdetails, _ = Authmodel.CreateUser(&newUser, auth.DB)

	}
	if userdetails.IsActive == 0 {

		return Tbluser{}, ErrorInactive
	}

	otp := generateOTP()

	var loginuser Tbluser

	loginuser.Id = userdetails.Id

	loginuser.Email = userdetails.Email

	loginuser.ModifiedBy = userdetails.Id

	loginuser.TenantId = userdetails.TenantId

	loginuser.Otp, _ = strconv.Atoi(otp)

	ExpirationTime := time.Now().UTC().Add(5 * time.Minute)

	loginuser.OtpExpiry = &ExpirationTime

	loginuser.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Authmodel.UpdateUserOtp(loginuser, auth.DB)

	if err != nil {

		return Tbluser{}, err
	}

	userdetails.Otp, err = strconv.Atoi(otp)

	userdetails.OtpExpiry = &ExpirationTime

	if err != nil {

		return Tbluser{}, err
	}

	return userdetails, nil
}

func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	otp := fmt.Sprintf("%6d", rand.Intn(900000)+100000)
	return otp
}

func (auth *Auth) UpdateS3FolderName(tenantid string, userid int, s3path string) error {

	err := Authmodel.UpdateS3FolderName(tenantid, userid, s3path, auth.DB)

	if err != nil {

		return err
	}
	return nil
}
