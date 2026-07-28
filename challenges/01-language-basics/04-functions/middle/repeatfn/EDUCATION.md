# Iterated application

## The idea

A closure over a function and a count applies it repeatedly; zero applications is the identity.

## Why it matters

Fixed-point iteration and repeated transforms build on this.

## Watch out

- n==0 must return the argument unchanged.
- The closure captures f and n by reference; don't mutate them after.
