// Package endpoint — Gopher Workplace challenge.
package endpoint

// The service's address is fixed at build time.
//
// TODO(candidate): give BaseURL and Version their values.
//
//	BaseURL = "https://api.example.com"
//	Version = "v2"
const (
	BaseURL = ""
	Version = ""
)

// Root is the API root: BaseURL, a slash, then Version.
//
// TODO(candidate): declare Root as a package-level var built *from the
// constants above* — do not paste the finished string. Constants are known at
// compile time, so the concatenation costs nothing at run time, and changing
// Version updates Root for free.
//
//	Root == "https://api.example.com/v2"
var Root = ""

// Path returns the full URL of a resource under Root: Root, a slash, then the
// resource name. An empty resource yields Root unchanged (no trailing slash).
//
// Examples:
//
//	Path("users")  => "https://api.example.com/v2/users"
//	Path("orders") => "https://api.example.com/v2/orders"
//	Path("")       => "https://api.example.com/v2"
func Path(resource string) string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
