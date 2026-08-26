package templatemeth

import "testing"

func TestTemplate(t *testing.T) {
	tmpl := &Template{impl: MyTask{}}
	if got := tmpl.Run(); got != "a-b" {
		t.Errorf("Run() = %q, want a-b", got)
	}
}
