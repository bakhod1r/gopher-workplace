# Handling the default case

## The idea

A switch without a default (or a catch-all return) silently produces the zero value for unlisted inputs.

## Why it matters

Unlabeled fall-through of unexpected inputs causes downstream logic to misbehave on empty strings.

## Watch out

- Always define behaviour for inputs outside the enumerated cases.
- `default:` documents the catch-all intent.
