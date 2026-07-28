// Package wraperr adds context to a returned error using a deferred closure. A
// planted bug passes err as a defer ARGUMENT, snapshotting its nil value before
// the body sets it, so the wrap never happens.
package wraperr

import (
	"errors"
	"fmt"
)

// Do returns an error wrapped with context. It fails when flag is true.
func Do(flag bool) (err error) {
	// CHANGE CODE BELOW THIS LINE
	defer func(e error) {
		if e != nil {
			err = fmt.Errorf("do: %w", e)
		}
	}(err)
	// CHANGE CODE ABOVE THIS LINE
	if flag {
		err = errors.New("boom")
	}
	return
}
