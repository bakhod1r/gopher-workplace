# Intersection with a lookup set

## The idea

Put one side in a set, then filter the other by membership, de-duplicating the
result:

```go
in := make(map[int]struct{})
for _, x := range a { in[x] = struct{}{} }
seen := make(map[int]struct{})
for _, x := range b {
	if _, ok := in[x]; ok {
		if _, dup := seen[x]; !dup { out = append(out, x); seen[x] = struct{}{} }
	}
}
sort.Ints(out)
```

## Why it matters

Intersection is a core set op (common tags, shared access). The lookup set turns
an O(n·m) nested scan into O(n+m).

## Watch out

- De-dup the output: `b` may repeat a common value.
- Membership is `_, ok := set[x]`.
- Sort for deterministic order.
