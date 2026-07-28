# Pointer receivers mutate

## The idea

A method with a **pointer receiver** operates on the original struct:

```go
func (a *Account) Deposit(amount int) { a.Balance += amount }
```

A value receiver would receive a copy and its changes would vanish.

## Why it matters

Any method that updates state needs a pointer receiver. Mixing value and pointer
receivers on one type is usually a smell.

## Watch out

- Pointer-receiver calls need an **addressable** value; a map element is not.
- Value receiver + field write silently mutates a copy.
- Large structs are cheaper to pass by pointer regardless of mutation.
