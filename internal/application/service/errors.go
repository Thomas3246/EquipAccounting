package service

import "errors"

var (
	ErrUserNotFound           = errors.New("user not found")
	ErrInvalidCredentials     = errors.New("invalid credentials")
	ErrNullParameter          = errors.New("missing required parameter")
	ErrInvalidRole            = errors.New("invalid role")
	ErrInvalidParameter       = errors.New("invalid parameter")
	ErrInvalidCookieParameter = errors.New("invalid cookie parameter")
	ErrNoAccess               = errors.New("no access to page")
)
