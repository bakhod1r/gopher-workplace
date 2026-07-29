# Zipping parallel slices

## Intuition

Iterate to the shorter length and pair by index:

```go
n := len(keys)
if len(vals) < n { n = len(vals) }
for i := 0; i < n; i++ { m[keys[i]] = vals[i] }
```

## Approach

1. n = min(len(keys),len(vals)).
2. For i in [0,n) set out[keys[i]]=vals[i].
3. Extra elements of the longer slice are ignored.
4. Return map.

## Solution

```go
func Zip(keys []string, vals []int) map[string]int {
	out := map[string]int{}
	n := len(keys)
	if len(vals) < n {
		n = len(vals)
	}
	for i := 0; i < n; i++ {
		out[keys[i]] = vals[i]
	}
	return out
}
```

## Walkthrough

keys=["a","b"] vals=[1,2,3]: n=2 -> a:1,b:2; vals[2] ignored.

## Pitfalls

- Mismatched lengths: iterate `min`, don't assume equal.
- Duplicate keys: later pairs overwrite earlier ones.
- `min` is a builtin in Go 1.21+.
