package closerifc

import (
	"errors"
	"testing"
)

func TestClose(t *testing.T) {
	f := &File{}
	if err := f.Close(); err != nil {
		t.Fatalf("first Close = %v, want nil", err)
	}
	if !f.Closed {
		t.Error("Closed = false after Close")
	}
	if err := f.Close(); !errors.Is(err, ErrAlreadyClosed) {
		t.Errorf("second Close = %v, want ErrAlreadyClosed", err)
	}
}

func TestCloseAll(t *testing.T) {
	t.Run("all_ok", func(t *testing.T) {
		if err := CloseAll([]Closer{&File{}, &File{}}); err != nil {
			t.Errorf("CloseAll = %v, want nil", err)
		}
	})

	t.Run("empty", func(t *testing.T) {
		if err := CloseAll(nil); err != nil {
			t.Errorf("CloseAll(nil) = %v, want nil", err)
		}
	})

	t.Run("first_error", func(t *testing.T) {
		bad := &File{Closed: true}
		good := &File{}
		if err := CloseAll([]Closer{bad, good}); !errors.Is(err, ErrAlreadyClosed) {
			t.Errorf("CloseAll = %v, want ErrAlreadyClosed", err)
		}
		if good.Closed {
			t.Error("CloseAll kept going after an error")
		}
	})
}
