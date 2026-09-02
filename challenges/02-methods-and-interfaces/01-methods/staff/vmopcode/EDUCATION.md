# Stack VM Dispatch

## Intuition

A bytecode VM is a loop with a cursor. Everything difficult about it comes from
one fact: instructions are not all the same width. `IP` must always land on the
next *opcode*, never on an operand, or the interpreter starts executing data.

Keeping `IP` and `Stack` on the receiver makes each `Step` a pure state
transition, which is what lets `Run` be a three-line loop and lets a debugger
single-step the machine.

## Approach

1. Bounds-check `IP`.
2. Fetch and immediately advance past the opcode.
3. Switch; `OpPush` consumes one more word.
4. Guard every pop against underflow.
5. Report whether to continue.

## Solution

```go
func (v *VM) Step() bool {
	if v.IP >= len(v.Prog) {
		return false
	}

	op := v.Prog[v.IP]
	v.IP++

	switch op {
	case OpPush:
		if v.IP < len(v.Prog) {
			v.Stack = append(v.Stack, v.Prog[v.IP])
			v.IP++
		}
	case OpAdd, OpMul:
		if len(v.Stack) >= 2 {
			b := v.Stack[len(v.Stack)-1]
			a := v.Stack[len(v.Stack)-2]
			v.Stack = v.Stack[:len(v.Stack)-2]
			if op == OpAdd {
				v.Stack = append(v.Stack, a+b)
			} else {
				v.Stack = append(v.Stack, a*b)
			}
		}
	case OpDup:
		if len(v.Stack) >= 1 {
			v.Stack = append(v.Stack, v.Stack[len(v.Stack)-1])
		}
	case OpHalt:
		return false
	}

	return true
}
```

## Walkthrough

`PUSH 2, PUSH 3, ADD, HALT`:

| IP before | opcode | stack after | IP after |
|-----------|--------|-------------|----------|
| 0 | PUSH 2 | [2] | 2 |
| 2 | PUSH 3 | [2 3] | 4 |
| 4 | ADD | [5] | 5 |
| 5 | HALT | [5] | 6, loop ends |

Note that after `PUSH 2` the pointer is at 2, not 1 — that second increment is
the whole decode contract. Get it wrong and step 2 fetches the literal `2` as an
opcode, which happens to be `OpMul`, and the program quietly computes garbage.

`{OpAdd, OpPush, 3, OpHalt}` starts with an underflowing `ADD`: the guard sees
fewer than two values, does nothing, and `IP` still moves on, so the `PUSH 3`
executes and `Run` returns 3.

## Pitfalls

- **Forgetting the operand increment.** The most common bytecode bug there is,
  and it fails silently because operands are valid opcodes.
- **Popping in the wrong order.** Harmless for `+` and `*`; catastrophic the
  moment you add `OpSub` or `OpDiv`.
- **Panicking on underflow.** A real VM raises a trap; the contract here is a
  no-op, which keeps `Step` total.
- **Advancing `IP` after the switch.** Then `OpPush` has to compensate, and the
  two increments drift apart as opcodes are added.

## Why `Step` and `Run` are separate

`Run` is a loop over a state machine, and keeping the transition in its own
method is what makes the VM inspectable: tests can single-step, a debugger can
break between instructions, and tracing is a wrapper rather than an edit to the
interpreter core.
