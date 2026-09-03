# Attach The Call Site

## Intuition

Go errors carry no stack by default, which keeps them cheap. When a location genuinely helps, capturing one frame at the point of creation is enough and costs almost nothing.

## Approach

1. Return nil for a nil error.
2. Call `runtime.Caller(1)` for the caller's frame.
3. Format the base file name, line and wrapped error.

## Solution

```go
if err == nil {
	return nil
}
_, file, line, ok := runtime.Caller(1)
if !ok {
	return err
}
return fmt.Errorf("%s:%d: %w", filepath.Base(file), line, err)
```

## Walkthrough

Two calls on different lines produce different prefixes, which is what makes the annotation useful for locating the failure.

## Pitfalls

- Using skip 0, which always reports this file.
- Emitting the full absolute path, which differs per build machine.
- Replacing the error rather than wrapping it.
