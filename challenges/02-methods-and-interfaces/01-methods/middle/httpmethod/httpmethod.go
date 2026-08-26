// Package httpmethod — Gopher Workplace challenge.
package httpmethod

import "net/http"

// App holds application state.
type App struct {
	Name string
}

// HealthHandler writes "OK: <Name>" to the response.
func (a *App) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK: " + a.Name))
}

// Handler returns an http.HandlerFunc from the App's HealthHandler method.
// This demonstrates binding a method to an http handler.
//
// Examples:
//
//	h := App{"myapp"}.Handler()
//	// h is an http.HandlerFunc that writes "OK: myapp"
func (a *App) Handler() http.HandlerFunc {
	// TODO(candidate): return a.HealthHandler as http.HandlerFunc.
	panic("not implemented")
}
