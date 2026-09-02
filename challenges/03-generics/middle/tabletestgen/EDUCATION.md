# Generic Table Runner

## Intuition

Because `I` and `W` are fixed per call, passing a table whose expectations are the wrong type fails to compile — something a `[]any` table could never catch.

## Approach

1. Mark the runner as a helper.
2. For each case, run a subtest comparing `fn(c.In)` with `c.Want`.

## Solution

```go
func Run[I, W comparable](t *testing.T, name string, cases []Case[I, W], fn func(I) W) {
	t.Helper()
	for _, c := range cases {
		t.Run(name+"/"+c.Name, func(t *testing.T) {
			t.Helper()
			if got := fn(c.In); got != c.Want {
				t.Errorf("%v: got %v, want %v", c.In, got, c.Want)
			}
		})
	}
}
```

## Walkthrough

A table of three cases produces three named subtests, each failing independently.

## Pitfalls

- Comparing with `reflect.DeepEqual` when `comparable` already gives `==`.
- Omitting `t.Helper()`, so every failure points at the runner.
- Using `t.Fatalf`, which stops the remaining cases.
