// Package commandpatt — Gopher Workplace challenge.
package commandpatt

// Command is a function that mutates state.
type Command func()

// Invoker queues and executes commands.
type Invoker struct {
	commands []Command
}

// Add appends a command.
func (inv *Invoker) Add(c Command) {
	inv.commands = append(inv.commands, c)
}

// ExecuteAll runs all commands and clears the queue.
func (inv *Invoker) ExecuteAll() {
	// TODO(candidate): run all commands, then set inv.commands to nil.
	panic("not implemented")
}
