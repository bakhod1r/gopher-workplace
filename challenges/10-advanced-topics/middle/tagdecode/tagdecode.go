// Package tagdecode — Gopher Workplace challenge.
package tagdecode

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// ErrTarget is returned when dst is not a non-nil pointer to a struct.
var ErrTarget = errors.New("dst must be a non-nil pointer to a struct")

// Decode fills dst's fields from src, matching by the field's `env` tag.
//
// Supported field kinds are string, int and bool. Fields without an env
// tag, unexported fields, and keys missing from src are left alone.
//
// Examples:
//
//	Decode(map[string]string{"PORT": "80"}, &cfg) => cfg.Port == 80
func Decode(src map[string]string, dst any) error {
	panic("not implemented")
}
