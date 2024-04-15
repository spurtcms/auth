package auth

import "errors"

var (
	ErrorPassword    = errors.New("invalid password")
	ErrorToken       = errors.New("invalid token")
	ErrorTokenExpiry = errors.New("token expired")
	ErrorConvertTime = errors.New("Could not convert interface to time.Time")
)

type Action string

const (
	Create Action = "Create"

	Read Action = "View"

	Update Action = "Update"

	Delete Action = "Delete"

	CRUD Action = "CRUD"
)
