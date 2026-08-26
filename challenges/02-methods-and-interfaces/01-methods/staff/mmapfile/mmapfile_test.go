package mmapfile

import "testing"

func TestMmap(t *testing.T) {
	m := &Mmap{Data: []byte{10, 20}}
	if got := m.ReadByteAt(1); got != 20 {
		t.Errorf("got %d", got)
	}
}
