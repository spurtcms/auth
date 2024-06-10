package auth

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
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
	}

}

// Check UserName Password - userlogin
func (auth *Auth) Checklogin(Username string, Password string) (string, int, error) {

	username := Username

	password := Password

	user, err := Authmodel.CheckLogin(username, password, auth.DB)

	if err != nil {

		fmt.Println(err)

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
	atClaims["expiry_time"] = time.Now().UTC().Add(time.Duration(auth.ExpiryTime) * time.Hour)
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
		t, ok := expiryTime.(time.Time)
		if !ok {
			fmt.Println("Could not convert interface to time.Time")
			return 0, "", ErrorConvertTime
		}
		if t.After(time.Now().UTC()) {

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

func (auth *Auth) CheckMemberLogin(memberlogin MemberLoginCheck) error {

	var member TblMember

	var err error

	if memberlogin.EmailwithPassword {

		member, err = Authmodel.CheckMemberLoginWithEmail(memberlogin.Email, memberlogin.Username, auth.DB)

		if err != nil {

			return err
		}

		passerr := bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(memberlogin.Password))

		if passerr != nil && passerr != bcrypt.ErrMismatchedHashAndPassword {
			return passerr
		}
		if passerr == bcrypt.ErrMismatchedHashAndPassword {
			return ErrorPassword
		}

	} else if memberlogin.UsernameWithPassword {

		member, err = Authmodel.CheckMemberLoginWithEmail(memberlogin.Email, memberlogin.Username, auth.DB)

		if err != nil {

			return err
		}

		passerr := bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(memberlogin.Password))

		if passerr != nil && passerr != bcrypt.ErrMismatchedHashAndPassword {
			return passerr
		}
		if passerr == bcrypt.ErrMismatchedHashAndPassword {
			return ErrorPassword
		}

	} else if memberlogin.EmailWithOTP {

		member, err = Authmodel.CheckEmailWithOtp(memberlogin.Email, auth.DB)

		if err != nil {
			return err
		}

		if member.Otp != memberlogin.OTP {
			return ErrorInvalidOTP
		}

	} else if memberlogin.UsernameWithOTP {

		member, err = Authmodel.CheckUsernameWithOtp(memberlogin.Email, auth.DB)

		if err != nil {
			return err
		}

		if member.Otp != memberlogin.OTP {
			return ErrorInvalidOTP
		}
	}

	return ErrorMemberLogin

}

// member token
func (auth *Auth) GenerateMemberToken(memberid int, loginType string, secretKey string) (token string, err error) {

	var MemberDetails TblMember

	if err := Authmodel.GetMemberDetailsByMemberId(&MemberDetails, memberid, auth.DB); err != nil {

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
	atClaims["expiry_time"] = time.Now().Add(2 * time.Hour).Unix()
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
func (auth *Auth) UpdateMemberOTP(otp OTP) (expiryTime time.Time, err error) {

	//generate otp
	generateotp := func() int {
		const digits = "0123456789"
		var otpp string
		for i := 0; i < otp.Length; i++ {
			randomInt, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
			if err == nil {
				otpp += string(digits[randomInt.Int64()])
			}
		}
		otpint, _ := strconv.Atoi(otpp)
		return otpint
	}

	otp_expiry_time := time.Now().UTC().Add(otp.Duration)

	if err := Authmodel.UpdateMemberOtp(otp.MemberId, generateotp(), otp_expiry_time.Format("2006-01-02 15:04:05"), auth.DB); err != nil {
		return otp_expiry_time, err
	}

	return otp_expiry_time, nil
}
