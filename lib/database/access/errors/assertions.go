package errors

import (
	"errors"
	"fmt"
)

const (
	InterfaceAssertionErrorFormat = "interface conversion error, interface: %s"
)

var (
	UserAssertionError = errors.New(fmt.Sprintf(InterfaceAssertionErrorFormat, "*model.User"))
)
