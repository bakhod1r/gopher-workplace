# `const` vs `var`, and package-level state

## The idea

A `const` is a value the compiler knows. A `var` is storage the program owns.

```go
const Version = "v2"                  // fixed at compile time
var Root = BaseURL + "/" + Version    // a variable, initialised before main runs
```

Constants may only hold values the compiler can compute: numbers, strings,
booleans, runes, and expressions over them. Anything that needs the machine at
run time — a function call, a slice, a map, a struct, `time.Now()` — must be a
`var`.

## Package-level variables initialise before `main`

Go initialises package-level variables in **dependency order**, not source
order, then runs `init()` functions, then `main`:

```go
var Root = BaseURL + "/" + Version   // depends on two constants
var host = parse(Root)               // depends on Root — Go runs it second
```

You never have to order the declarations by hand; the compiler works out what
each one needs. A cycle is a compile error.

## Derive, do not duplicate

The whole point of the exercise:

```go
var Root = "https://api.example.com/v2"   // pasted: goes stale silently
var Root = BaseURL + "/" + Version        // derived: one place to change
```

Both produce the same string today. Bump `Version` to `"v3"` and only the second
one follows. Because both operands are constants, the compiler folds the
concatenation — the derived form costs nothing at run time and documents where
the value comes from.

## Constant expressions are free

```go
const KB = 1 << 10
const timeout = 30 * 1000       // computed at compile time
var greeting = "hello, " + name // name is a var → run-time concatenation
```

The first two exist only in the compiled instructions. The third allocates.

## Watch out

- Package-level `var`s are global mutable state. Prefer `const` when the value
  never changes, and prefer passing values as parameters over reaching for a
  package-level variable.
- A `var` cannot be used where a constant is required — array lengths, `case`
  values in a constant switch, other constant declarations.
- Initialisation order across *packages* follows imports: imported packages are
  fully initialised first.
- `init()` runs after variable initialisation, once per package, and cannot be
  called by hand. Use it sparingly; explicit setup is easier to follow and test.

## Try it yourself

```go
const (
	BaseURL = "https://api.example.com"
	Version = "v2"
)

var Root = BaseURL + "/" + Version

func Path(resource string) string {
	if resource == "" {
		return Root
	}
	return Root + "/" + resource
}

fmt.Println(Root)            // https://api.example.com/v2
fmt.Println(Path("users"))   // https://api.example.com/v2/users
fmt.Println(Path(""))        // https://api.example.com/v2 — no trailing slash
```
