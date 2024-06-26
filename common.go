package auth

import "errors"

var (
	ErrorPassword       = errors.New("invalid password")
	ErrorToken          = errors.New("invalid token")
	ErrorOtpExpiry      = errors.New("otp expired")
	ErrorConvertTime    = errors.New("could not convert interface to time.Time")
	ErrorMemberLogin    = errors.New("select any one of the config for member login")
	ErrorUnauthorized   = errors.New("Unauthorized")
	ErrorInactive       = errors.New("user disabled please contact admin")
	ErrorInvalidOTP     = errors.New("invalid OTP")
	ErrorInactiveMember = errors.New("inactive member")
	ErrorTokenExpiry    = errors.New("token expired")
)

type Action string

const (
	Create Action = "Create"

	Read Action = "View"

	Update Action = "Update"

	Delete Action = "Delete"

	CRUD Action = "CRUD"
)
