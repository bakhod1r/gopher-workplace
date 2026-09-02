# HyperLogLog Register

## Intuition

Hash an item and count the leading zeros. Getting *k* zeros has probability
2⁻ᵏ, so the longest run you have ever seen is evidence about how many distinct
hashes went past — roughly 2^maxZeros of them. Duplicates hash identically and
cannot push the register higher, which is why the count is of *distinct* items
and why memory stays constant.

## Approach

1. Compute the run length for this hash.
2. Keep it only if it beats the record.

## Solution

```go
func (h *HLL) Add(hash uint32) {
	if zeros := leadingZeros(hash); zeros > h.maxZeros {
		h.maxZeros = zeros
	}
}
```

## Walkthrough

The mock `leadingZeros` returns `v % 5`. `Add(1)` yields 1, which beats the
zero value: register 1. `Add(4)` yields 4: register 4. `Add(2)` yields 2, which
does not beat 4, so the register holds. Re-adding any of them changes nothing —
the operation is idempotent, exactly as the real algorithm requires.

## Pitfalls

- **Assigning unconditionally.** The register would then track the *last* item,
  making the estimate depend on arrival order and collapse on a duplicate.
- **`>=` instead of `>`.** Harmless here, but it signals a misunderstanding: the
  register is a maximum, not a counter.
- **Value receiver.** The maximum never persists.

## What the real structure adds

Real HyperLogLog splits the hash: the first *p* bits pick one of 2^p registers,
the rest supply the zero run. Averaging across registers (harmonically, with a
bias correction) is what turns a wildly noisy single estimate into roughly 2%
error using a few kilobytes. Merging two sketches is then just a per-register
max — the same operation implemented here, which is why it must be monotone.
