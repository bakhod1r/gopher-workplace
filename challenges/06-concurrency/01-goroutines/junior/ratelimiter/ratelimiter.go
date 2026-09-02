// Package ratelimiter — Gopher Workplace challenge.
package ratelimiter

// RemainingQuota returns each tenant's remaining request quota, never below zero.
//
// Examples:
//
//	RemainingQuota([]int{10, 90}, 100)  => [90 10]
//	RemainingQuota([]int{150}, 100)     => [0]
//	RemainingQuota(nil, 100)            => []
func RemainingQuota(used []int, limit int) []int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
