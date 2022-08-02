package hferrors

import (
	"errors"
)

var (
	ErrNoAPIKeyFound = errors.New("error: please configure an API key")
)
