// Command gencatalog generates site/web/assets/js/problems.js from the
// challenges/ tree.
//
// It walks every puzzle (a dir containing go.mod), takes the starter code from
// its non-test .go file and the description from README.md, and emits
// window.PROBLEMS + window.CATALOG. Every puzzle runs against the real Go
// toolchain via the localhost backend (site/cmd/localrunner), so there are no
// per-puzzle in-browser runners to register.
//
//	go run ./site/cmd/gencatalog             # from the repo root
//	go run . -root /path/to/repo             # from this directory
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Test seams: these wrap calls whose failure modes cannot be provoked through
// the filesystem alone, so the error handling below stays exercised.
var (
	getwd       = os.Getwd
	readDir     = os.ReadDir
	marshalJSON = marshal
)

// levelOrder ranks the sidebar groups; unknown levels sink to the bottom.
var levelOrder = map[string]int{"junior": 0, "middle": 1, "senior": 2, "staff": 3}

// excluded puzzles stay on disk but are not published to the site.
var excluded = map[string]bool{}

type problem struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Level         string `json:"level"`
	Tag           string `json:"tag"`
	Description   string `json:"description"`
	Education     string `json:"education"`
	Starter       string `json:"starter"`
	Debug         string `json:"debug"`
	File          string `json:"file"`
	Run           string `json:"run"`
	RunCustom     string `json:"runCustom"`
	CustomDefault string `json:"customDefault"`
}

type catItem struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Sub    string `json:"sub"`
	Lv     string `json:"lv"`
	Level  string `json:"level"`
	Locked bool   `json:"locked"`
	Tag    string `json:"tag"`
}

type catGroup struct {
	Topic string    `json:"topic"`
	Items []catItem `json:"items"`
}

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		log.Fatal(err)
	}
}

// run is main's body, split out so tests can drive it without a process exit.
func run(args []string, stdout io.Writer) error {
	fs := flag.NewFlagSet("gencatalog", flag.ContinueOnError)
	fs.SetOutput(stdout)
	root := fs.String("root", "", "repo root (dir containing challenges/); auto-detected if empty")
	out := fs.String("out", "", "output path (default <root>/site/web/assets/js/problems.js)")
	if err := fs.Parse(args); err != nil {
		return err
	}

	repo, err := findRoot(*root)
	if err != nil {
		return err
	}
	outPath := *out
	if outPath == "" {
		outPath = filepath.Join(repo, "site", "web", "assets", "js", "problems.js")
	}

	problems, groups, err := build(repo)
	if err != nil {
		return err
	}
	if err := write(outPath, problems, groups); err != nil {
		return err
	}

	rel, err := filepath.Rel(repo, outPath)
	if err != nil {
		rel = outPath
	}
	fmt.Fprintf(stdout, "wrote %s: %d puzzles (all run on the local Go toolchain backend)\n", rel, len(problems))
	return nil
}

// findRoot honours an explicit root, else walks up from the working directory
// looking for challenges/.
func findRoot(explicit string) (string, error) {
	if explicit != "" {
		return filepath.Abs(explicit)
	}
	dir, err := getwd()
	if err != nil {
		return "", err
	}
	for {
		if fi, err := os.Stat(filepath.Join(dir, "challenges")); err == nil && fi.IsDir() {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("no challenges/ dir found above %s (pass -root)", dir)
		}
		dir = parent
	}
}

// findPuzzles lists every puzzle dir (one holding a go.mod), sorted.
func findPuzzles(challenges string) ([]string, error) {
	var dirs []string
	err := filepath.WalkDir(challenges, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && d.Name() == "go.mod" {
			dirs = append(dirs, filepath.Dir(path))
		}
		return nil
	})
	sort.Strings(dirs)
	return dirs, err
}

func build(repo string) (map[string]problem, []catGroup, error) {
	challenges := filepath.Join(repo, "challenges")
	dirs, err := findPuzzles(challenges)
	if err != nil {
		return nil, nil, err
	}

	problems := map[string]problem{}
	groups := map[string][]catItem{}

	for _, dir := range dirs {
		// findPuzzles only ever returns paths under challenges/, so the slug is
		// a plain prefix trim.
		slug := filepath.ToSlash(strings.TrimPrefix(dir, challenges+string(os.PathSeparator)))
		if excluded[slug] {
			continue
		}
		parts := strings.Split(slug, "/")
		if len(parts) < 3 {
			continue // not a <level>/<topic>/<subtopic>/<name> puzzle
		}

		entries, err := readDir(dir)
		if err != nil {
			return nil, nil, err
		}
		var goFiles, debugFiles []string
		hasTests := false
		for _, e := range entries {
			n := e.Name()
			switch {
			case strings.HasSuffix(n, "_test.go"):
				hasTests = true
			case strings.HasSuffix(n, ".go"):
				goFiles = append(goFiles, n)
			case strings.HasSuffix(n, ".debug.txt"):
				debugFiles = append(debugFiles, n)
			}
		}
		sort.Strings(goFiles)
		sort.Strings(debugFiles)

		// The stub the candidate edits: shortest-named non-test .go file.
		starterFile := ""
		if len(goFiles) > 0 {
			starterFile = goFiles[0]
			for _, f := range goFiles {
				if len(f) < len(starterFile) {
					starterFile = f
				}
			}
		}
		starter := readFile(dir, starterFile)

		// Optional planted-bug variant for Debug mode: a `*.debug.txt` sibling
		// that compiles but is wrong, exercised by the same tests.
		debug := ""
		if len(debugFiles) > 0 {
			debug = readFile(dir, debugFiles[0])
		}

		// Optional teaching material for the puzzle's concept, shown in its own
		// tab beside the description.
		education := readFile(dir, "EDUCATION.md")

		md := readFile(dir, "README.md")
		title := titleOf(md, slug)
		level := strings.ToLower(field(md, "Level"))
		if level == "" {
			level = strings.ToLower(parts[0])
		}
		// A puzzle with tests runs on the backend; that is all of them today.
		tag := ""
		if hasTests {
			tag = "backend"
		}

		problems[slug] = problem{
			ID:    slug,
			Title: title,
			Level: level,
			Tag:   tag,
			Description: markSolutionNocopy(
				headerHTML(md, title) +
					mdToHTML(stripHeadFields(stripTopicsSection(md))) +
					topicsFooterHTML(md)),
			Education: educationHTML(education),
			Starter:   starter,
			Debug:     debug,
			// goFiles is sorted, so this is the first non-test .go file by name,
			// which is not necessarily the starter (shortest name).
			File: firstOr(goFiles, ""),
		}

		grp := groupKey(slug)
		groups[grp] = append(groups[grp], catItem{
			ID:    slug,
			Name:  cleanTitle(title, slug),
			Sub:   cleanDirName(parts[2]),
			Lv:    strings.ToUpper(parts[0][:1]),
			Level: level,
			// Never locked: the local runner serves every puzzle. Playability is
			// gated at run time by whether the backend is connected.
			Locked: false,
			Tag:    tag,
		})
	}

	names := make([]string, 0, len(groups))
	for g := range groups {
		names = append(names, g)
	}
	sort.Slice(names, func(i, j int) bool {
		a, b := groups[names[i]][0].ID, groups[names[j]][0].ID
		ap, bp := strings.Split(a, "/"), strings.Split(b, "/")
		ao, bo := levelRank(ap[0]), levelRank(bp[0])
		if ao != bo {
			return ao < bo
		}
		return ap[1] < bp[1]
	})

	list := make([]catGroup, 0, len(names))
	for _, g := range names {
		items := groups[g]
		sort.Slice(items, func(i, j int) bool { return items[i].ID < items[j].ID })
		list = append(list, catGroup{Topic: g, Items: items})
	}
	return problems, list, nil
}

func levelRank(level string) int {
	if r, ok := levelOrder[level]; ok {
		return r
	}
	return 9
}

func firstOr(ss []string, def string) string {
	if len(ss) > 0 {
		return ss[0]
	}
	return def
}

func readFile(dir, name string) string {
	if name == "" {
		return ""
	}
	b, err := os.ReadFile(filepath.Join(dir, name))
	if err != nil {
		return ""
	}
	return string(b)
}

func write(path string, problems map[string]problem, groups []catGroup) error {
	var buf bytes.Buffer
	buf.WriteString("/* AUTO-GENERATED by site/cmd/gencatalog — do not edit by hand.\n")
	buf.WriteString("   Source of truth: the challenges/ tree. Re-run build.sh to refresh. */\n")
	buf.WriteString("window.PROBLEMS = ")
	p, err := marshalJSON(problems)
	if err != nil {
		return err
	}
	buf.Write(p)
	buf.WriteString(";\n\nwindow.CATALOG = ")
	c, err := marshalJSON(groups)
	if err != nil {
		return err
	}
	buf.Write(c)
	buf.WriteString(";\n")
	return os.WriteFile(path, buf.Bytes(), 0o644)
}

// marshal emits 2-space-indented JSON with HTML left intact (the descriptions
// are HTML) and non-ASCII escaped, so the file is safe to serve under any
// charset.
func marshal(v any) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return escapeNonASCII(bytes.TrimRight(buf.Bytes(), "\n")), nil
}

// escapeNonASCII rewrites every rune above U+007F as a \uXXXX escape. Non-ASCII
// only ever appears inside JSON strings here, so a flat scan is safe.
func escapeNonASCII(b []byte) []byte {
	var out bytes.Buffer
	for _, r := range string(b) {
		if r < 0x80 {
			out.WriteRune(r)
			continue
		}
		if r > 0xFFFF { // outside the BMP: surrogate pair
			r -= 0x10000
			fmt.Fprintf(&out, "\\u%04x\\u%04x", 0xD800+(r>>10), 0xDC00+(r&0x3FF))
			continue
		}
		fmt.Fprintf(&out, "\\u%04x", r)
	}
	return out.Bytes()
}
