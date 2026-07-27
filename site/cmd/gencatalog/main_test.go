package main

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestMain re-execs the binary as the real command so main() itself is covered.
func TestMain(m *testing.M) {
	if os.Getenv("GW_TEST_MAIN") != "" {
		os.Args = append([]string{"gencatalog"}, strings.Fields(os.Getenv("GW_TEST_MAIN_ARGS"))...)
		main()
		os.Exit(0)
	}
	os.Exit(m.Run())
}

func writeFile(t *testing.T, dir, name, body string) {
	t.Helper()
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

func repoFixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	p := filepath.Join(root, "challenges", "junior", "01-basics", "01-vars", "swap")
	writeFile(t, p, "go.mod", "module swap\n")
	writeFile(t, p, "swap.go", "package swap\n")
	writeFile(t, p, "README.md", "# Swap\n\n**Level:** Junior\n")
	return root
}

func TestRunWritesCatalog(t *testing.T) {
	root := repoFixture(t)
	out := filepath.Join(t.TempDir(), "problems.js")

	var stdout bytes.Buffer
	if err := run([]string{"-root", root, "-out", out}, &stdout); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(stdout.String(), "1 puzzles") {
		t.Errorf("stdout = %q", stdout.String())
	}
	b, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	body := string(b)
	for _, want := range []string{"AUTO-GENERATED", "window.PROBLEMS = {", "window.CATALOG = ["} {
		if !strings.Contains(body, want) {
			t.Errorf("output missing %q", want)
		}
	}
}

// With no -out the catalog lands at the conventional path inside the repo.
func TestRunDefaultOutPath(t *testing.T) {
	root := repoFixture(t)
	if err := os.MkdirAll(filepath.Join(root, "site", "web", "assets", "js"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := run([]string{"-root", root}, &bytes.Buffer{}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, "site", "web", "assets", "js", "problems.js")); err != nil {
		t.Errorf("default output not written: %v", err)
	}
}

// No -root: walk up from the working directory.
func TestRunFindsRootFromWorkingDir(t *testing.T) {
	root := repoFixture(t)
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(wd) })
	if err := os.Chdir(filepath.Join(root, "challenges", "junior")); err != nil {
		t.Fatal(err)
	}
	out := filepath.Join(t.TempDir(), "p.js")
	if err := run([]string{"-out", out}, &bytes.Buffer{}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(out); err != nil {
		t.Errorf("no output: %v", err)
	}
}

func TestRunErrors(t *testing.T) {
	root := repoFixture(t)

	if err := run([]string{"-nope"}, &bytes.Buffer{}); err == nil {
		t.Error("unknown flag: want an error")
	}
	// An unwritable output path.
	err := run([]string{"-root", root, "-out", filepath.Join(t.TempDir(), "missing", "p.js")}, &bytes.Buffer{})
	if err == nil {
		t.Error("unwritable out: want an error")
	}
	// No challenges/ above cwd and no -root.
	wd, _ := os.Getwd()
	t.Cleanup(func() { _ = os.Chdir(wd) })
	if err := os.Chdir(t.TempDir()); err != nil {
		t.Fatal(err)
	}
	if err := run(nil, &bytes.Buffer{}); err == nil {
		t.Log("a challenges/ dir exists above the temp dir; environment-dependent")
	}
}

// build must fail loudly if challenges/ cannot be walked at all.
func TestBuildMissingChallengesDir(t *testing.T) {
	if _, _, err := build(filepath.Join(t.TempDir(), "nope")); err == nil {
		t.Error("want an error when challenges/ is absent")
	}
}

func TestFindPuzzlesError(t *testing.T) {
	if _, err := findPuzzles(filepath.Join(t.TempDir(), "nope")); err == nil {
		t.Error("want an error walking a missing dir")
	}
}

func TestBuildUnreadablePuzzleDir(t *testing.T) {
	if os.Geteuid() == 0 {
		t.Skip("running as root: permission bits do not apply")
	}
	root := t.TempDir()
	p := filepath.Join(root, "challenges", "junior", "01-a", "01-b", "x")
	writeFile(t, p, "go.mod", "module x\n")
	if err := os.Chmod(p, 0o111); err != nil { // executable but not readable
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chmod(p, 0o755) })
	if _, _, err := build(root); err == nil {
		t.Error("want an error for an unreadable puzzle dir")
	}
}

// A go.mod directly under challenges/ is not a puzzle: the layout is
// level/topic/subtopic/name.
func TestBuildSkipsShallowModules(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "challenges", "stray"), "go.mod", "module stray\n")
	problems, groups, err := build(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(problems) != 0 || len(groups) != 0 {
		t.Errorf("problems=%d groups=%d, want 0/0", len(problems), len(groups))
	}
}

func TestBuildHonoursExcluded(t *testing.T) {
	root := repoFixture(t)
	slug := "junior/01-basics/01-vars/swap"
	excluded[slug] = true
	t.Cleanup(func() { delete(excluded, slug) })

	problems, _, err := build(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(problems) != 0 {
		t.Errorf("problems = %d, want the excluded puzzle skipped", len(problems))
	}
}

// An unknown level sorts after the known ones rather than crashing.
func TestLevelRankUnknown(t *testing.T) {
	if got := levelRank("wizard"); got != 9 {
		t.Errorf("levelRank(wizard) = %d, want 9", got)
	}
	if got := levelRank("junior"); got != 0 {
		t.Errorf("levelRank(junior) = %d", got)
	}
}

func TestFirstOrAndReadFile(t *testing.T) {
	if got := firstOr(nil, "fallback"); got != "fallback" {
		t.Errorf("firstOr(nil) = %q", got)
	}
	if got := firstOr([]string{"a", "b"}, "fallback"); got != "a" {
		t.Errorf("firstOr = %q", got)
	}
	if got := readFile(t.TempDir(), ""); got != "" {
		t.Errorf("readFile with no name = %q", got)
	}
	if got := readFile(t.TempDir(), "missing.go"); got != "" {
		t.Errorf("readFile of a missing file = %q", got)
	}
}

// A puzzle with no .go file at all still generates an entry, just with no
// starter code.
func TestBuildPuzzleWithoutGoFile(t *testing.T) {
	root := t.TempDir()
	p := filepath.Join(root, "challenges", "junior", "01-a", "01-b", "empty")
	writeFile(t, p, "go.mod", "module empty\n")
	problems, _, err := build(root)
	if err != nil {
		t.Fatal(err)
	}
	got := problems["junior/01-a/01-b/empty"]
	if got.Starter != "" || got.File != "" {
		t.Errorf("starter=%q file=%q, want both empty", got.Starter, got.File)
	}
}

// The starter is the shortest-named .go file; File reports the first by name.
func TestBuildStarterVsFile(t *testing.T) {
	root := t.TempDir()
	p := filepath.Join(root, "challenges", "junior", "01-a", "01-b", "x")
	writeFile(t, p, "go.mod", "module x\n")
	writeFile(t, p, "aaa_helpers.go", "package x // helpers\n")
	writeFile(t, p, "x.go", "package x // starter\n")
	problems, _, err := build(root)
	if err != nil {
		t.Fatal(err)
	}
	got := problems["junior/01-a/01-b/x"]
	if !strings.Contains(got.Starter, "starter") {
		t.Errorf("starter = %q, want the shortest-named file", got.Starter)
	}
	if got.File != "aaa_helpers.go" {
		t.Errorf("file = %q, want the first by name", got.File)
	}
}

func TestMarshalError(t *testing.T) {
	if _, err := marshal(map[string]any{"ch": make(chan int)}); err == nil {
		t.Error("marshalling a channel: want an error")
	}
}

func TestWriteErrors(t *testing.T) {
	// Unmarshalable problems.
	if err := write(filepath.Join(t.TempDir(), "p.js"), nil, nil); err != nil {
		t.Errorf("empty catalog: %v", err)
	}
	// Unwritable path.
	if err := write(filepath.Join(t.TempDir(), "missing", "p.js"), nil, nil); err == nil {
		t.Error("unwritable path: want an error")
	}
}

// main() must exit non-zero when generation fails.
func TestMainExitsOnFailure(t *testing.T) {
	out, err := runSelf(t, "-root "+filepath.Join(t.TempDir(), "no-such-root"))
	if err == nil {
		t.Fatalf("exit status 0; want failure. output:\n%s", out)
	}
}

func TestMainSucceeds(t *testing.T) {
	root := repoFixture(t)
	out, err := runSelf(t, "-root "+root+" -out "+filepath.Join(t.TempDir(), "p.js"))
	if err != nil {
		t.Fatalf("main failed: %v\n%s", err, out)
	}
	if !strings.Contains(out, "1 puzzles") {
		t.Errorf("output = %q", out)
	}
}

// runSelf re-execs this test binary as the real command.
func runSelf(t *testing.T, args string) (string, error) {
	t.Helper()
	cmd := exec.Command(os.Args[0], "-test.run=TestMainSucceeds")
	cmd.Env = append(os.Environ(), "GW_TEST_MAIN=1", "GW_TEST_MAIN_ARGS="+args)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

// The seams exist so these error paths stay exercised; provoke each one.
func TestSeamErrorPaths(t *testing.T) {
	boom := errors.New("boom")

	t.Run("getwd", func(t *testing.T) {
		old := getwd
		getwd = func() (string, error) { return "", boom }
		t.Cleanup(func() { getwd = old })
		if _, err := findRoot(""); !errors.Is(err, boom) {
			t.Errorf("findRoot = %v, want the getwd error", err)
		}
	})

	t.Run("readDir", func(t *testing.T) {
		old := readDir
		readDir = func(string) ([]os.DirEntry, error) { return nil, boom }
		t.Cleanup(func() { readDir = old })
		if _, _, err := build(repoFixture(t)); !errors.Is(err, boom) {
			t.Errorf("build = %v, want the readDir error", err)
		}
	})

	t.Run("marshal", func(t *testing.T) {
		old := marshalJSON
		marshalJSON = func(any) ([]byte, error) { return nil, boom }
		t.Cleanup(func() { marshalJSON = old })
		if err := write(filepath.Join(t.TempDir(), "p.js"), nil, nil); !errors.Is(err, boom) {
			t.Errorf("write = %v, want the marshal error", err)
		}
	})

	t.Run("marshal-catalog", func(t *testing.T) {
		old := marshalJSON
		calls := 0
		marshalJSON = func(v any) ([]byte, error) {
			calls++
			if calls == 1 {
				return old(v)
			}
			return nil, boom
		}
		t.Cleanup(func() { marshalJSON = old })
		if err := write(filepath.Join(t.TempDir(), "p.js"), nil, nil); !errors.Is(err, boom) {
			t.Errorf("write = %v, want the catalog marshal error", err)
		}
	})
}

// A relative -out against an absolute root has no relative path between them,
// so the printed name falls back to the raw output path.
func TestRunRelativeOutPath(t *testing.T) {
	root := repoFixture(t)
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(wd) })
	tmp := t.TempDir()
	if err := os.Chdir(tmp); err != nil {
		t.Fatal(err)
	}
	var stdout bytes.Buffer
	if err := run([]string{"-root", root, "-out", "out.js"}, &stdout); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(stdout.String(), "out.js") {
		t.Errorf("stdout = %q", stdout.String())
	}
}

// Two topics in the same level, and two puzzles in one topic: exercises both
// arms of the sidebar ordering.
func TestBuildOrdersWithinAndAcrossTopics(t *testing.T) {
	root := t.TempDir()
	for _, slug := range []string{
		"junior/02-second/01-a/beta",
		"junior/01-first/01-a/alpha",
		"junior/01-first/01-a/zeta",
	} {
		writeFile(t, filepath.Join(root, "challenges", filepath.FromSlash(slug)), "go.mod", "module m\n")
	}
	_, groups, err := build(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(groups) != 2 {
		t.Fatalf("groups = %d, want 2", len(groups))
	}
	if groups[0].Topic != "Junior · First" || groups[1].Topic != "Junior · Second" {
		t.Errorf("topic order = %q, %q", groups[0].Topic, groups[1].Topic)
	}
	// No README, so the title falls back to the dir name as-is.
	if len(groups[0].Items) != 2 || groups[0].Items[0].Name != "alpha" {
		t.Errorf("items = %+v, want alpha before zeta", groups[0].Items)
	}
}
