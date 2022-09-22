package domain

import (
	"errors"
)

var (
	// labstack/echo Variables
	// https://pkg.go.dev/github.com/labstack/echo#pkg-variables
	// Any Internal Server Error occurs
	ErrInternalServerError = errors.New("internal server error")
	// The requested item does not exist
	ErrNotFound = errors.New("your requested item is not found")
	// The current action already exists
	ErrConflict = errors.New("your item already exist")
	// The request-body is not valid
	ErrBadRequestBodyInput = errors.New("given param is not valid")
	// The param is not valid
	ErrBadParamInput      = errors.New("given param is not valid")
	ErrRowsAffectedNotOne = errors.New("the number of affected rows is not 1")
)
