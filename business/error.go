package business

import "errors"

var (
	ErrInternalServer = errors.New("something gone wrong, contact administrator")
	ErrNotFound       = errors.New("data not found")
	ErrIDNotFound     = errors.New("id not found")
	ErrDuplicateData  = errors.New("duplicate data")
	ErrUsernameEmpty  = errors.New("username empty")
	ErrEmailEmpty     = errors.New("email empty")
	ErrPasswordEmpty  = errors.New("password empty")
	ErrWrongPassword  = errors.New("password wrong")
	ErrUnathorized    = errors.New("unauthorized")
)
