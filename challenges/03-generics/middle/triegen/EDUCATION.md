# Prefix Tree

## Intuition

The prefix question and the exact-key question are different: `HasPrefix` only needs the path to exist, while `Get` needs the terminal node to be marked.

## Approach

1. `Insert`: create missing children while walking, then store the value and set the flag.
2. `Get`: walk, failing on a missing child, then check the flag.
3. `HasPrefix`: walk and report whether every rune had a child.

## Solution

```go
func NewTrie[V any]() *Trie[V] {
	return &Trie[V]{children: make(map[rune]*Trie[V])}
}

func (t *Trie[V]) Insert(key string, v V) {
	node := t
	for _, r := range key {
		next, ok := node.children[r]
		if !ok {
			next = &Trie[V]{children: make(map[rune]*Trie[V])}
			node.children[r] = next
		}
		node = next
	}
	node.value = v
	node.set = true
}

func (t *Trie[V]) Get(key string) (V, bool) {
	node := t
	for _, r := range key {
		next, ok := node.children[r]
		if !ok {
			var zero V
			return zero, false
		}
		node = next
	}
	if !node.set {
		var zero V
		return zero, false
	}
	return node.value, true
}

func (t *Trie[V]) HasPrefix(prefix string) bool {
	node := t
	for _, r := range prefix {
		next, ok := node.children[r]
		if !ok {
			return false
		}
		node = next
	}
	return true
}
```

## Walkthrough

After `Insert("go", 1)`, the node for `"g"` exists but is unset, so `Get("g")` reports `false` while `HasPrefix("g")` reports `true`.

## Pitfalls

- Treating any reachable node as a stored key.
- Comparing the stored value against its zero value instead of using the flag.
- Indexing the key by byte, which breaks for multi-byte runes.
