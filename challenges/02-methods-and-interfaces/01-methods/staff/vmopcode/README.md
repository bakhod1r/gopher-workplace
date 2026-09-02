# Stack VM Dispatch

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A bytecode interpreter is a fetch-decode-execute loop over a flat slice of
ints. The instruction pointer is state on the VM; each `Step` consumes exactly
one instruction — including its operand, when it has one.

## Task

Implement `Step` on `*VM` in [vmopcode.go](vmopcode.go):

1. If `IP >= len(v.Prog)`, return `false`.
2. Fetch the opcode and advance `IP` past it.
3. `OpPush` reads the next word as its operand (advancing `IP` again) and
   pushes it.
4. `OpAdd` and `OpMul` pop two values and push the result; `OpDup` pushes a copy
   of the top.
5. `OpHalt` returns `false`.
6. An opcode with too few operands on the stack does nothing but still advances.
7. Otherwise return `true`.

**Constraint (staff):** a 200,000-instruction program must finish in under a second, allocation-free, leaving exactly one value on the stack.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  PUSH 2, PUSH 3, ADD, HALT
Output: 5
```

**Example 2:**

```
Input:  PUSH 4, DUP, MUL, HALT
Output: 16
```

**Example 3:**

```
Input:  PUSH 1, HALT, PUSH 99
Output: 1   (nothing after HALT runs)
```

_Explanation:_ `Run` stops the moment `Step` reports false.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Variable-width instructions** | `OpPush` occupies two words, so `IP` must advance by two — the classic decode bug. |
| 2 | **Stack discipline** | Pop right operand first; for `Add` and `Mul` the order does not matter, but the habit does. |
| 3 | **Total dispatch** | Underflow and end-of-program are handled, never panicked on. |

## Hint

Advance `IP` right after the fetch, before executing. Then `OpPush`'s operand is
simply `v.Prog[v.IP]` followed by one more increment.

## Validate

```bash
make verify
```
