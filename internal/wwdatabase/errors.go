package client

import (
	"errors"
)

var (
	ErrInternal            = errors.New("service internal error")
	ErrUnavailable         = errors.New("service unavailable")
	ErrUnknown             = errors.New("service error")
	ErrNoNameOrID          = errors.New("must provide name or ID")
	ErrDateIncorrect       = errors.New("date incorrectly formatted")
	ErrNoChampionsReturned = errors.New("no champions returned from query")
	ErrTableDoesNotExist   = errors.New("error in query, the table you are trying to access does not exist")
)
