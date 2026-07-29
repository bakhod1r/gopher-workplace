# Method Value Captures a Copy

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

`(*a).Get` dereferences to a value first, binding the method value to a COPY of
the account; later changes to `*a` aren't seen. Bind on the pointer (`a.Get`) to
capture the live object.

## Task

Fix [methodcopy.go](methodcopy.go) so the getter reflects later balance changes.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  get := Getter(a); a.Balance = 250; get()
Output: 250
```

_Explanation:_ The closure must observe later mutations.

**Example 2:**

```
Input:  a := &Account{Balance: 100}; Getter(a)()
Output: 100
```

**Example 3:**

```
Input:  get := Getter(a); a.Balance = 0; get()
Output: 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Method value binding** | `(*a).Get` copies the receiver. |
| 2 | **Pointer binding** | `a.Get` keeps the live pointer. |
| 3 | **Live vs snapshot** | Copy freezes the state. |

## Hint

Bind on the pointer: `return a.Get`.

## Validate

```bash
make verify
```
