// Package classifybug classifies an HTTP status code class. A planted bug omits
// the default, so unknown codes return an empty string instead of "unknown".
package classifybug

// Class returns "success" (2xx), "client" (4xx), "server" (5xx), or "unknown".
func Class(code int) string {
	switch code / 100 {
	case 2:
		return "success"
	case 4:
		return "client"
	case 5:
		return "server"
		// CHANGE CODE BELOW THIS LINE
	}
	return ""
	// CHANGE CODE ABOVE THIS LINE
}
