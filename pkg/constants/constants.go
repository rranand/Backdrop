package constants

import (
	"errors"
	"time"
)

var (
	ErrNotFound                       = errors.New("resource not found")
	ErrConflict                       = errors.New("resource already exists")
	QueryTimeoutDuration              = time.Second * 5
	DatabaseConnectionTimeoutDuration = time.Second * 5
	ServerStopForcefulTimeoutDuration = time.Second * 5
)

type contextKey string

const AuthDataKey = contextKey("authData")
