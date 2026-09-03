"""10-advanced-topics / 03-reflection — rotation 1: 5 puzzles per level."""

SUB = "03-reflection"

SPECS = []


def P(level, **kw):
    kw.setdefault("sub", SUB)
    kw["level"] = level
    kw.setdefault("mode", "stub")
    kw.setdefault("stub", 'panic("not implemented")')
    kw.setdefault("imports", [])
    kw.setdefault("extra", "")
    SPECS.append(kw)


# ---------------------------------------------------------------- junior -----

P(
    "junior",
    name="kindof",
    title="Ask A Value What It Is",
    sig="func KindName(v any) string",
    doc="""KindName returns the name of v's underlying kind: "int", "slice",
"struct" and so on.

A nil interface has no type at all, so it reports "invalid".

Examples:

	KindName(3) => "int" """,
    imports=['"reflect"'],
    solution="""return reflect.ValueOf(v).Kind().String()""",
    tests="""
import "testing"

type point struct{ X int }

func TestKindName(t *testing.T) {
	cases := []struct {
		in   any
		want string
	}{
		{3, "int"},
		{"s", "string"},
		{3.5, "float64"},
		{[]int{1}, "slice"},
		{map[string]int{}, "map"},
		{point{}, "struct"},
		{&point{}, "ptr"},
		{true, "bool"},
		{nil, "invalid"},
	}
	for _, c := range cases {
		if got := KindName(c.in); got != c.want {
			t.Errorf("KindName(%#v) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestKindNameIgnoresTheNamedType(t *testing.T) {
	type myInt int
	if got := KindName(myInt(1)); got != "int" {
		t.Errorf("KindName = %q, want \\"int\\": a kind is not a type name", got)
	}
}
""",
    context="A debug endpoint prints the shape of whatever it is handed. `%T` gives the declared type name, but the handler needs to branch on the underlying shape instead.",
    task=[
        "Return the name of `v`'s kind.",
        "A nil interface reports `\"invalid\"`.",
        "A named type reports its underlying kind, not its own name.",
    ],
    examples=[
        ("KindName(3)", '"int"', None),
        ("KindName([]int{1})", '"slice"', None),
        ("KindName(nil)", '"invalid"', "There is no type inside a nil interface."),
    ],
    topics=[
        ("reflect.ValueOf", "Opens an interface value for inspection at run time."),
        ("Kind vs Type", "`Kind` is the underlying shape; `Type` is the declared name."),
        ("The zero Value", "`reflect.ValueOf(nil)` is invalid, and `Kind()` says so instead of panicking."),
    ],
    hint="One expression: value, kind, string.",
    intuition="Reflection starts by opening the interface: `ValueOf` gives you a handle you can ask questions of. `Kind` is the coarsest question — what shape is this, regardless of what it is called.",
    approach=[
        "`reflect.ValueOf(v)`.",
        "`.Kind()` for the shape, `.String()` for its name.",
    ],
    walkthrough="`myInt(1)` has type name `myInt` and kind `int`. `KindName` reports the kind, so both `1` and `myInt(1)` give \"int\".",
    pitfalls=[
        "`reflect.TypeOf(nil)` returns nil and panics when you call `Kind` on it; `ValueOf` gives a usable zero Value instead.",
        "Confusing `Kind().String()` with `Type().String()` — the latter gives \"main.myInt\".",
    ],
)

P(
    "junior",
    name="fieldnames",
    title="List A Struct's Exported Fields",
    sig="func FieldNames(v any) []string",
    doc="""FieldNames returns the names of v's exported fields, in declaration
order.

A non-struct, or a nil interface, yields nil.

Examples:

	FieldNames(struct{ A, b int }{}) => []string{"A"}""",
    imports=['"reflect"'],
    solution="""t := reflect.TypeOf(v)
if t == nil || t.Kind() != reflect.Struct {
	return nil
}
var out []string
for i := 0; i < t.NumField(); i++ {
	if f := t.Field(i); f.IsExported() {
		out = append(out, f.Name)
	}
}
return out""",
    tests="""
import (
	"reflect"
	"testing"
)

type user struct {
	Name  string
	Age   int
	admin bool
}

func TestFieldNames(t *testing.T) {
	if got := FieldNames(user{}); !reflect.DeepEqual(got, []string{"Name", "Age"}) {
		t.Errorf("FieldNames = %v, want [Name Age]", got)
	}
	if got := FieldNames(struct{}{}); got != nil {
		t.Errorf("FieldNames = %v, want nil for an empty struct", got)
	}
}

func TestFieldNamesRejectsNonStructs(t *testing.T) {
	for _, in := range []any{3, "s", []int{1}, map[string]int{}, nil, &user{}} {
		if got := FieldNames(in); got != nil {
			t.Errorf("FieldNames(%#v) = %v, want nil", in, got)
		}
	}
}

func TestFieldNamesIsInDeclarationOrder(t *testing.T) {
	type ordered struct {
		Z, A, M int
	}
	if got := FieldNames(ordered{}); !reflect.DeepEqual(got, []string{"Z", "A", "M"}) {
		t.Errorf("FieldNames = %v, want [Z A M]", got)
	}
}
""",
    context="A config auditor wants to report which settings a struct exposes. Hand-maintaining the list drifts from the struct within a release.",
    task=[
        "Return the exported field names of `v`, in declaration order.",
        "Skip unexported fields.",
        "Return nil for anything that is not a struct, including nil and a pointer to a struct.",
    ],
    examples=[
        ("FieldNames(user{})", "[Name Age]", "`admin` is unexported."),
        ("FieldNames(&user{})", "<nil>", "A pointer is not a struct."),
        ("FieldNames(3)", "<nil>", None),
    ],
    topics=[
        ("reflect.Type.NumField / Field", "Struct layout is walked by index, in declaration order."),
        ("StructField.IsExported", "Export status is part of the field's metadata."),
        ("Guarding the kind", "`NumField` panics on a non-struct, so check first."),
    ],
    hint="`reflect.TypeOf(nil)` is nil — check that before you check the kind.",
    intuition="A struct type carries its full field list at run time: names, types, tags and export status. Reflection just reads that table.",
    approach=[
        "Take `reflect.TypeOf(v)`; return nil if it is nil or not a struct.",
        "Loop over `NumField()`, appending each exported field's `Name`.",
    ],
    walkthrough="`user` has three fields; `admin` starts with a lower-case letter, so `IsExported` is false and it is skipped, leaving [Name Age].",
    pitfalls=[
        "Calling `NumField` before checking the kind — that is a panic, not an error.",
        "Dereferencing pointers; the spec here says a pointer yields nil.",
    ],
)

P(
    "junior",
    name="taglookup",
    title="Read A Struct Tag",
    sig="func Tag(v any, field, key string) (string, bool)",
    doc="""Tag returns the value of the given key in the named field's struct tag.

The second result reports whether the field exists and carries that key.

Examples:

	Tag(row{}, "ID", "json") => "id", true""",
    imports=['"reflect"'],
    solution="""t := reflect.TypeOf(v)
if t == nil || t.Kind() != reflect.Struct {
	return "", false
}
f, ok := t.FieldByName(field)
if !ok {
	return "", false
}
return f.Tag.Lookup(key)""",
    tests="""
import "testing"

type row struct {
	ID   int    `json:"id" db:"row_id"`
	Name string `json:"name"`
	Skip string
}

func TestTag(t *testing.T) {
	cases := []struct {
		field, key, want string
		ok               bool
	}{
		{"ID", "json", "id", true},
		{"ID", "db", "row_id", true},
		{"Name", "json", "name", true},
		{"Name", "db", "", false},
		{"Skip", "json", "", false},
		{"Missing", "json", "", false},
	}
	for _, c := range cases {
		got, ok := Tag(row{}, c.field, c.key)
		if got != c.want || ok != c.ok {
			t.Errorf("Tag(%q, %q) = %q, %v, want %q, %v", c.field, c.key, got, ok, c.want, c.ok)
		}
	}
}

func TestTagRejectsNonStructs(t *testing.T) {
	for _, in := range []any{nil, 3, []int{1}} {
		if _, ok := Tag(in, "ID", "json"); ok {
			t.Errorf("Tag(%#v) reported ok, want false", in)
		}
	}
}
""",
    context="A serialiser needs to honour the same `json` tags the standard library reads, plus a `db` tag of its own. Both live on the field, and both have to be read at run time.",
    task=[
        "Return the value of `key` in `field`'s struct tag, and whether it was present.",
        "A missing field, a missing key, or a non-struct all report false.",
    ],
    examples=[
        ('Tag(row{}, "ID", "json")', '"id", true', None),
        ('Tag(row{}, "Name", "db")', '"", false', "The field has no db tag."),
        ('Tag(row{}, "Missing", "json")', '"", false', None),
    ],
    topics=[
        ("StructTag.Lookup", "Returns the value and whether the key was present — unlike `Get`, which cannot tell empty from absent."),
        ("FieldByName", "Finds a field by name and reports whether it exists."),
        ("Tags are strings", "The conventional `key:\"value\"` format is parsed by the `StructTag` methods."),
    ],
    hint="`Get` returns \"\" for both an empty tag and a missing one. You need the other method.",
    intuition="Tags are just a string attached to each field. The `StructTag` type parses the conventional format for you, and `Lookup` is the version that distinguishes \"absent\" from \"present but empty\".",
    approach=[
        "Guard against a nil or non-struct type.",
        "`FieldByName(field)`; return false if it is missing.",
        "Return `f.Tag.Lookup(key)` directly — its results already match the signature.",
    ],
    walkthrough='`row.ID` carries `json:"id" db:"row_id"`. Looking up "db" parses the tag and returns "row_id", true. Looking up "db" on `Name` returns "", false.',
    pitfalls=[
        "Using `Tag.Get`, which cannot report absence.",
        "Forgetting that `FieldByName` also finds promoted fields from embedded structs.",
    ],
)

P(
    "junior",
    name="zerocheck",
    title="Is This Value The Zero Value",
    sig="func IsZero(v any) bool",
    doc="""IsZero reports whether v holds the zero value for its type.

A nil interface counts as zero: there is nothing in it.

Examples:

	IsZero(0) => true""",
    imports=['"reflect"'],
    solution="""rv := reflect.ValueOf(v)
if !rv.IsValid() {
	return true
}
return rv.IsZero()""",
    tests="""
import "testing"

type pair struct {
	A int
	B string
}

func TestIsZero(t *testing.T) {
	cases := []struct {
		in   any
		want bool
	}{
		{nil, true},
		{0, true},
		{1, false},
		{"", true},
		{"x", false},
		{false, true},
		{true, false},
		{pair{}, true},
		{pair{A: 1}, false},
		{(*pair)(nil), true},
		{&pair{}, false},
		{[]int(nil), true},
		{[]int{}, false},
		{0.0, true},
	}
	for _, c := range cases {
		if got := IsZero(c.in); got != c.want {
			t.Errorf("IsZero(%#v) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestIsZeroArrays(t *testing.T) {
	if !IsZero([2]int{}) {
		t.Error("IsZero([2]int{}) = false, want true")
	}
	if IsZero([2]int{0, 1}) {
		t.Error("IsZero([2]int{0,1}) = true, want false")
	}
}
""",
    context="A patch endpoint must tell \"field omitted\" from \"field set to its zero value\". The first cut compared against `nil` and treated every explicit zero as absent.",
    task=[
        "Report whether `v` holds the zero value for its type.",
        "A nil interface is zero.",
        "A nil pointer is zero; a pointer to a zero struct is not.",
    ],
    examples=[
        ("IsZero(0)", "true", None),
        ("IsZero(&pair{})", "false", "The pointer itself is not nil."),
        ("IsZero([]int{})", "false", "An empty non-nil slice is not the zero slice."),
    ],
    topics=[
        ("Value.IsZero", "Compares against the type's zero value, whatever the type is."),
        ("Value.IsValid", "False exactly when the interface held nothing."),
        ("nil vs empty", "A nil slice is the zero value; an allocated empty one is not."),
    ],
    hint="Calling a method on an invalid Value panics. Check validity first.",
    intuition="\"Zero\" is a per-type notion — 0, \"\", nil, an all-zero struct. Reflection knows the type, so it can answer the question generically, but only once you have handled the case where there is no type at all.",
    approach=[
        "`reflect.ValueOf(v)`.",
        "Return true when the Value is invalid — a nil interface holds nothing.",
        "Otherwise return `rv.IsZero()`.",
    ],
    walkthrough="`IsZero((*pair)(nil))` sees a valid Value of kind ptr whose value is nil, so `IsZero` is true. `IsZero(&pair{})` sees a non-nil pointer, so it is false even though what it points at is zero.",
    pitfalls=[
        "`v == nil` on the interface misses a typed nil pointer, which is not `nil` as an interface.",
        "Calling `IsZero` on the invalid Value from `reflect.ValueOf(nil)` panics.",
    ],
)

P(
    "junior",
    name="setfield",
    title="Write A Field Through A Pointer",
    sig="func SetInt(ptr any, field string, n int) error",
    doc="""SetInt sets the named int field of the struct ptr points at.

Reflection can only write through a pointer: a value passed by interface
is a copy, and the reflect package refuses to modify it.

Examples:

	SetInt(&counters{}, "Hits", 3) => nil, Hits is 3""",
    imports=['"errors"', '"reflect"'],
    extra="""// ErrNotSettable is returned when the target cannot be written.
var ErrNotSettable = errors.New("target is not a settable int field")""",
    solution="""rv := reflect.ValueOf(ptr)
if rv.Kind() != reflect.Pointer || rv.IsNil() {
	return ErrNotSettable
}
rv = rv.Elem()
if rv.Kind() != reflect.Struct {
	return ErrNotSettable
}
f := rv.FieldByName(field)
if !f.IsValid() || !f.CanSet() || f.Kind() != reflect.Int {
	return ErrNotSettable
}
f.SetInt(int64(n))
return nil""",
    tests="""
import (
	"errors"
	"testing"
)

type counters struct {
	Hits   int
	Name   string
	hidden int
}

func TestSetInt(t *testing.T) {
	c := &counters{}
	if err := SetInt(c, "Hits", 42); err != nil {
		t.Fatalf("SetInt = %v, want nil", err)
	}
	if c.Hits != 42 {
		t.Errorf("Hits = %d, want 42", c.Hits)
	}
}

func TestSetIntRejectsBadTargets(t *testing.T) {
	cases := []struct {
		name string
		ptr  any
		f    string
	}{
		{"value not pointer", counters{}, "Hits"},
		{"nil pointer", (*counters)(nil), "Hits"},
		{"nil interface", nil, "Hits"},
		{"missing field", &counters{}, "Nope"},
		{"wrong kind", &counters{}, "Name"},
		{"unexported", &counters{}, "hidden"},
		{"not a struct", new(int), "Hits"},
	}
	for _, c := range cases {
		if err := SetInt(c.ptr, c.f, 1); !errors.Is(err, ErrNotSettable) {
			t.Errorf("%s: err = %v, want ErrNotSettable", c.name, err)
		}
	}
}

func TestSetIntOverwrites(t *testing.T) {
	c := &counters{Hits: 5}
	if err := SetInt(c, "Hits", 0); err != nil || c.Hits != 0 {
		t.Errorf("Hits = %d, err = %v, want 0, nil", c.Hits, err)
	}
}
""",
    context="A test fixture builder wants to poke one field of a struct by name. The first attempt passed the struct by value and every write silently went nowhere — until reflect started panicking instead.",
    task=[
        "Set the named int field of the struct `ptr` points at.",
        "Return `ErrNotSettable` for a non-pointer, a nil pointer, a non-struct, a missing field, an unexported field, or a field that is not an int.",
    ],
    examples=[
        ('SetInt(&counters{}, "Hits", 42)', "nil, Hits is 42", None),
        ('SetInt(counters{}, "Hits", 1)', "ErrNotSettable", "A value, not a pointer."),
        ('SetInt(&counters{}, "hidden", 1)', "ErrNotSettable", "Unexported fields are not settable."),
    ],
    topics=[
        ("Addressability", "Only a Value obtained through a pointer's `Elem` can be set."),
        ("CanSet", "The check that covers both addressability and export status."),
        ("Kind before Set", "`SetInt` panics on a non-int field, so verify the kind first."),
    ],
    hint="`reflect.ValueOf(x)` is never settable. What makes a Value addressable?",
    intuition="Reflection follows Go's own rules: you cannot assign to a copy. `ValueOf` gives you a copy of whatever was in the interface, so writing requires starting from a pointer and stepping through `Elem` to the real storage.",
    approach=[
        "Verify `ptr` is a non-nil pointer; take `Elem()`.",
        "Verify it is a struct and find the field by name.",
        "Verify the field is valid, settable and of kind int, then `SetInt`.",
    ],
    walkthrough="`&counters{}` gives a ptr Value; `Elem()` is the addressable struct; `FieldByName(\"Hits\")` is an addressable exported int, so `CanSet` is true and the write lands in the caller's struct.",
    pitfalls=[
        "Skipping `CanSet` — an unexported field is valid and of the right kind, and setting it panics.",
        "`SetInt` on a string field panics; the kind check is not optional.",
    ],
)

# ---------------------------------------------------------------- middle -----

P(
    "middle",
    name="tagdecode",
    title="Fill A Struct From A String Map",
    sig="func Decode(src map[string]string, dst any) error",
    doc="""Decode fills dst's fields from src, matching by the field's `env` tag.

Supported field kinds are string, int and bool. Fields without an env
tag, unexported fields, and keys missing from src are left alone.

Examples:

	Decode(map[string]string{"PORT": "80"}, &cfg) => cfg.Port == 80""",
    imports=['"errors"', '"fmt"', '"reflect"', '"strconv"'],
    extra="""// ErrTarget is returned when dst is not a non-nil pointer to a struct.
var ErrTarget = errors.New("dst must be a non-nil pointer to a struct")""",
    solution="""rv := reflect.ValueOf(dst)
if rv.Kind() != reflect.Pointer || rv.IsNil() {
	return ErrTarget
}
rv = rv.Elem()
if rv.Kind() != reflect.Struct {
	return ErrTarget
}
rt := rv.Type()
for i := 0; i < rt.NumField(); i++ {
	f := rt.Field(i)
	key, ok := f.Tag.Lookup("env")
	if !ok || key == "" || key == "-" || !f.IsExported() {
		continue
	}
	s, ok := src[key]
	if !ok {
		continue
	}
	fv := rv.Field(i)
	switch fv.Kind() {
	case reflect.String:
		fv.SetString(s)
	case reflect.Int:
		n, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("field %s: %w", f.Name, err)
		}
		fv.SetInt(int64(n))
	case reflect.Bool:
		b, err := strconv.ParseBool(s)
		if err != nil {
			return fmt.Errorf("field %s: %w", f.Name, err)
		}
		fv.SetBool(b)
	default:
		return fmt.Errorf("field %s: unsupported kind %s", f.Name, fv.Kind())
	}
}
return nil""",
    tests="""
import (
	"errors"
	"testing"
)

type config struct {
	Host    string `env:"HOST"`
	Port    int    `env:"PORT"`
	Debug   bool   `env:"DEBUG"`
	Ignored string
	skipped string `env:"SKIPPED"`
	Dashed  string `env:"-"`
}

func TestDecode(t *testing.T) {
	var c config
	src := map[string]string{
		"HOST": "localhost", "PORT": "8080", "DEBUG": "true",
		"SKIPPED": "x", "-": "y", "UNKNOWN": "z",
	}
	if err := Decode(src, &c); err != nil {
		t.Fatalf("Decode = %v, want nil", err)
	}
	if c.Host != "localhost" || c.Port != 8080 || !c.Debug {
		t.Errorf("c = %+v, want {localhost 8080 true ...}", c)
	}
	if c.Ignored != "" || c.skipped != "" || c.Dashed != "" {
		t.Errorf("c = %+v: untagged, unexported and \\"-\\" fields must be left alone", c)
	}
}

func TestDecodeLeavesMissingKeys(t *testing.T) {
	c := config{Host: "keep", Port: 1}
	if err := Decode(map[string]string{"DEBUG": "1"}, &c); err != nil {
		t.Fatal(err)
	}
	if c.Host != "keep" || c.Port != 1 || !c.Debug {
		t.Errorf("c = %+v, want the untouched fields preserved", c)
	}
}

func TestDecodeBadValues(t *testing.T) {
	var c config
	if err := Decode(map[string]string{"PORT": "eighty"}, &c); err == nil {
		t.Error("want an error for a non-numeric port, got nil")
	}
	if err := Decode(map[string]string{"DEBUG": "maybe"}, &c); err == nil {
		t.Error("want an error for a non-boolean debug, got nil")
	}
}

func TestDecodeBadTarget(t *testing.T) {
	for _, dst := range []any{config{}, nil, (*config)(nil), new(int)} {
		if err := Decode(nil, dst); !errors.Is(err, ErrTarget) {
			t.Errorf("Decode(%#v) = %v, want ErrTarget", dst, err)
		}
	}
}

func TestDecodeUnsupportedKind(t *testing.T) {
	var bad struct {
		F float64 `env:"F"`
	}
	if err := Decode(map[string]string{"F": "1.5"}, &bad); err == nil {
		t.Error("want an error for an unsupported kind, got nil")
	}
}
""",
    context="Every service in the fleet parses its own environment by hand. The parsing is identical, the mistakes are not, and each one is found in production.",
    task=[
        "Fill `dst`'s fields from `src`, matching each field's `env` tag against the map key.",
        "Support string, int and bool fields; report an error for any other tagged kind.",
        "Skip unexported fields, untagged fields, `env:\"-\"`, and keys absent from `src`.",
        "Return `ErrTarget` unless `dst` is a non-nil pointer to a struct.",
    ],
    examples=[
        ('Decode(map[string]string{"PORT":"8080"}, &cfg)', "nil, cfg.Port is 8080", None),
        ('Decode(map[string]string{"PORT":"eighty"}, &cfg)', "a parse error", None),
        ("Decode(src, config{})", "ErrTarget", "A value cannot be written through."),
    ],
    topics=[
        ("Tag-driven mapping", "The tag is the contract between the external name and the field."),
        ("Kind switching", "Each supported kind needs its own parse and its own `Set` method."),
        ("Settability", "The pointer plus `Elem` is what makes the fields writable."),
        ("Error wrapping", "`%w` keeps the underlying `strconv` error inspectable."),
    ],
    hint="One loop over the fields. Per field: tag, lookup, kind switch, set.",
    intuition="Decoding is a table walk. The struct type is the table — each row has a name, a kind and a tag — and the map is the data. Reflection lets you write the walk once instead of once per config struct.",
    approach=[
        "Validate `dst` and step to the struct with `Elem`.",
        "For each field, read the `env` tag and skip the exclusions.",
        "Look the key up in `src`; skip if absent.",
        "Switch on the field's kind, parse, and set. Wrap parse failures with the field name.",
    ],
    walkthrough='With `PORT` set to "8080", the loop reaches `Port`, reads tag "PORT", finds the value, sees kind int, parses 8080 and calls `SetInt`. `Ignored` has no tag and is never touched.',
    pitfalls=[
        "Checking `IsExported` after calling `Set` — the panic comes first.",
        "Treating `env:\"-\"` as a key named \"-\".",
        "Returning early on the first missing key instead of skipping it.",
    ],
)

P(
    "middle",
    name="mapkeys",
    title="Sorted Keys Of Any String-Keyed Map",
    sig="func Keys(m any) []string",
    doc="""Keys returns the keys of m sorted in ascending order.

m must be a map with string keys; anything else yields nil. The value
type does not matter.

Examples:

	Keys(map[string]int{"b": 1, "a": 2}) => []string{"a", "b"}""",
    imports=['"reflect"', '"sort"'],
    solution="""rv := reflect.ValueOf(m)
if rv.Kind() != reflect.Map || rv.Type().Key().Kind() != reflect.String {
	return nil
}
out := make([]string, 0, rv.Len())
for _, k := range rv.MapKeys() {
	out = append(out, k.String())
}
sort.Strings(out)
return out""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestKeys(t *testing.T) {
	got := Keys(map[string]int{"b": 1, "a": 2, "c": 3})
	if !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("Keys = %v, want [a b c]", got)
	}
}

func TestKeysIgnoresTheValueType(t *testing.T) {
	got := Keys(map[string][]byte{"z": nil, "y": nil})
	if !reflect.DeepEqual(got, []string{"y", "z"}) {
		t.Errorf("Keys = %v, want [y z]", got)
	}
}

func TestKeysRejectsOtherShapes(t *testing.T) {
	for _, in := range []any{nil, 3, []string{"a"}, map[int]string{1: "a"}} {
		if got := Keys(in); got != nil {
			t.Errorf("Keys(%#v) = %v, want nil", in, got)
		}
	}
}

func TestKeysEmptyMap(t *testing.T) {
	if got := Keys(map[string]int{}); len(got) != 0 {
		t.Errorf("Keys = %v, want empty", got)
	}
}

func TestKeysIsDeterministic(t *testing.T) {
	m := map[string]int{"d": 1, "a": 2, "c": 3, "b": 4}
	first := Keys(m)
	for i := 0; i < 50; i++ {
		if got := Keys(m); !reflect.DeepEqual(got, first) {
			t.Fatalf("run %d = %v, want %v: map iteration order must be sorted away", i, got, first)
		}
	}
}
""",
    context="A diff tool prints two configuration maps side by side. Map iteration order is randomised, so the same input produces a different diff on every run.",
    task=[
        "Return the sorted keys of the string-keyed map `m`.",
        "Any value type is acceptable.",
        "Return nil for a non-map, a nil interface, or a map with non-string keys.",
    ],
    examples=[
        ('Keys(map[string]int{"b":1,"a":2})', "[a b]", None),
        ("Keys(map[int]string{1:\"a\"})", "<nil>", "Keys must be strings."),
        ("Keys(map[string]int{})", "[]", None),
    ],
    topics=[
        ("Type.Key()", "The map's key type is available without touching any entry."),
        ("Value.MapKeys", "Returns the keys as reflect Values, in unspecified order."),
        ("Determinism", "Sorting is what makes reflective output reproducible."),
    ],
    hint="Check the key kind through `rv.Type().Key()` before touching the entries.",
    intuition="Map iteration order in Go is deliberately randomised, so any reflective dump must sort. The key type is part of the map's type, so you can validate the shape before reading a single entry.",
    approach=[
        "Reject anything that is not a map with a string key type.",
        "Preallocate to `rv.Len()` and collect `k.String()` for each key.",
        "Sort and return.",
    ],
    walkthrough="`Keys` on a four-entry map collects the keys in whatever order the runtime offers, then sorts them — so fifty calls produce fifty identical slices.",
    pitfalls=[
        "`reflect.ValueOf(nil).Kind()` is invalid, which the map check already rejects.",
        "Using `fmt.Sprint(k)` instead of `k.String()`; the kind check has already guaranteed a string.",
    ],
)

P(
    "middle",
    name="callfn",
    title="Call A Function You Only Know At Run Time",
    sig="func CallInts(fn any, args ...int) ([]int, error)",
    doc="""CallInts calls fn with args and returns its int results.

fn must be a function taking exactly len(args) int parameters and
returning only ints. Anything else is an error, not a panic.

Examples:

	CallInts(func(a, b int) int { return a + b }, 1, 2) => []int{3}""",
    imports=['"errors"', '"reflect"'],
    extra="""// ErrSignature is returned when fn does not match the expected shape.
var ErrSignature = errors.New("fn must take and return only ints")""",
    solution="""rv := reflect.ValueOf(fn)
if rv.Kind() != reflect.Func {
	return nil, ErrSignature
}
t := rv.Type()
if t.IsVariadic() || t.NumIn() != len(args) {
	return nil, ErrSignature
}
for i := 0; i < t.NumIn(); i++ {
	if t.In(i).Kind() != reflect.Int {
		return nil, ErrSignature
	}
}
for i := 0; i < t.NumOut(); i++ {
	if t.Out(i).Kind() != reflect.Int {
		return nil, ErrSignature
	}
}
in := make([]reflect.Value, len(args))
for i, a := range args {
	in[i] = reflect.ValueOf(a)
}
res := rv.Call(in)
out := make([]int, len(res))
for i, r := range res {
	out[i] = int(r.Int())
}
return out, nil""",
    tests="""
import (
	"errors"
	"reflect"
	"testing"
)

func TestCallInts(t *testing.T) {
	got, err := CallInts(func(a, b int) int { return a + b }, 1, 2)
	if err != nil || !reflect.DeepEqual(got, []int{3}) {
		t.Errorf("CallInts = %v, %v, want [3], nil", got, err)
	}
}

func TestCallIntsMultipleResults(t *testing.T) {
	got, err := CallInts(func(a int) (int, int) { return a, -a }, 5)
	if err != nil || !reflect.DeepEqual(got, []int{5, -5}) {
		t.Errorf("CallInts = %v, %v, want [5 -5], nil", got, err)
	}
}

func TestCallIntsNoArgsNoResults(t *testing.T) {
	got, err := CallInts(func() {})
	if err != nil || len(got) != 0 {
		t.Errorf("CallInts = %v, %v, want empty, nil", got, err)
	}
}

func TestCallIntsBadSignatures(t *testing.T) {
	cases := []struct {
		name string
		fn   any
		args []int
	}{
		{"not a func", 3, nil},
		{"nil", nil, nil},
		{"wrong arity", func(a, b int) int { return a }, []int{1}},
		{"string param", func(s string) int { return 0 }, []int{1}},
		{"string result", func(a int) string { return "" }, []int{1}},
		{"variadic", func(a ...int) int { return 0 }, []int{1}},
	}
	for _, c := range cases {
		if _, err := CallInts(c.fn, c.args...); !errors.Is(err, ErrSignature) {
			t.Errorf("%s: err = %v, want ErrSignature", c.name, err)
		}
	}
}
""",
    context="A plugin registry stores handler functions as `any`. Calling one means checking, at run time, that it has the shape the registry promised.",
    task=[
        "Call `fn` with `args` and return its int results.",
        "Reject anything that is not a function, has the wrong arity, is variadic, or has a non-int parameter or result.",
        "A function with no results returns an empty slice, not an error.",
    ],
    examples=[
        ("CallInts(func(a, b int) int { return a+b }, 1, 2)", "[3], nil", None),
        ("CallInts(func(a int) (int,int) { return a,-a }, 5)", "[5 -5], nil", "Every result is collected."),
        ("CallInts(func(s string) int {...}, 1)", "ErrSignature", None),
    ],
    topics=[
        ("reflect.Value.Call", "Takes and returns `[]reflect.Value`; a mismatch panics, so validate first."),
        ("Type.NumIn / In / NumOut / Out", "The function's signature is fully inspectable."),
        ("IsVariadic", "A variadic function needs `CallSlice`, not `Call` — reject it here."),
        ("Turning panics into errors", "Validating up front is better than recovering afterwards."),
    ],
    hint="Every way `Call` can panic corresponds to a check you can make on the `Type` first.",
    intuition="`Call` is unforgiving: the wrong arity or the wrong argument type is a panic, not an error. The function's type carries everything needed to check the shape before you call, which turns a crash into a returned error.",
    approach=[
        "Verify the kind is func, it is not variadic, and the arity matches.",
        "Verify every parameter and every result is of kind int.",
        "Build the `[]reflect.Value` arguments, `Call`, and read the results with `Int()`.",
    ],
    walkthrough="For `func(a, b int) int` with args 1 and 2: arity 2 matches, both parameters are int, the single result is int. `Call` returns one Value holding 3.",
    pitfalls=[
        "Checking arity but not parameter types — `Call` panics on a type mismatch.",
        "Forgetting the variadic case, where `NumIn` counts the slice parameter as one.",
    ],
)

P(
    "middle",
    name="appendzero",
    title="Grow A Slice Whose Type You Do Not Know",
    sig="func AppendZero(slicePtr any, n int) error",
    doc="""AppendZero appends n zero values to the slice that slicePtr points at.

The element type comes from the slice itself, so one implementation
serves every slice type.

Examples:

	s := []int{1}; AppendZero(&s, 2) => s is [1 0 0]""",
    imports=['"errors"', '"reflect"'],
    extra="""// ErrTarget is returned when slicePtr is not a non-nil pointer to a slice.
var ErrTarget = errors.New("target must be a non-nil pointer to a slice")""",
    solution="""rv := reflect.ValueOf(slicePtr)
if rv.Kind() != reflect.Pointer || rv.IsNil() {
	return ErrTarget
}
rv = rv.Elem()
if rv.Kind() != reflect.Slice {
	return ErrTarget
}
if n < 0 {
	return ErrTarget
}
elem := rv.Type().Elem()
for i := 0; i < n; i++ {
	rv.Set(reflect.Append(rv, reflect.Zero(elem)))
}
return nil""",
    tests="""
import (
	"errors"
	"reflect"
	"testing"
)

func TestAppendZeroInts(t *testing.T) {
	s := []int{1}
	if err := AppendZero(&s, 2); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s, []int{1, 0, 0}) {
		t.Errorf("s = %v, want [1 0 0]", s)
	}
}

func TestAppendZeroStrings(t *testing.T) {
	s := []string{"a"}
	if err := AppendZero(&s, 1); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s, []string{"a", ""}) {
		t.Errorf("s = %q, want [a \\"\\"]", s)
	}
}

func TestAppendZeroStructs(t *testing.T) {
	type item struct{ A, B int }
	s := []item{{1, 2}}
	if err := AppendZero(&s, 1); err != nil {
		t.Fatal(err)
	}
	if len(s) != 2 || s[1] != (item{}) {
		t.Errorf("s = %v, want a zero item appended", s)
	}
}

func TestAppendZeroNilSlice(t *testing.T) {
	var s []int
	if err := AppendZero(&s, 3); err != nil {
		t.Fatal(err)
	}
	if len(s) != 3 {
		t.Errorf("len = %d, want 3", len(s))
	}
}

func TestAppendZeroNoop(t *testing.T) {
	s := []int{1}
	if err := AppendZero(&s, 0); err != nil || len(s) != 1 {
		t.Errorf("s = %v, err = %v, want [1], nil", s, err)
	}
}

func TestAppendZeroBadTarget(t *testing.T) {
	s := []int{1}
	for _, c := range []any{s, nil, (*[]int)(nil), new(int)} {
		if err := AppendZero(c, 1); !errors.Is(err, ErrTarget) {
			t.Errorf("AppendZero(%#v) = %v, want ErrTarget", c, err)
		}
	}
	if err := AppendZero(&s, -1); !errors.Is(err, ErrTarget) {
		t.Errorf("negative n = %v, want ErrTarget", err)
	}
}
""",
    context="A fixture helper pads every slice in a test table to the same length. Writing it once per element type produced six near-identical functions that drifted apart.",
    task=[
        "Append `n` zero values to the slice `slicePtr` points at.",
        "Work for any element type, including structs and a nil slice.",
        "Return `ErrTarget` for a non-pointer, a nil pointer, a non-slice, or a negative `n`.",
    ],
    examples=[
        ("s := []int{1}; AppendZero(&s, 2)", "s is [1 0 0]", None),
        ('s := []string{"a"}; AppendZero(&s, 1)', 's is ["a" ""]', "The zero value follows the element type."),
        ("AppendZero(&s, 0)", "s unchanged", None),
    ],
    topics=[
        ("reflect.Append", "The reflective twin of `append`; it returns a new Value that must be stored back."),
        ("reflect.Zero", "Produces the zero Value for any type."),
        ("Type.Elem()", "A slice type knows its element type."),
        ("Set writes the result back", "`Append` alone does not modify the caller's slice."),
    ],
    hint="`reflect.Append` returns a new slice Value. Something has to happen to it.",
    intuition="Reflection mirrors the language: `append` returns a new slice header, so the reflective version does too. The extra step is storing it back through the addressable Value you got from `Elem`.",
    approach=[
        "Validate the pointer, step to the slice with `Elem`, and reject a negative `n`.",
        "Take the element type from `rv.Type().Elem()`.",
        "Append `reflect.Zero(elem)` `n` times, storing each result with `rv.Set`.",
    ],
    walkthrough="`&s` where `s` is `[]int{1}` gives an addressable slice Value. Two appends of `reflect.Zero(int)` produce [1 0] then [1 0 0], each stored back into the caller's slice header.",
    pitfalls=[
        "Discarding `Append`'s result — the caller's slice never changes.",
        "Building the zero value as `reflect.ValueOf(0)`, which only works for int.",
    ],
)

P(
    "middle",
    name="structclone",
    title="Copy A Struct Without Knowing Its Type",
    sig="func Clone(v any) any",
    doc="""Clone returns a copy of the struct v, as a value of the same type.

The copy is shallow: fields are assigned, so slices and maps inside it
still share their storage with v.

Examples:

	Clone(pt{1, 2}) => pt{1, 2}, a distinct value""",
    imports=['"reflect"'],
    solution="""rv := reflect.ValueOf(v)
if !rv.IsValid() || rv.Kind() != reflect.Struct {
	return nil
}
out := reflect.New(rv.Type()).Elem()
out.Set(rv)
return out.Interface()""",
    tests="""
import (
	"reflect"
	"testing"
)

type pt struct {
	X, Y int
	Tags []string
}

func TestCloneCopiesFields(t *testing.T) {
	in := pt{X: 1, Y: 2}
	got := Clone(in)
	out, ok := got.(pt)
	if !ok {
		t.Fatalf("Clone returned %T, want pt", got)
	}
	if out.X != 1 || out.Y != 2 || out.Tags != nil {
		t.Errorf("Clone = %+v, want {1 2 []}", out)
	}
}

func TestCloneIsIndependentForValueFields(t *testing.T) {
	in := pt{X: 1}
	out := Clone(in).(pt)
	out.X = 99
	if in.X != 1 {
		t.Errorf("in.X = %d, want 1", in.X)
	}
}

func TestCloneIsShallow(t *testing.T) {
	in := pt{Tags: []string{"a"}}
	out := Clone(in).(pt)
	out.Tags[0] = "changed"
	if in.Tags[0] != "changed" {
		t.Error("the clone copied the slice; a shallow copy shares it")
	}
}

func TestCloneRejectsNonStructs(t *testing.T) {
	for _, in := range []any{nil, 3, []int{1}, &pt{}} {
		if got := Clone(in); got != nil {
			t.Errorf("Clone(%#v) = %v, want nil", in, got)
		}
	}
}

func TestClonePreservesTheType(t *testing.T) {
	type other struct{ A string }
	got := Clone(other{A: "x"})
	if reflect.TypeOf(got) != reflect.TypeOf(other{}) {
		t.Errorf("type = %T, want other", got)
	}
}
""",
    context="A snapshot helper stores a copy of whatever struct it is handed. Type-asserting through a list of known types covered four of them and missed the fifth.",
    task=[
        "Return a copy of the struct `v`, with the same dynamic type.",
        "The copy is shallow — reference fields stay shared.",
        "Return nil for a nil interface or anything that is not a struct.",
    ],
    examples=[
        ("Clone(pt{1,2})", "pt{1,2}", "Same type, distinct value."),
        ("out := Clone(in).(pt); out.X = 99", "in.X unchanged", "Value fields are copied."),
        ("Clone(&pt{})", "<nil>", "A pointer is not a struct."),
    ],
    topics=[
        ("reflect.New", "Allocates a new value of a type and returns a pointer Value to it."),
        ("Value.Set", "Struct assignment through reflection is the same shallow copy as `=`."),
        ("Value.Interface", "Boxes the reflected value back into an `any` with its real type."),
    ],
    hint="`reflect.New(t)` gives a pointer. You want what it points at.",
    intuition="Reflection can build a value of any type it is shown. `New` allocates one, `Elem` gives the addressable value inside, and `Set` performs exactly the assignment the language would.",
    approach=[
        "Reject a nil interface and non-struct kinds.",
        "`reflect.New(rv.Type()).Elem()` for a fresh addressable value.",
        "`Set` the original into it and return `Interface()`.",
    ],
    walkthrough="`Clone(pt{1,2})` allocates a new `pt`, assigns the original's fields, and boxes the result. Assigning to `out.X` afterwards cannot reach `in`, but writing through `out.Tags` can — that is what shallow means.",
    pitfalls=[
        "Returning `out` (a `reflect.Value`) instead of `out.Interface()`.",
        "Promising a deep copy; slices and maps are still shared.",
    ],
)

# ---------------------------------------------------------------- senior -----

P(
    "senior",
    name="typednil",
    title="The Interface That Is Not Nil",
    sig="func IsNilValue(v any) bool",
    doc="""IsNilValue reports whether v is nil or holds a nil pointer, map,
slice, channel, function or interface.

An interface holding a typed nil pointer is not == nil, which is the trap
this function exists to close.

Examples:

	var p *int; IsNilValue(p) => true""",
    imports=['"reflect"'],
    solution="""if v == nil {
	return true
}
rv := reflect.ValueOf(v)
switch rv.Kind() {
case reflect.Pointer, reflect.Map, reflect.Slice, reflect.Chan,
	reflect.Func, reflect.Interface, reflect.UnsafePointer:
	return rv.IsNil()
default:
	return false
}""",
    tests="""
import "testing"

type myErr struct{}

func (*myErr) Error() string { return "boom" }

func TestIsNilValue(t *testing.T) {
	var p *int
	var m map[string]int
	var s []int
	var c chan int
	var f func()
	cases := []struct {
		name string
		in   any
		want bool
	}{
		{"untyped nil", nil, true},
		{"nil pointer", p, true},
		{"nil map", m, true},
		{"nil slice", s, true},
		{"nil chan", c, true},
		{"nil func", f, true},
		{"non-nil pointer", new(int), false},
		{"empty slice", []int{}, false},
		{"zero int", 0, false},
		{"empty string", "", false},
		{"struct", struct{}{}, false},
	}
	for _, c := range cases {
		if got := IsNilValue(c.in); got != c.want {
			t.Errorf("%s: IsNilValue = %v, want %v", c.name, got, c.want)
		}
	}
}

func TestIsNilValueCatchesTheTypedNilError(t *testing.T) {
	var e error = (*myErr)(nil)
	if e == nil {
		t.Fatal("the fixture is wrong: a typed nil error is not == nil")
	}
	if !IsNilValue(e) {
		t.Error("IsNilValue = false for a typed nil error: this is the case the function is for")
	}
}

func TestIsNilValueDoesNotPanic(t *testing.T) {
	for _, in := range []any{0, "", 1.5, [2]int{}, struct{ A int }{}} {
		_ = IsNilValue(in)
	}
}
""",
    context="A handler returns `err` from a helper that declares `*ValidationError` as its result type. Every call site sees a non-nil error, every request fails validation, and the error message is empty.",
    task=[
        "Report whether `v` is nil or wraps a nil pointer, map, slice, channel, function or interface.",
        "Values that cannot be nil — ints, strings, structs, arrays — report false.",
        "Never panic, whatever the input.",
    ],
    examples=[
        ("var p *int; IsNilValue(p)", "true", "An interface holding a typed nil pointer."),
        ("IsNilValue([]int{})", "false", "Empty is not nil."),
        ("IsNilValue(0)", "false", "An int cannot be nil."),
    ],
    topics=[
        ("Interface representation", "An interface is a (type, value) pair; a typed nil has a type, so it is not == nil."),
        ("Value.IsNil", "Valid only for the nilable kinds — it panics on the rest."),
        ("Why the trap bites", "Assigning a typed nil pointer to an `error` variable makes it non-nil forever after."),
    ],
    hint="`v == nil` is the first check, not the only one. And `IsNil` panics on an int.",
    intuition="An interface value carries a type word and a data word. A nil `*T` fills in the type word, so the interface is not nil even though the pointer inside is. Reflection is the only way to look inside and ask about the data word alone.",
    approach=[
        "Return true immediately for an untyped nil.",
        "Switch on the kind and call `IsNil` only for the kinds that can be nil.",
        "Return false for everything else.",
    ],
    walkthrough="`var e error = (*myErr)(nil)` has type `*myErr` and value nil. `e == nil` is false; `reflect.ValueOf(e).IsNil()` is true, which is the answer the caller wanted.",
    pitfalls=[
        "Calling `IsNil` without the kind switch — that panics on an int.",
        "Fixing the symptom at the call site instead of never assigning a typed nil to an interface in the first place.",
    ],
)

P(
    "senior",
    name="skipunexported",
    title="Sum The Fields You Are Allowed To Read",
    mode="bug",
    sig="func SumInts(v any) int64",
    doc="""SumInts returns the total of v's exported int fields.

Unexported fields can be read as reflect Values but not converted back
through Interface, and reaching for them panics.

Examples:

	SumInts(mix{A: 1, b: 2}) => 1""",
    imports=['"reflect"'],
    buggy="""rv := reflect.ValueOf(v)
if !rv.IsValid() || rv.Kind() != reflect.Struct {
	return 0
}
var total int64
for i := 0; i < rv.NumField(); i++ {
	f := rv.Field(i)
	if f.Kind() != reflect.Int {
		continue
	}
	total += f.Interface().(int64)
}
return total""",
    solution="""rv := reflect.ValueOf(v)
if !rv.IsValid() || rv.Kind() != reflect.Struct {
	return 0
}
rt := rv.Type()
var total int64
for i := 0; i < rv.NumField(); i++ {
	if !rt.Field(i).IsExported() {
		continue
	}
	f := rv.Field(i)
	if f.Kind() != reflect.Int {
		continue
	}
	total += f.Int()
}
return total""",
    tests="""
import "testing"

type mix struct {
	A      int
	B      int
	hidden int
	Name   string
}

func TestSumInts(t *testing.T) {
	if got := SumInts(mix{A: 1, B: 2, hidden: 100, Name: "x"}); got != 3 {
		t.Errorf("SumInts = %d, want 3: only exported int fields count", got)
	}
	if got := SumInts(mix{}); got != 0 {
		t.Errorf("SumInts = %d, want 0", got)
	}
}

func TestSumIntsNegative(t *testing.T) {
	if got := SumInts(mix{A: -5, B: 5}); got != 0 {
		t.Errorf("SumInts = %d, want 0", got)
	}
}

func TestSumIntsRejectsNonStructs(t *testing.T) {
	for _, in := range []any{nil, 3, []int{1}, &mix{}} {
		if got := SumInts(in); got != 0 {
			t.Errorf("SumInts(%#v) = %d, want 0", in, got)
		}
	}
}

func TestSumIntsAllExported(t *testing.T) {
	type open struct{ A, B, C int }
	if got := SumInts(open{1, 2, 3}); got != 6 {
		t.Errorf("SumInts = %d, want 6", got)
	}
}
""",
    context="An audit helper totals the numeric fields of whatever record it is handed. It works in the unit tests and panics the moment a struct with a private field reaches it.",
    task=[
        "Total the exported int fields of `v`.",
        "Skip unexported fields and fields of any other kind.",
        "Return 0 for a nil interface or a non-struct.",
        "Fix the single bug so no input can make the function panic.",
    ],
    examples=[
        ("SumInts(mix{A:1, B:2, hidden:100})", "3", "`hidden` is unexported."),
        ("SumInts(mix{})", "0", None),
        ("SumInts(&mix{})", "0", "A pointer is not a struct."),
    ],
    topics=[
        ("CanInterface / IsExported", "An unexported field's Value cannot be boxed back into an `any`."),
        ("Typed accessors", "`f.Int()` reads the value without going through `Interface`."),
        ("Type and Value in step", "Export status lives on the `StructField`, not the `Value`."),
    ],
    hint="Two things are wrong with one line. What does `Interface()` require, and what does an `int` field actually assert to?",
    intuition="Reflection enforces the language's visibility rules: it will show you an unexported field but refuse to hand it out as an `any`. And a field of kind int holds an `int`, not an `int64` — the type assertion was never going to succeed.",
    approach=[
        "Consult `rt.Field(i).IsExported()` and skip unexported fields.",
        "Read the value with `f.Int()`, which returns int64 for every signed integer kind.",
    ],
    walkthrough="For `mix{A:1, B:2, hidden:100}`: A and B are exported ints and contribute 3; `hidden` is skipped before anything can panic; `Name` fails the kind check.",
    pitfalls=[
        "`f.Interface()` on an unexported field panics with \"cannot return value obtained from unexported field\".",
        "`f.Interface().(int64)` panics even on an exported int field — the dynamic type is `int`.",
    ],
)

P(
    "senior",
    name="fieldindex",
    title="Look The Field Up Once, Not Once Per Row",
    sig="func SumColumn(rows any, field string) (int64, error)",
    doc="""SumColumn totals the named int field across a slice of structs.

Resolving the field by name is a string search through the struct's field
table; doing it per row makes the cost O(rows x fields).

Examples:

	SumColumn([]rec{{N: 1}, {N: 2}}, "N") => 3, nil""",
    imports=['"errors"', '"reflect"'],
    extra="""// ErrShape is returned when rows is not a slice of structs with an int
// field of the given name.
var ErrShape = errors.New("rows must be a slice of structs with that int field")""",
    solution="""rv := reflect.ValueOf(rows)
if rv.Kind() != reflect.Slice {
	return 0, ErrShape
}
et := rv.Type().Elem()
if et.Kind() != reflect.Struct {
	return 0, ErrShape
}
sf, ok := et.FieldByName(field)
if !ok || !sf.IsExported() || sf.Type.Kind() != reflect.Int {
	return 0, ErrShape
}
idx := sf.Index
var total int64
for i := 0; i < rv.Len(); i++ {
	total += rv.Index(i).FieldByIndex(idx).Int()
}
return total, nil""",
    tests="""
import (
	"errors"
	"testing"
)

type rec struct {
	N      int
	M      int
	Label  string
	hidden int
}

func TestSumColumn(t *testing.T) {
	rows := []rec{{N: 1, M: 10}, {N: 2, M: 20}, {N: 3, M: 30}}
	if got, err := SumColumn(rows, "N"); err != nil || got != 6 {
		t.Errorf("SumColumn(N) = %d, %v, want 6, nil", got, err)
	}
	if got, err := SumColumn(rows, "M"); err != nil || got != 60 {
		t.Errorf("SumColumn(M) = %d, %v, want 60, nil", got, err)
	}
}

func TestSumColumnEmpty(t *testing.T) {
	if got, err := SumColumn([]rec{}, "N"); err != nil || got != 0 {
		t.Errorf("SumColumn = %d, %v, want 0, nil", got, err)
	}
}

func TestSumColumnBadShape(t *testing.T) {
	cases := []struct {
		name  string
		rows  any
		field string
	}{
		{"not a slice", rec{}, "N"},
		{"nil", nil, "N"},
		{"slice of ints", []int{1}, "N"},
		{"missing field", []rec{{}}, "Nope"},
		{"wrong kind", []rec{{}}, "Label"},
		{"unexported", []rec{{}}, "hidden"},
	}
	for _, c := range cases {
		if _, err := SumColumn(c.rows, c.field); !errors.Is(err, ErrShape) {
			t.Errorf("%s: err = %v, want ErrShape", c.name, err)
		}
	}
}

func TestSumColumnResolvesTheFieldOnce(t *testing.T) {
	rows := make([]rec, 4096)
	for i := range rows {
		rows[i].N = 1
	}
	got, err := SumColumn(rows, "N")
	if err != nil || got != 4096 {
		t.Fatalf("SumColumn = %d, %v, want 4096, nil", got, err)
	}
	n := testing.AllocsPerRun(20, func() { _, _ = SumColumn(rows, "N") })
	if n > 4 {
		t.Errorf("SumColumn made %v allocations for 4096 rows, want a handful: resolve the field once", n)
	}
}

func BenchmarkSumColumn(b *testing.B) {
	rows := make([]rec, 4096)
	for i := range rows {
		rows[i].N = i
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = SumColumn(rows, "N")
	}
}
""",
    context="A reporting layer sums a column by name over a few million rows. The profile is dominated by `FieldByName`, which is doing a string comparison per field per row.",
    task=[
        "Total the named int field over the slice of structs `rows`.",
        "Resolve the field against the element type once, before the loop.",
        "Return `ErrShape` for a non-slice, a non-struct element, or a field that is missing, unexported, or not an int.",
        "An empty slice totals 0.",
    ],
    examples=[
        ('SumColumn([]rec{{N:1},{N:2}}, "N")', "3, nil", None),
        ('SumColumn([]rec{{}}, "Label")', "ErrShape", "Not an int field."),
        ('SumColumn([]rec{}, "N")', "0, nil", None),
    ],
    topics=[
        ("Type metadata is per type, not per value", "The field's position is the same for every row."),
        ("StructField.Index / FieldByIndex", "The resolved position is an index path, usable without another name search."),
        ("Validate once", "Shape errors are properties of the type, so they can be decided before the loop."),
        ("Reflection's real cost", "Name resolution, not field access, is what dominates."),
    ],
    hint="`FieldByName` returns a `StructField`. What is in its `Index`?",
    intuition="Every row in the slice has the same type, so the field's offset is decided once by the type, not once per row. `FieldByName` hands you that position — keep it and index directly.",
    approach=[
        "Validate the slice and its struct element type.",
        "`FieldByName` once on the element type; check exported and int.",
        "Keep `sf.Index` and use `FieldByIndex` inside the loop.",
    ],
    walkthrough="For 4096 rows, `FieldByName` runs once and the loop does 4096 direct index lookups. Calling it per row would run 4096 string searches over the field table instead.",
    pitfalls=[
        "Calling `rv.Index(i).FieldByName(field)` inside the loop — correct and slow.",
        "Validating the field inside the loop, so a bad shape is reported 4096 times.",
    ],
)

P(
    "senior",
    name="walkfields",
    title="Total Every Int, However Deep",
    sig="func DeepSum(v any) int64",
    doc="""DeepSum totals every exported int field in v, descending into nested
structs, slices of structs and pointers.

A nil pointer contributes nothing. Cycles are not part of the input.

Examples:

	DeepSum(outer{N: 1, In: inner{M: 2}}) => 3""",
    imports=['"reflect"'],
    solution="""return deepSum(reflect.ValueOf(v))""",
    extra="""func deepSum(rv reflect.Value) int64 {
	if !rv.IsValid() {
		return 0
	}
	switch rv.Kind() {
	case reflect.Int:
		return rv.Int()
	case reflect.Pointer, reflect.Interface:
		if rv.IsNil() {
			return 0
		}
		return deepSum(rv.Elem())
	case reflect.Struct:
		rt := rv.Type()
		var total int64
		for i := 0; i < rv.NumField(); i++ {
			if !rt.Field(i).IsExported() {
				continue
			}
			total += deepSum(rv.Field(i))
		}
		return total
	case reflect.Slice, reflect.Array:
		var total int64
		for i := 0; i < rv.Len(); i++ {
			total += deepSum(rv.Index(i))
		}
		return total
	default:
		return 0
	}
}""",
    tests="""
import "testing"

type inner struct {
	M      int
	hidden int
}

type outer struct {
	N     int
	In    inner
	Ptr   *inner
	List  []inner
	Label string
}

func TestDeepSumFlat(t *testing.T) {
	if got := DeepSum(inner{M: 5}); got != 5 {
		t.Errorf("DeepSum = %d, want 5", got)
	}
}

func TestDeepSumNested(t *testing.T) {
	v := outer{
		N:    1,
		In:   inner{M: 2},
		Ptr:  &inner{M: 4},
		List: []inner{{M: 8}, {M: 16}},
	}
	if got := DeepSum(v); got != 31 {
		t.Errorf("DeepSum = %d, want 31", got)
	}
}

func TestDeepSumNilPointer(t *testing.T) {
	if got := DeepSum(outer{N: 1}); got != 1 {
		t.Errorf("DeepSum = %d, want 1: a nil pointer contributes nothing", got)
	}
}

func TestDeepSumSkipsUnexported(t *testing.T) {
	if got := DeepSum(inner{M: 1, hidden: 100}); got != 1 {
		t.Errorf("DeepSum = %d, want 1", got)
	}
}

func TestDeepSumOtherInputs(t *testing.T) {
	if got := DeepSum(7); got != 7 {
		t.Errorf("DeepSum(7) = %d, want 7", got)
	}
	if got := DeepSum([]int{1, 2, 3}); got != 6 {
		t.Errorf("DeepSum = %d, want 6", got)
	}
	if got := DeepSum(nil); got != 0 {
		t.Errorf("DeepSum(nil) = %d, want 0", got)
	}
	if got := DeepSum("x"); got != 0 {
		t.Errorf("DeepSum = %d, want 0", got)
	}
}

func TestDeepSumPointerInput(t *testing.T) {
	if got := DeepSum(&inner{M: 9}); got != 9 {
		t.Errorf("DeepSum = %d, want 9", got)
	}
}
""",
    context="A metrics exporter needs the total of every counter in a nested settings tree. The tree gains a level every quarter and the hand-written walker is always one release behind.",
    task=[
        "Total every exported int reachable from `v`.",
        "Descend into nested structs, pointers, interfaces, slices and arrays.",
        "Skip unexported fields; a nil pointer contributes 0.",
        "A bare int is its own total; anything else contributes 0.",
    ],
    examples=[
        ("DeepSum(outer{N:1, In:inner{M:2}})", "3", None),
        ("DeepSum(outer{N:1})", "1", "The nil `Ptr` contributes nothing."),
        ("DeepSum([]int{1,2,3})", "6", "Slices are walked too."),
    ],
    topics=[
        ("Recursive reflection", "One function per kind, recursing on the contained Values."),
        ("Value.Elem", "Steps through a pointer or an interface to what it holds."),
        ("Nil guards", "`Elem` on a nil pointer yields an invalid Value."),
        ("Export status on the way down", "Every struct level must be filtered, not just the top one."),
    ],
    hint="One switch on the kind, four interesting cases, and a recursive call in each.",
    intuition="A value tree is walked the same way any tree is: handle the leaf kind, recurse on the container kinds, and stop everywhere else. Reflection turns \"what shape is this node\" into a `Kind` switch.",
    approach=[
        "Return 0 for an invalid Value.",
        "Int is the leaf: return it.",
        "Pointer and interface: return 0 when nil, otherwise recurse on `Elem`.",
        "Struct: recurse on every exported field. Slice and array: recurse on every element.",
    ],
    walkthrough="For the nested `outer`: N gives 1, In.M gives 2, *Ptr gives 4, and the two list entries give 8 and 16 — 31 in total. `Label` falls through the switch to 0.",
    pitfalls=[
        "Recursing into unexported fields, which panics as soon as something tries to read them out.",
        "Calling `Elem` on a nil pointer and then `Kind` on the invalid result.",
    ],
)

P(
    "senior",
    name="settableslice",
    title="Write Into A Slice Through Reflection",
    mode="bug",
    sig="func Double(slice any) error",
    doc="""Double multiplies every element of the int slice in place.

reflect.ValueOf gives a copy of the interface's contents, but a slice's
elements live in the shared backing array — which is exactly why the
elements are settable and the slice header is not.

Examples:

	s := []int{1, 2}; Double(s) => s is [2 4]""",
    imports=['"errors"', '"reflect"'],
    extra="""// ErrShape is returned when slice is not a slice of ints.
var ErrShape = errors.New("argument must be a slice of ints")""",
    buggy="""rv := reflect.ValueOf(slice)
if rv.Kind() != reflect.Slice || rv.Type().Elem().Kind() != reflect.Int {
	return ErrShape
}
for i := 0; i < rv.Len(); i++ {
	e := rv.Index(i)
	v := e.Int()
	e = reflect.ValueOf(int(v * 2))
	_ = e
}
return nil""",
    solution="""rv := reflect.ValueOf(slice)
if rv.Kind() != reflect.Slice || rv.Type().Elem().Kind() != reflect.Int {
	return ErrShape
}
for i := 0; i < rv.Len(); i++ {
	e := rv.Index(i)
	e.SetInt(e.Int() * 2)
}
return nil""",
    tests="""
import (
	"errors"
	"reflect"
	"testing"
)

func TestDouble(t *testing.T) {
	s := []int{1, 2, 3}
	if err := Double(s); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s, []int{2, 4, 6}) {
		t.Errorf("s = %v, want [2 4 6]: the elements were not written", s)
	}
}

func TestDoubleEmptyAndNil(t *testing.T) {
	if err := Double([]int{}); err != nil {
		t.Errorf("Double([]) = %v, want nil", err)
	}
	if err := Double([]int(nil)); err != nil {
		t.Errorf("Double(nil slice) = %v, want nil", err)
	}
}

func TestDoubleWritesThroughAView(t *testing.T) {
	s := []int{1, 2, 3, 4}
	if err := Double(s[1:3]); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s, []int{1, 4, 6, 4}) {
		t.Errorf("s = %v, want [1 4 6 4]", s)
	}
}

func TestDoubleBadShape(t *testing.T) {
	for _, in := range []any{nil, 3, []string{"a"}, map[string]int{}} {
		if err := Double(in); !errors.Is(err, ErrShape) {
			t.Errorf("Double(%#v) = %v, want ErrShape", in, err)
		}
	}
}
""",
    context="A normalisation pass written with reflection runs without error and changes nothing. The author adds logging inside the loop, sees the right values computed, and cannot see where they go.",
    task=[
        "Double every element of the int slice, in place.",
        "Return `ErrShape` for anything that is not a slice of ints.",
        "Fix the single bug so the writes actually reach the caller's slice.",
    ],
    examples=[
        ("s := []int{1,2,3}; Double(s)", "s is [2 4 6]", None),
        ("s := []int{1,2,3,4}; Double(s[1:3])", "s is [1 4 6 4]", "Only the view is written."),
        ("Double(3)", "ErrShape", None),
    ],
    topics=[
        ("Assignment to a Value does nothing", "`e = reflect.ValueOf(x)` rebinds a local variable, it does not store anything."),
        ("Value.SetInt", "The write path; it requires the Value to be addressable."),
        ("Why slice elements are settable", "The elements live in the backing array, which the header points at."),
    ],
    hint="`e` is a handle, not a slot. Which method writes through a handle?",
    intuition="A `reflect.Value` is a handle onto storage. Assigning to the Go variable holding the handle changes only the variable — the storage is written through `Set` methods, and only when the handle is addressable.",
    approach=[
        "Take each element's Value with `rv.Index(i)`.",
        "Write with `e.SetInt(e.Int() * 2)`.",
    ],
    walkthrough="`rv.Index(0)` is an addressable handle onto `s[0]`. `SetInt` writes 2 into it. The buggy line built a new, unaddressable Value and dropped it.",
    pitfalls=[
        "Expecting `reflect.ValueOf(s)` to be unsettable to make the elements unsettable too — the header is a copy, the array is not.",
        "`e.Set(reflect.ValueOf(v*2))` also works; the point is that a `Set` call is what is missing.",
    ],
)

# ----------------------------------------------------------------- staff -----

P(
    "staff",
    name="decodecache",
    title="Resolve The Layout Once Per Type, Safely",
    sig="func Decode(src map[string]string, dst any) error",
    doc="""Decode fills dst's string fields from src by their `env` tag, caching
each struct type's tag-to-index layout.

The cache is shared by concurrent callers, so it must be safe under
parallel use — and must not resolve the layout twice for one type.

Examples:

	Decode(map[string]string{"H": "x"}, &cfg) => cfg.H == "x" """,
    imports=['"errors"', '"reflect"', '"sync"'],
    extra="""// ErrTarget is returned when dst is not a non-nil pointer to a struct.
var ErrTarget = errors.New("dst must be a non-nil pointer to a struct")

// layouts caches the tag-to-field-index map for each struct type.
var layouts sync.Map // reflect.Type -> map[string]int

// layoutOf returns the tag-to-index map for t, computing it at most once
// per type as far as any caller can observe.
func layoutOf(t reflect.Type) map[string]int {
	if v, ok := layouts.Load(t); ok {
		return v.(map[string]int)
	}
	m := make(map[string]int, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() || f.Type.Kind() != reflect.String {
			continue
		}
		if key, ok := f.Tag.Lookup("env"); ok && key != "" && key != "-" {
			m[key] = i
		}
	}
	actual, _ := layouts.LoadOrStore(t, m)
	return actual.(map[string]int)
}""",
    solution="""rv := reflect.ValueOf(dst)
if rv.Kind() != reflect.Pointer || rv.IsNil() {
	return ErrTarget
}
rv = rv.Elem()
if rv.Kind() != reflect.Struct {
	return ErrTarget
}
layout := layoutOf(rv.Type())
for key, i := range layout {
	if s, ok := src[key]; ok {
		rv.Field(i).SetString(s)
	}
}
return nil""",
    tests="""
import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

type cfg struct {
	Host   string `env:"HOST"`
	Region string `env:"REGION"`
	Plain  string
	hidden string `env:"HIDDEN"`
	Count  int    `env:"COUNT"`
}

func TestDecode(t *testing.T) {
	var c cfg
	src := map[string]string{"HOST": "h", "REGION": "r", "HIDDEN": "x", "COUNT": "5"}
	if err := Decode(src, &c); err != nil {
		t.Fatal(err)
	}
	if c.Host != "h" || c.Region != "r" {
		t.Errorf("c = %+v, want Host=h Region=r", c)
	}
	if c.Plain != "" || c.hidden != "" || c.Count != 0 {
		t.Errorf("c = %+v: untagged, unexported and non-string fields must be left alone", c)
	}
}

func TestDecodeMissingKeys(t *testing.T) {
	c := cfg{Host: "keep"}
	if err := Decode(map[string]string{"REGION": "r"}, &c); err != nil {
		t.Fatal(err)
	}
	if c.Host != "keep" || c.Region != "r" {
		t.Errorf("c = %+v", c)
	}
}

func TestDecodeBadTarget(t *testing.T) {
	for _, dst := range []any{cfg{}, nil, (*cfg)(nil), new(int)} {
		if err := Decode(nil, dst); !errors.Is(err, ErrTarget) {
			t.Errorf("Decode(%#v) = %v, want ErrTarget", dst, err)
		}
	}
}

func TestDecodeConcurrent(t *testing.T) {
	const workers = 16
	var wg sync.WaitGroup
	wg.Add(workers)
	errs := make([]error, workers)
	got := make([]cfg, workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			src := map[string]string{
				"HOST":   fmt.Sprintf("h%d", w),
				"REGION": fmt.Sprintf("r%d", w),
			}
			for i := 0; i < 200; i++ {
				var c cfg
				if err := Decode(src, &c); err != nil {
					errs[w] = err
					return
				}
				if c.Host != src["HOST"] || c.Region != src["REGION"] {
					errs[w] = fmt.Errorf("worker %d got %+v", w, c)
					return
				}
				got[w] = c
			}
		}(w)
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestDecodeManyTypes(t *testing.T) {
	type a struct {
		V string `env:"V"`
	}
	type b struct {
		W string `env:"W"`
	}
	var av a
	var bv b
	if err := Decode(map[string]string{"V": "1", "W": "2"}, &av); err != nil {
		t.Fatal(err)
	}
	if err := Decode(map[string]string{"V": "1", "W": "2"}, &bv); err != nil {
		t.Fatal(err)
	}
	if av.V != "1" || bv.W != "2" {
		t.Errorf("av = %+v, bv = %+v", av, bv)
	}
}
""",
    context="A decoder walks the struct's field table on every call. It shows up in the profile of a service that decodes the same three config types a hundred thousand times a second.",
    task=[
        "Fill `dst`'s exported string fields from `src`, matching by `env` tag.",
        "Use the shared `layoutOf` cache so a type's field table is walked once.",
        "Leave untagged, unexported, non-string fields and missing keys alone.",
        "Return `ErrTarget` unless `dst` is a non-nil pointer to a struct.",
        "Correct under concurrent use — many goroutines decoding at once.",
    ],
    examples=[
        ('Decode(map[string]string{"HOST":"h"}, &cfg)', "nil, cfg.Host is \"h\"", None),
        ("Decode(src, cfg{})", "ErrTarget", None),
        ("16 goroutines x 200 decodes", "every result correct", "The cache is shared; the targets are not."),
    ],
    topics=[
        ("Type metadata is immutable", "A resolved layout is valid forever, which is what makes caching sound."),
        ("sync.Map for read-heavy caches", "Loads after the first write take no lock."),
        ("LoadOrStore races benignly", "Two goroutines may both compute the layout; only one is published, and both are equal."),
        ("Shared cache, private targets", "The cache is read-only after publication; each caller writes its own struct."),
    ],
    hint="The cache is already written for you. The body is a validate, a lookup and a loop over the layout.",
    intuition="Reflection's expensive part is asking the type questions, and the answers never change. Cache the answer per type; the per-call work then drops to a map lookup and a few field writes.",
    approach=[
        "Validate `dst` and step to the struct.",
        "Fetch the layout for `rv.Type()` from the cache.",
        "For each cached tag-to-index pair, set the field when `src` has the key.",
    ],
    walkthrough="The first `Decode` of `cfg` walks five fields and stores a two-entry map. Every later call — from any goroutine — does one `sync.Map` load and up to two `SetString` calls.",
    pitfalls=[
        "Ranging the struct's fields instead of the cached layout, which reintroduces the walk.",
        "Caching a `reflect.Value` instead of an index; Values are bound to a particular variable.",
        "Guarding a plain map with no lock — the concurrency test is there to catch it.",
    ],
)

P(
    "staff",
    name="deepcopy",
    title="A Copy That Shares Nothing",
    sig="func DeepCopy(v any) any",
    doc="""DeepCopy returns a copy of v that shares no mutable storage with it.

Structs, slices, maps, arrays and pointers are copied recursively; scalars
and strings are copied by value. Cycles are not part of the input.

Examples:

	DeepCopy(node{Tags: []string{"a"}}) => an independent node""",
    imports=['"reflect"'],
    extra="""func deepCopy(rv reflect.Value) reflect.Value {
	switch rv.Kind() {
	case reflect.Pointer:
		if rv.IsNil() {
			return rv
		}
		out := reflect.New(rv.Type().Elem())
		out.Elem().Set(deepCopy(rv.Elem()))
		return out
	case reflect.Interface:
		if rv.IsNil() {
			return rv
		}
		out := reflect.New(rv.Type()).Elem()
		out.Set(deepCopy(rv.Elem()))
		return out
	case reflect.Slice:
		if rv.IsNil() {
			return rv
		}
		out := reflect.MakeSlice(rv.Type(), rv.Len(), rv.Len())
		for i := 0; i < rv.Len(); i++ {
			out.Index(i).Set(deepCopy(rv.Index(i)))
		}
		return out
	case reflect.Array:
		out := reflect.New(rv.Type()).Elem()
		for i := 0; i < rv.Len(); i++ {
			out.Index(i).Set(deepCopy(rv.Index(i)))
		}
		return out
	case reflect.Map:
		if rv.IsNil() {
			return rv
		}
		out := reflect.MakeMapWithSize(rv.Type(), rv.Len())
		iter := rv.MapRange()
		for iter.Next() {
			out.SetMapIndex(deepCopy(iter.Key()), deepCopy(iter.Value()))
		}
		return out
	case reflect.Struct:
		out := reflect.New(rv.Type()).Elem()
		for i := 0; i < rv.NumField(); i++ {
			if rv.Type().Field(i).IsExported() {
				out.Field(i).Set(deepCopy(rv.Field(i)))
			}
		}
		return out
	default:
		return rv
	}
}""",
    solution="""rv := reflect.ValueOf(v)
if !rv.IsValid() {
	return nil
}
return deepCopy(rv).Interface()""",
    tests="""
import (
	"reflect"
	"testing"
)

type node struct {
	Name  string
	Tags  []string
	Meta  map[string]int
	Child *node
	Fixed [2]int
}

func TestDeepCopyEqual(t *testing.T) {
	in := node{
		Name:  "root",
		Tags:  []string{"a", "b"},
		Meta:  map[string]int{"k": 1},
		Child: &node{Name: "kid", Tags: []string{"c"}},
		Fixed: [2]int{7, 8},
	}
	out, ok := DeepCopy(in).(node)
	if !ok {
		t.Fatalf("DeepCopy returned %T, want node", DeepCopy(in))
	}
	if !reflect.DeepEqual(in, out) {
		t.Fatalf("DeepCopy = %+v, want %+v", out, in)
	}
}

func TestDeepCopySharesNothing(t *testing.T) {
	in := node{
		Tags:  []string{"a"},
		Meta:  map[string]int{"k": 1},
		Child: &node{Name: "kid"},
	}
	out := DeepCopy(in).(node)

	out.Tags[0] = "changed"
	if in.Tags[0] != "a" {
		t.Error("the slice is shared")
	}
	out.Meta["k"] = 99
	if in.Meta["k"] != 1 {
		t.Error("the map is shared")
	}
	out.Child.Name = "changed"
	if in.Child.Name != "kid" {
		t.Error("the pointed-at struct is shared")
	}
	if out.Child == in.Child {
		t.Error("the pointer was copied, not what it points at")
	}
}

func TestDeepCopyNils(t *testing.T) {
	in := node{Name: "bare"}
	out := DeepCopy(in).(node)
	if out.Tags != nil || out.Meta != nil || out.Child != nil {
		t.Errorf("out = %+v, want the nil fields preserved", out)
	}
}

func TestDeepCopyScalarsAndNil(t *testing.T) {
	if got := DeepCopy(7); got != 7 {
		t.Errorf("DeepCopy(7) = %v, want 7", got)
	}
	if got := DeepCopy("s"); got != "s" {
		t.Errorf("DeepCopy = %v, want s", got)
	}
	if got := DeepCopy(nil); got != nil {
		t.Errorf("DeepCopy(nil) = %v, want nil", got)
	}
}

func TestDeepCopyNestedSlices(t *testing.T) {
	in := [][]int{{1, 2}, {3}}
	out := DeepCopy(in).([][]int)
	out[0][0] = 99
	if in[0][0] != 1 {
		t.Error("the inner slice is shared")
	}
}
""",
    context="A cache hands out the struct it stores. Callers mutate a slice three levels down, the cached entry changes with it, and the bug is reported as \"the cache returns wrong data at random\".",
    task=[
        "Return a copy of `v` that shares no mutable storage with it.",
        "Recurse through structs, pointers, interfaces, slices, arrays and maps.",
        "Preserve nil slices, maps and pointers as nil; skip unexported fields.",
        "Scalars, strings and a nil interface are returned as they are.",
    ],
    examples=[
        ("out := DeepCopy(in).(node); out.Tags[0] = \"x\"", "in.Tags is unchanged", None),
        ("out.Child == in.Child", "false", "The pointed-at struct is copied, not the pointer."),
        ("DeepCopy(node{Name:\"bare\"})", "nil fields stay nil", None),
    ],
    topics=[
        ("reflect.MakeSlice / MakeMapWithSize", "Building new containers of a type known only at run time."),
        ("Nil is not empty", "Copying a nil map into a made map changes observable behaviour."),
        ("Recursion over kinds", "Each container kind needs its own construction step."),
        ("Unexported fields are uncopyable", "`Set` refuses them, so they must be skipped."),
    ],
    hint="The recursive helper is written for you. The exported function validates, recurses and boxes.",
    intuition="A shallow copy duplicates the headers and shares the storage. A deep copy has to rebuild every container it meets, which means one construction rule per container kind — and a decision about nil, which is a distinct state from empty.",
    approach=[
        "Return nil for an invalid Value.",
        "Delegate to the recursive helper.",
        "Box the result back with `Interface()`.",
    ],
    walkthrough="Copying `node{Tags:[\"a\"], Child:&node{...}}` makes a new two-element `[]string`, a new `node` for the child, and a new pointer to it. Writing through any of them cannot reach the original.",
    pitfalls=[
        "Turning a nil slice into an empty one, which breaks `== nil` checks downstream.",
        "Copying map values without copying keys — keys can be structs containing pointers too.",
        "Forgetting that unexported fields cannot be set, so a copy of such a struct is necessarily partial.",
    ],
)

P(
    "staff",
    name="methodcache",
    title="Call A Method By Name, Once Resolved",
    sig="func CallNamed(v any, method string, arg int) (int, error)",
    doc="""CallNamed calls the named method on v with one int argument and
returns its single int result.

Method lookup by name is a search; the resolved index is cached per
(type, name) so repeated calls cost a map lookup.

Examples:

	CallNamed(adder{2}, "Add", 3) => 5, nil""",
    imports=['"errors"', '"reflect"', '"sync"'],
    extra="""// ErrMethod is returned when the method is missing or has the wrong shape.
var ErrMethod = errors.New("no such method with signature func(int) int")

type methodKey struct {
	t    reflect.Type
	name string
}

// methods caches the resolved method index, or -1 when the lookup failed.
var methods sync.Map // methodKey -> int

// methodIndex resolves the method's index on t, caching the answer.
func methodIndex(t reflect.Type, name string) int {
	k := methodKey{t, name}
	if v, ok := methods.Load(k); ok {
		return v.(int)
	}
	idx := -1
	if m, ok := t.MethodByName(name); ok {
		mt := m.Type
		// mt includes the receiver as parameter 0.
		if mt.NumIn() == 2 && mt.In(1).Kind() == reflect.Int &&
			mt.NumOut() == 1 && mt.Out(0).Kind() == reflect.Int {
			idx = m.Index
		}
	}
	actual, _ := methods.LoadOrStore(k, idx)
	return actual.(int)
}""",
    solution="""rv := reflect.ValueOf(v)
if !rv.IsValid() {
	return 0, ErrMethod
}
idx := methodIndex(rv.Type(), method)
if idx < 0 {
	return 0, ErrMethod
}
out := rv.Method(idx).Call([]reflect.Value{reflect.ValueOf(arg)})
return int(out[0].Int()), nil""",
    tests="""
import (
	"errors"
	"sync"
	"testing"
)

type adder struct{ Base int }

func (a adder) Add(n int) int      { return a.Base + n }
func (a adder) Twice(n int) int    { return 2 * n }
func (a adder) Name() string       { return "adder" }
func (a adder) Pair(n int) (int, int) { return n, n }

type ptrAdder struct{ Base int }

func (p *ptrAdder) Add(n int) int { return p.Base + n }

func TestCallNamed(t *testing.T) {
	if got, err := CallNamed(adder{Base: 2}, "Add", 3); err != nil || got != 5 {
		t.Errorf("CallNamed = %d, %v, want 5, nil", got, err)
	}
	if got, err := CallNamed(adder{}, "Twice", 4); err != nil || got != 8 {
		t.Errorf("CallNamed = %d, %v, want 8, nil", got, err)
	}
}

func TestCallNamedPointerReceiver(t *testing.T) {
	if got, err := CallNamed(&ptrAdder{Base: 10}, "Add", 5); err != nil || got != 15 {
		t.Errorf("CallNamed = %d, %v, want 15, nil", got, err)
	}
	if _, err := CallNamed(ptrAdder{Base: 10}, "Add", 5); !errors.Is(err, ErrMethod) {
		t.Error("a pointer-receiver method must not be reachable on the value")
	}
}

func TestCallNamedBadShapes(t *testing.T) {
	cases := []struct {
		name   string
		v      any
		method string
	}{
		{"missing", adder{}, "Nope"},
		{"wrong results", adder{}, "Name"},
		{"too many results", adder{}, "Pair"},
		{"nil value", nil, "Add"},
		{"no methods", 3, "Add"},
	}
	for _, c := range cases {
		if _, err := CallNamed(c.v, c.method, 1); !errors.Is(err, ErrMethod) {
			t.Errorf("%s: err = %v, want ErrMethod", c.name, err)
		}
	}
}

func TestCallNamedConcurrent(t *testing.T) {
	var wg sync.WaitGroup
	const workers = 16
	wg.Add(workers)
	errs := make([]error, workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			for i := 0; i < 200; i++ {
				got, err := CallNamed(adder{Base: w}, "Add", i)
				if err != nil {
					errs[w] = err
					return
				}
				if got != w+i {
					errs[w] = errors.New("wrong result")
					return
				}
			}
		}(w)
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestCallNamedRepeatsUseTheCache(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if got, err := CallNamed(adder{Base: 1}, "Add", i); err != nil || got != 1+i {
			t.Fatalf("iteration %d: %d, %v", i, got, err)
		}
	}
}
""",
    context="A command dispatcher resolves handler methods by name on every request. `MethodByName` is a linear scan of the method table, and the table is not small.",
    task=[
        "Call the named method on `v` with `arg` and return its int result.",
        "Use the shared cache so a (type, name) pair is resolved at most once, as far as callers can tell.",
        "Return `ErrMethod` for a missing method, a wrong signature, or a nil value.",
        "A pointer-receiver method must not be reachable through the value type.",
        "Correct under concurrent use.",
    ],
    examples=[
        ('CallNamed(adder{2}, "Add", 3)', "5, nil", None),
        ('CallNamed(adder{}, "Name", 1)', "ErrMethod", "The result is a string."),
        ('CallNamed(ptrAdder{}, "Add", 1)', "ErrMethod", "Pointer-receiver methods are not in the value's method set."),
    ],
    topics=[
        ("Method sets", "`T`'s method set excludes pointer-receiver methods; `*T`'s includes both."),
        ("Method.Type includes the receiver", "Parameter 0 is the receiver, so a one-argument method has `NumIn() == 2`."),
        ("Value.Method(i).Call", "The bound method value takes only the declared arguments."),
        ("Negative caching", "Storing -1 for a failed lookup keeps the failure path as cheap as the success path."),
    ],
    hint="The cache and the signature check are given. Resolve, guard, call, read the result.",
    intuition="Method resolution is a property of the type, so it is cacheable exactly like field layout. The subtlety is the receiver: it occupies parameter 0 of the method's type but is already bound by the time you call through `Value.Method`.",
    approach=[
        "Reject an invalid Value.",
        "Resolve the index through the cache; a negative index means `ErrMethod`.",
        "`rv.Method(idx).Call` with the single argument and read `out[0].Int()`.",
    ],
    walkthrough="`adder.Add` has method type `func(adder, int) int`, so `NumIn()` is 2 and `In(1)` is int — it passes. `ptrAdder.Add` is not in `ptrAdder`'s method set at all, so `MethodByName` fails and -1 is cached.",
    pitfalls=[
        "Checking `NumIn() == 1` and rejecting every valid method — the receiver is counted.",
        "Using `rv.Type().Method(i).Func` and forgetting to pass the receiver as the first argument.",
        "Caching without the type in the key, so two types with the same method name collide.",
    ],
)

P(
    "staff",
    name="fielddiff",
    title="Report Which Fields Differ",
    sig="func Diff(a, b any) []string",
    doc="""Diff returns the dotted paths of the exported fields where a and b
differ, in declaration order.

a and b must have the same type; otherwise the result is nil. Nested
structs contribute dotted paths.

Examples:

	Diff(cfg{A: 1}, cfg{A: 2}) => []string{"A"}""",
    imports=['"reflect"'],
    extra="""func diff(av, bv reflect.Value, prefix string, out *[]string) {
	if av.Kind() == reflect.Struct {
		rt := av.Type()
		for i := 0; i < av.NumField(); i++ {
			f := rt.Field(i)
			if !f.IsExported() {
				continue
			}
			name := f.Name
			if prefix != "" {
				name = prefix + "." + name
			}
			diff(av.Field(i), bv.Field(i), name, out)
		}
		return
	}
	if !av.Equal(bv) {
		*out = append(*out, prefix)
	}
}""",
    solution="""av, bv := reflect.ValueOf(a), reflect.ValueOf(b)
if !av.IsValid() || !bv.IsValid() || av.Type() != bv.Type() {
	return nil
}
var out []string
diff(av, bv, "", &out)
return out""",
    tests="""
import (
	"reflect"
	"testing"
)

type limits struct {
	Soft int
	Hard int
}

type settings struct {
	Name   string
	Retry  int
	Limits limits
	hidden int
}

func TestDiffFlat(t *testing.T) {
	a := settings{Name: "x", Retry: 1}
	b := settings{Name: "y", Retry: 1}
	if got := Diff(a, b); !reflect.DeepEqual(got, []string{"Name"}) {
		t.Errorf("Diff = %v, want [Name]", got)
	}
}

func TestDiffNested(t *testing.T) {
	a := settings{Limits: limits{Soft: 1, Hard: 2}}
	b := settings{Limits: limits{Soft: 9, Hard: 2}}
	if got := Diff(a, b); !reflect.DeepEqual(got, []string{"Limits.Soft"}) {
		t.Errorf("Diff = %v, want [Limits.Soft]", got)
	}
}

func TestDiffMultipleInDeclarationOrder(t *testing.T) {
	a := settings{Name: "x", Retry: 1, Limits: limits{Soft: 1}}
	b := settings{Name: "y", Retry: 2, Limits: limits{Soft: 2}}
	want := []string{"Name", "Retry", "Limits.Soft"}
	if got := Diff(a, b); !reflect.DeepEqual(got, want) {
		t.Errorf("Diff = %v, want %v", got, want)
	}
}

func TestDiffIdentical(t *testing.T) {
	a := settings{Name: "x", Retry: 1, Limits: limits{Soft: 1, Hard: 2}}
	if got := Diff(a, a); len(got) != 0 {
		t.Errorf("Diff = %v, want empty", got)
	}
}

func TestDiffIgnoresUnexported(t *testing.T) {
	a := settings{hidden: 1}
	b := settings{hidden: 2}
	if got := Diff(a, b); len(got) != 0 {
		t.Errorf("Diff = %v, want empty: unexported fields are not compared", got)
	}
}

func TestDiffMismatchedTypes(t *testing.T) {
	if got := Diff(settings{}, limits{}); got != nil {
		t.Errorf("Diff = %v, want nil", got)
	}
	if got := Diff(nil, settings{}); got != nil {
		t.Errorf("Diff = %v, want nil", got)
	}
	if got := Diff(nil, nil); got != nil {
		t.Errorf("Diff = %v, want nil", got)
	}
}

func TestDiffScalars(t *testing.T) {
	if got := Diff(1, 2); !reflect.DeepEqual(got, []string{""}) {
		t.Errorf("Diff = %v, want [\\"\\"]", got)
	}
	if got := Diff(1, 1); len(got) != 0 {
		t.Errorf("Diff = %v, want empty", got)
	}
}
""",
    context="A config reload logs \"settings changed\" and nothing else. On-call needs to know which setting, and the struct has forty fields across four nested blocks.",
    task=[
        "Return the dotted paths of the exported fields where `a` and `b` differ, in declaration order.",
        "Descend into nested structs, joining names with `.`.",
        "Skip unexported fields.",
        "Return nil when the types differ or either value is a nil interface.",
    ],
    examples=[
        ('Diff(settings{Name:"x"}, settings{Name:"y"})', "[Name]", None),
        ("Diff(a, b) with only Limits.Soft differing", "[Limits.Soft]", "Nested paths are dotted."),
        ("Diff(settings{}, limits{})", "<nil>", "Different types are not comparable."),
    ],
    topics=[
        ("Parallel traversal", "Two Values of one type have identical field indices, so they walk in lockstep."),
        ("Value.Equal", "Compares two Values of the same type without boxing them."),
        ("Path accumulation", "The prefix is built on the way down, so leaves know their full name."),
        ("Type identity", "`av.Type() != bv.Type()` is the only sound precondition for a field-by-field walk."),
    ],
    hint="The recursive helper is given. Validate the pair, then let it walk.",
    intuition="Two values of the same type have the same shape, so the walk visits identical positions in both. Only the leaves need comparing, and the path to each leaf is whatever names you passed on the way down.",
    approach=[
        "Reject invalid Values and mismatched types.",
        "Call the helper with an empty prefix and a pointer to the result slice.",
        "Return the accumulated paths.",
    ],
    walkthrough="For `settings`, the walk visits Name, Retry, then descends into Limits and visits Limits.Soft and Limits.Hard. `hidden` is skipped. Only the differing leaves are appended, in that order.",
    pitfalls=[
        "Comparing structs at the top level with `Equal` — you learn that they differ, not where.",
        "Using `reflect.DeepEqual` on the leaves, which boxes both sides on every comparison.",
        "Descending into unexported fields, which `Equal` refuses to compare.",
    ],
)

P(
    "staff",
    name="typeregistry",
    title="Build A Value From A Name",
    sig="func New(name string) (any, error)",
    doc="""New returns a freshly allocated pointer to the type registered under
name.

The registry is written once at init and read by many goroutines, so the
lookup must be safe without serialising every construction.

Examples:

	New("job") => *job, nil""",
    imports=['"errors"', '"reflect"', '"sync"'],
    extra="""// ErrUnknown is returned when no type is registered under the name.
var ErrUnknown = errors.New("unknown type name")

// registry maps a name to the registered type.
var registry sync.Map // string -> reflect.Type

// Register records a prototype under name. It is safe for concurrent use.
func Register(name string, prototype any) {
	registry.Store(name, reflect.TypeOf(prototype))
}

// lookup returns the registered type, or nil.
func lookup(name string) reflect.Type {
	v, ok := registry.Load(name)
	if !ok {
		return nil
	}
	t, _ := v.(reflect.Type)
	return t
}""",
    solution="""t := lookup(name)
if t == nil {
	return nil, ErrUnknown
}
return reflect.New(t).Interface(), nil""",
    tests="""
import (
	"errors"
	"sync"
	"testing"
)

type job struct {
	ID   int
	Name string
}

type task struct{ N int }

func init() {
	Register("job", job{})
	Register("task", task{})
	Register("int", 0)
}

func TestNewReturnsAPointerToTheRegisteredType(t *testing.T) {
	v, err := New("job")
	if err != nil {
		t.Fatal(err)
	}
	p, ok := v.(*job)
	if !ok {
		t.Fatalf("New = %T, want *job", v)
	}
	if *p != (job{}) {
		t.Errorf("New = %+v, want the zero job", *p)
	}
	p.ID = 7
}

func TestNewGivesDistinctValues(t *testing.T) {
	a, _ := New("job")
	b, _ := New("job")
	if a.(*job) == b.(*job) {
		t.Error("New returned the same pointer twice")
	}
	a.(*job).ID = 1
	if b.(*job).ID != 0 {
		t.Error("the two values share storage")
	}
}

func TestNewOtherTypes(t *testing.T) {
	v, err := New("task")
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := v.(*task); !ok {
		t.Errorf("New = %T, want *task", v)
	}
	v, err = New("int")
	if err != nil {
		t.Fatal(err)
	}
	if p, ok := v.(*int); !ok || *p != 0 {
		t.Errorf("New = %T, want *int holding 0", v)
	}
}

func TestNewUnknown(t *testing.T) {
	if _, err := New("nope"); !errors.Is(err, ErrUnknown) {
		t.Errorf("err = %v, want ErrUnknown", err)
	}
	if _, err := New(""); !errors.Is(err, ErrUnknown) {
		t.Errorf("err = %v, want ErrUnknown", err)
	}
}

func TestNewConcurrent(t *testing.T) {
	const workers = 16
	var wg sync.WaitGroup
	wg.Add(workers)
	errs := make([]error, workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			for i := 0; i < 200; i++ {
				v, err := New("job")
				if err != nil {
					errs[w] = err
					return
				}
				p := v.(*job)
				p.ID = w
				if p.ID != w {
					errs[w] = errors.New("value is shared between goroutines")
					return
				}
			}
		}(w)
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			t.Fatal(err)
		}
	}
}
""",
    context="A message broker decodes payloads into the struct named in the envelope. The dispatch is a switch over forty type names that has to be edited every time a message type is added.",
    task=[
        "Return a pointer to a freshly allocated zero value of the type registered under `name`.",
        "Every call returns a distinct value.",
        "Return `ErrUnknown` when nothing is registered under that name.",
        "Correct under concurrent use — many goroutines constructing at once.",
    ],
    examples=[
        ('New("job")', "*job pointing at the zero job, nil", None),
        ('a, _ := New("job"); b, _ := New("job")', "a != b", "Each call allocates its own value."),
        ('New("nope")', "ErrUnknown", None),
    ],
    topics=[
        ("reflect.New", "Allocates a zero value of a run-time type and returns a pointer Value to it."),
        ("Value.Interface", "Boxes the pointer back with its real dynamic type, so `v.(*job)` succeeds."),
        ("Types are immutable and shareable", "A `reflect.Type` is safe to hold and read from any goroutine."),
        ("sync.Map for a read-mostly registry", "Written at init, read on every message."),
    ],
    hint="`Register` and `lookup` are given. Three lines: look up, guard, construct.",
    intuition="A type value is a run-time handle to everything the compiler knew about a type — enough to allocate one. That turns a forty-case switch into a map from name to type plus a single `New`.",
    approach=[
        "`lookup(name)`; return `ErrUnknown` on a nil type.",
        "`reflect.New(t)` for a pointer to a fresh zero value.",
        "Return `Interface()` so the caller can type-assert it.",
    ],
    walkthrough='`New("job")` finds `reflect.TypeOf(job{})`, allocates a zero `job`, and boxes the `*job` pointer. The caller\'s `v.(*job)` succeeds because the dynamic type survived the round trip.',
    pitfalls=[
        "Returning `reflect.New(t).Elem().Interface()`, which boxes a copy and gives the caller nothing to write through.",
        "Caching a constructed value instead of the type, which would hand every caller the same struct.",
        "Registering a pointer prototype, which makes `New` return `**T`.",
    ],
)
