package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMdInline(t *testing.T) {
	cases := map[string]string{
		"plain":                   "plain",
		"a `code` b":              "a <code>code</code> b",
		"a **bold** b":            "a <strong>bold</strong> b",
		"an *emphasis* here":      "an <em>emphasis</em> here",
		"**bold** and *italic*":   "<strong>bold</strong> and <em>italic</em>",
		"see [docs](http://x)":    "see docs",
		"5 < 6 & 7 > 2":           "5 &lt; 6 &amp; 7 &gt; 2",
		"`a<b>` stays escaped":    "<code>a&lt;b&gt;</code> stays escaped",
		"**bold** and `code` mix": "<strong>bold</strong> and <code>code</code> mix",
	}
	for in, want := range cases {
		if got := mdInline(in); got != want {
			t.Errorf("mdInline(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestMdToHTML(t *testing.T) {
	md := strings.Join([]string{
		"## Task",
		"",
		"Do the thing.",
		"",
		"- one",
		"- two",
		"",
		"```go",
		"if a < b {}",
		"```",
		"",
		"<details><summary>Hint</summary>",
	}, "\n")
	got := mdToHTML(md)
	for _, want := range []string{
		"<h2>Task</h2>",
		"<p>Do the thing.</p>",
		"<ul>\n<li>one</li>\n<li>two</li>\n</ul>",
		`<pre class="md"><code>`,
		"if a &lt; b {}", // code is escaped, not interpreted
		"</code></pre>",
		"<details><summary>Hint</summary>", // raw HTML passes through
	} {
		if !strings.Contains(got, want) {
			t.Errorf("output missing %q\n---\n%s", want, got)
		}
	}
}

func TestFieldAndTitle(t *testing.T) {
	md := "# Challenge 03 — Swap\n\n**Level:** Junior\n**Estimated time:** 10 min\n\nBody.\n"
	if got := field(md, "Level"); got != "Junior" {
		t.Errorf("field(Level) = %q", got)
	}
	if got := field(md, "Estimated time"); got != "10 min" {
		t.Errorf("field(Estimated time) = %q", got)
	}
	if got := field(md, "Nope"); got != "" {
		t.Errorf("field(Nope) = %q, want empty", got)
	}
	if got := titleOf(md, "junior/01-a/01-b/swap"); got != "Challenge 03 — Swap" {
		t.Errorf("titleOf = %q", got)
	}
	if got := titleOf("no heading here", "junior/01-a/01-b/swap"); got != "swap" {
		t.Errorf("titleOf fallback = %q, want the dir name", got)
	}
}

func TestStripHeadFields(t *testing.T) {
	md := "# Title\n\n**Level:** Junior\n**Topic:** Variables\n\nBody line.\n"
	got := stripHeadFields(md)
	if strings.Contains(got, "# Title") || strings.Contains(got, "**Level:**") || strings.Contains(got, "**Topic:**") {
		t.Errorf("meta not stripped:\n%s", got)
	}
	if !strings.Contains(got, "Body line.") {
		t.Errorf("body lost:\n%s", got)
	}
}

func TestTopicsSectionAndTags(t *testing.T) {
	md := strings.Join([]string{
		"# T",
		"",
		"## Task",
		"Body.",
		"",
		"## Topics",
		"| Concept | Example |",
		"| **iota** | `_ = iota` |",
		"| **const** | `const x = 1` |",
		"",
		"## Hint",
		"| **notatopic** | x |",
	}, "\n")

	tags := topicTags(md)
	want := []string{"iota", "const"}
	if len(tags) != len(want) {
		t.Fatalf("topicTags = %q, want %q", tags, want)
	}
	for i := range want {
		if tags[i] != want[i] {
			t.Errorf("topicTags[%d] = %q, want %q", i, tags[i], want[i])
		}
	}
	// The example code must not leak in as a chip, and the section must stop at
	// the next H2.
	for _, tag := range tags {
		if strings.Contains(tag, "=") || tag == "notatopic" {
			t.Errorf("unexpected chip %q", tag)
		}
	}

	if got := stripTopicsSection(md); strings.Contains(got, "## Topics") {
		t.Errorf("Topics section not stripped:\n%s", got)
	}
	if !strings.Contains(topicsFooterHTML(md), "<span class=\"topics-n\">2</span>") {
		t.Errorf("footer count wrong: %s", topicsFooterHTML(md))
	}
	if topicsFooterHTML("# T\n\nno topics\n") != "" {
		t.Error("footer should be empty when there are no topics")
	}
}

func TestTopicTagsFallsBackToMetaLine(t *testing.T) {
	md := "# T\n\n**Topic:** Slices → Aliasing (`append`, `copy`)\n"
	got := topicTags(md)
	want := []string{"Slices", "Aliasing", "append", "copy"}
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Errorf("topicTags = %q, want %q", got, want)
	}
}

func TestCleanDirNameAndTitle(t *testing.T) {
	cases := map[string]string{
		"03-composite-types":         "Composite Types",
		"01-variables-and-constants": "Variables And Constants",
		"dedupe":                     "Dedupe",
	}
	for in, want := range cases {
		if got := cleanDirName(in); got != want {
			t.Errorf("cleanDirName(%q) = %q, want %q", in, got, want)
		}
	}
	if got := cleanTitle("Challenge 07 — Swap Two", "junior/01-a/01-b/swap"); got != "Swap Two" {
		t.Errorf("cleanTitle = %q", got)
	}
	if got := cleanTitle("Challenge M01 — Pool", "junior/01-a/01-b/pool"); got != "Pool" {
		t.Errorf("cleanTitle = %q", got)
	}
	if got := cleanTitle("Plain Title", "junior/01-a/01-b/x"); got != "Plain Title" {
		t.Errorf("cleanTitle = %q", got)
	}
	if got := groupKey("junior/01-language-basics/02-data-types/swap"); got != "Junior · Language Basics" {
		t.Errorf("groupKey = %q", got)
	}
}

// Python's json.dumps escapes non-ASCII; the generated file is compared against
// that historical output, so the escaping must match.
func TestEscapeNonASCII(t *testing.T) {
	cases := map[string]string{
		`"a"`: `"a"`,
		`"→"`: `"\u2192"`,
		`"·"`: `"\u00b7"`,
		`"🎉"`: `"\ud83c\udf89"`, // outside the BMP: surrogate pair
	}
	for in, want := range cases {
		if got := string(escapeNonASCII([]byte(in))); got != want {
			t.Errorf("escapeNonASCII(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestMarshalKeepsHTMLUnescaped(t *testing.T) {
	b, err := marshal(map[string]string{"d": "<p>a & b</p>"})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(b), "<p>a & b</p>") {
		t.Errorf("HTML was escaped: %s", b)
	}
}

func TestBuildEndToEnd(t *testing.T) {
	root := t.TempDir()
	write := func(dir, name, body string) {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	pj := filepath.Join(root, "challenges", "junior", "01-basics", "01-vars", "swap")
	write(pj, "go.mod", "module x\n")
	write(pj, "swap.go", "package swap\n")
	write(pj, "swap_test.go", "package swap\n")
	write(pj, "swap.debug.txt", "buggy\n")
	write(pj, "README.md", "# Challenge 01 — Swap\n\n**Level:** Junior\n\nBody.\n\n## Topics\n| **iota** |\n")

	ps := filepath.Join(root, "challenges", "staff", "01-conc", "01-race", "pool")
	write(ps, "go.mod", "module y\n")
	write(ps, "pool.go", "package pool\n")

	problems, groups, err := build(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(problems) != 2 {
		t.Fatalf("problems = %d, want 2", len(problems))
	}

	swap := problems["junior/01-basics/01-vars/swap"]
	if swap.Title != "Challenge 01 — Swap" || swap.Level != "junior" {
		t.Errorf("swap = %+v", swap)
	}
	if swap.Tag != "backend" {
		t.Errorf("tag = %q; a puzzle with tests runs on the backend", swap.Tag)
	}
	if swap.Starter != "package swap\n" || swap.Debug != "buggy\n" {
		t.Errorf("starter/debug = %q/%q", swap.Starter, swap.Debug)
	}
	if !strings.Contains(swap.Description, "<h1 class=\"p-title\">") {
		t.Errorf("description missing header: %s", swap.Description)
	}

	// No test file → not a backend puzzle.
	if got := problems["staff/01-conc/01-race/pool"].Tag; got != "" {
		t.Errorf("pool tag = %q, want empty", got)
	}

	// Groups are ordered by level, junior before staff.
	if len(groups) != 2 {
		t.Fatalf("groups = %d, want 2", len(groups))
	}
	if groups[0].Topic != "Junior · Basics" || groups[1].Topic != "Staff · Conc" {
		t.Errorf("group order = %q, %q", groups[0].Topic, groups[1].Topic)
	}
	it := groups[0].Items[0]
	if it.Lv != "J" || it.Sub != "Vars" || it.Name != "Swap" || it.Locked {
		t.Errorf("item = %+v", it)
	}

	// The emitted JSON must be valid and round-trip.
	b, err := marshal(problems)
	if err != nil {
		t.Fatal(err)
	}
	var back map[string]problem
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatalf("emitted JSON does not parse: %v", err)
	}
	if back["junior/01-basics/01-vars/swap"].Title != swap.Title {
		t.Error("round-trip lost the title")
	}
}

func TestBuildSkipsDirsWithoutGoMod(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "challenges", "junior", "01-a", "01-b", "empty"), 0o755); err != nil {
		t.Fatal(err)
	}
	problems, groups, err := build(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(problems) != 0 || len(groups) != 0 {
		t.Errorf("problems=%d groups=%d, want 0/0", len(problems), len(groups))
	}
}

// An unterminated ``` fence must still close the <pre> block.
func TestMdToHTMLUnclosedCodeFence(t *testing.T) {
	got := mdToHTML("```go\nx := 1\n")
	if !strings.HasSuffix(got, "</code></pre>") {
		t.Errorf("output = %q, want a closed code block", got)
	}
}

// Older list-style Topics sections bold nothing; the backticked items are the
// chips instead.
func TestTopicsFromSectionBacktickFallback(t *testing.T) {
	md := "# T\n\n## Topics\n- `iota`\n- `const`\n"
	got := topicsFromSection(md)
	if strings.Join(got, ",") != "iota,const" {
		t.Errorf("topicsFromSection = %q, want [iota const]", got)
	}
	// A Topics section with neither is empty, not a crash.
	if got := topicsFromSection("# T\n\n## Topics\nplain prose\n"); len(got) != 0 {
		t.Errorf("topicsFromSection = %q, want empty", got)
	}
}

func TestHeaderHTMLIncludesEstimatedTime(t *testing.T) {
	got := headerHTML("**Level:** Junior\n**Estimated time:** 15 min\n", "Swap")
	if !strings.Contains(got, `<span class="p-time">15 min</span>`) {
		t.Errorf("header = %q, want the time pill", got)
	}
	if !strings.Contains(got, `lv-junior`) {
		t.Errorf("header = %q, want the level pill", got)
	}
	// Neither field: just the title.
	bare := headerHTML("no meta here", "Swap")
	if strings.Contains(bare, "pill") || strings.Contains(bare, "p-time") {
		t.Errorf("header = %q, want no meta", bare)
	}
}

// "Challenge 07 — " with nothing after it falls back to the directory name.
func TestCleanTitleFallsBackToDirName(t *testing.T) {
	if got := cleanTitle("Challenge 07 — ", "junior/01-a/01-b/swap-two"); got != "Swap Two" {
		t.Errorf("cleanTitle = %q, want Swap Two", got)
	}
}

func TestCapitalizeEmpty(t *testing.T) {
	if got := capitalize(""); got != "" {
		t.Errorf("capitalize(\"\") = %q", got)
	}
	if got := capitalize("jUNIOR"); got != "Junior" {
		t.Errorf("capitalize = %q", got)
	}
}

// EDUCATION.md is optional: no file (or a blank one) means no education tab.
func TestEducationHTML(t *testing.T) {
	if got := educationHTML(""); got != "" {
		t.Errorf("educationHTML(\"\") = %q, want empty", got)
	}
	if got := educationHTML("\n\n  \n"); got != "" {
		t.Errorf("educationHTML(whitespace) = %q, want empty", got)
	}
	got := educationHTML("# Scope\n\nA variable lives in its block.\n")
	if !strings.Contains(got, "<h1>Scope</h1>") || !strings.Contains(got, "<p>A variable lives in its block.</p>") {
		t.Errorf("educationHTML = %q", got)
	}
}

// A puzzle shipping EDUCATION.md carries it into the catalog; one without it
// gets an empty field.
func TestBuildPicksUpEducation(t *testing.T) {
	root := t.TempDir()
	with := filepath.Join(root, "challenges", "junior", "01-a", "01-b", "taught")
	writeFile(t, with, "go.mod", "module taught\n")
	writeFile(t, with, "EDUCATION.md", "# Blank identifier\n\nIt discards a value.\n")
	without := filepath.Join(root, "challenges", "junior", "01-a", "01-b", "bare")
	writeFile(t, without, "go.mod", "module bare\n")

	problems, _, err := build(root)
	if err != nil {
		t.Fatal(err)
	}
	if got := problems["junior/01-a/01-b/taught"].Education; !strings.Contains(got, "Blank identifier") {
		t.Errorf("education = %q", got)
	}
	if got := problems["junior/01-a/01-b/bare"].Education; got != "" {
		t.Errorf("education = %q, want empty for a puzzle without EDUCATION.md", got)
	}
}
