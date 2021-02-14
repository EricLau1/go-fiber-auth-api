package util

import "errors"

var (
	ErrInvalidEmail       = errors.New("invalid email")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrEmptyPassword      = errors.New("password can't be empty")
	ErrInvalidAuthToken   = errors.New("invalid auth-token")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnauthorized       = errors.New("Unauthorized")
)
