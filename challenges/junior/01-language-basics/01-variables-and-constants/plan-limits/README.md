# Challenge 01 — Plan Rate Limits

**Level:** Junior  
**Estimated time:** 15 min

## Context

You maintain the billing service. Each subscription tier gets a different
API rate limit. The `Tier` type and the `Limit` function are stubbed out —
you need to implement them from scratch.

## Task

Implement [plan.go](plan.go) so that:

1. `Free`, `Pro`, `Enterprise` are **distinct, ascending** values (`0, 1, 2`).
2. `Limit` returns the correct requests-per-minute for each tier.
3. An unknown tier falls back to the `Free` allowance.

Do **not** change the function signature or the tests.

## Examples

```go
Limit(Free)        // => 60
Limit(Pro)         // => 600    (baseLimit * 10)
Limit(Enterprise)  // => 6000   (baseLimit * 100)
```

## Hints

<details>
<summary>Hint 1</summary>

Two pieces: the `const` block that defines the tiers, and the body of `Limit`.
Start with the constants.
</details>

<details>
<summary>Hint 2</summary>

There is a Go keyword that auto-increments successive constants in a `const`
block, starting from 0 — perfect for `Free, Pro, Enterprise`.
</details>

<details>
<summary>Hint 3</summary>

For `Limit`, a `switch t` (or `if`/`else`) mapping each tier to its multiple of
`baseLimit` works. The `default` case covers unknown tiers by returning
`baseLimit`.
</details>

<details>
<summary>Solution</summary>

```go
const (
	Free Tier = iota // 0
	Pro              // 1
	Enterprise       // 2
)

func Limit(t Tier) int {
	switch t {
	case Pro:
		return baseLimit * 10
	case Enterprise:
		return baseLimit * 100
	default:
		return baseLimit
	}
}
```
</details>

## Topics

<details>
<summary>Topics (5)</summary>

`Variables & Constants` · `iota` · `typed constants` · `variable shadowing` · `zero value`
</details>
