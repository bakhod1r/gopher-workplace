// Package templatemeth — Gopher Workplace challenge.
package templatemeth

// Step defines steps.
type Step interface {
	DoStep1() string
	DoStep2() string
}

// Template runs the steps in order.
type Template struct {
	impl Step
}

// Run executes the template method.
func (t *Template) Run() string {
	// TODO(candidate): return t.impl.DoStep1() + "-" + t.impl.DoStep2()
	panic("not implemented")
}

// MyTask implements Step.
type MyTask struct{}

func (MyTask) DoStep1() string { return "a" }
func (MyTask) DoStep2() string { return "b" }
