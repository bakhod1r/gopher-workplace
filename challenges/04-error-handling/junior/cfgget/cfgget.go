// Package cfgget — Gopher Workplace challenge.
package cfgget

import "errors"

// ErrMissingKey reports an absent configuration key.
var ErrMissingKey = errors.New("missing config key")

// Get returns cfg[key], or ErrMissingKey when the key is absent.
//
// Examples:
//
//	Get(map[string]string{"port": "80"}, "port") => "80", nil
//	Get(nil, "port")                             => "", ErrMissingKey
func Get(cfg map[string]string, key string) (string, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
