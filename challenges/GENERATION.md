# Gopher Workplace — Puzzle Generation Spec

How new puzzles are authored so the curriculum stays **monotonic**: a learner
never meets a concept before it was taught. Read this together with
[PLAN.md](PLAN.md) (structure) and copy [`_template/`](_template/) (shape).

---

## 1. The learning path is a total order

Every subtopic slot has one position in a single ordered sequence. The order is:

```
level  →  topic (01..17)  →  subtopic (01..NN)
```

with **level** as the outermost key:

```
junior  <  middle  <  senior  <  staff
```

So the full path is:

```
junior/01/01, junior/01/02, … junior/17/NN,
middle/01/01, …               middle/17/NN,
senior/01/01, …               senior/17/NN,
staff/01/01,  …               staff/17/NN
```

Call a slot's position `P`. The **covered set** at `P` = every subtopic whose
position is `≤ P` (this slot and everything before it).

## 2. Scope rule (the one law)

> A puzzle at position `P` may only rely on Go concepts introduced **at or
> before `P`** in the learning path. Never use a not-yet-taught concept.

- The subtopic at `P` is the concept **being taught** — it is the point of the
  puzzle and must be exercised.
- Everything in the covered set (`< P`) is **fair game** — reuse freely.
- Everything after `P` is **forbidden** — if solving the puzzle needs it, the
  puzzle is mis-scoped. Move it later or simplify.

## 3. Level accumulation (a consequence, stated explicitly)

Because level is the outermost key, a level fully precedes the next. Therefore:

| Puzzle level | Covered topics available |
|--------------|--------------------------|
| junior       | junior topics up to `P` |
| middle       | **all junior** + middle up to `P` |
| senior       | **all junior + middle** + senior up to `P` |
| staff        | **all junior + middle + senior** + staff up to `P` |

Example: a `middle` puzzle may use *every* junior topic (all 17), plus the
middle topics taught before it. A junior puzzle may **not** reach forward into
middle.

## 4. Authoring checklist (scope gate)

Before writing a puzzle at `<level>/<NN-topic>/<MM-subtopic>/<name>`:

1. **Locate `P`.** Note its level, topic index, subtopic index.
2. **List the covered set.** All junior (if middle+), all earlier topics this
   level, earlier subtopics this topic.
3. **Pick the authoring mode (§5e)** — from-scratch stub for junior/concept
   puzzles, planted-bug for senior/staff/debugging puzzles.
4. **Scope-audit the solution.** Write the intended solution in your head; every
   language feature / stdlib package it uses must be in the covered set or be
   the target subtopic. If not — rescope.
5. Build from [`_template/`](_template/) (next section).
6. Verify red→green: the shipped stub/bug **fails** tests, the solution
   **passes**.
   Run `make verify`.
7. Delete the slot's `.keep`.

## 5. Generation template

Copy `_template/` into the slot and fill the `{{...}}` placeholders:

```
_template/
  go.mod.tmpl        module path = the slot path
  Makefile           standard verify gate (copy as-is, no edits)
  puzzle.go.tmpl     doc comment + implement-from-scratch stub (see §5e)
  puzzle_test.go.tmpl  red-until-fixed table tests
  README.tmpl.md     Context / Task / Examples(≥3) / Topics to Master / hint
```

Placeholders:

| Placeholder | Meaning |
|-------------|---------|
| `{{LEVEL}}` | junior \| middle \| senior \| staff |
| `{{TOPIC}}` | `NN-topic` dir name |
| `{{SUBTOPIC}}` | `MM-subtopic` dir name |
| `{{NAME}}` | puzzle/package short name (also go file base) |
| `{{PKG}}` | Go package name (= NAME, lowercase, no dashes) |
| `{{FUNC}}` | exported entry function under test |
| `{{TITLE}}` | human title for README |

Module path convention:

```
github.com/gopher-workplace/challenges/{{LEVEL}}/{{TOPIC}}/{{SUBTOPIC}}/{{NAME}}
```

## 5b. Difficulty dial — constraints per level

Scope (§2) controls *which concepts* appear. This controls *how hard the same
concept is pushed*. Higher levels don't just add new topics — they add
**resource + scale constraints** on top of correctness.

| Level  | What the puzzle demands beyond "passes" |
|--------|-----------------------------------------|
| junior | correctness only. Small inputs. |
| middle | correctness + basic efficiency (no needless O(n²), no obvious extra allocs). |
| senior | correctness **at scale**: 100M-row input, a **RAM ceiling** (stream, don't buffer), bounded allocations. |
| staff  | senior + a **CPU/time ceiling** and concurrency: hit a deadline under load, no data race, parallelism must actually help. |

How to encode a constraint in the puzzle (so it is graded, not just prose):

- **Scale / RAM:** a test that feeds a large or streaming input and asserts the
  result; pair with `testing.B` + `-benchmem` and a bounded
  `b.ReportMetric` / `AllocsPerOp` check, or make the naive buffered solution
  OOM/blow up so only the streaming fix survives.
- **CPU / time:** `testing.B` with a target ns/op, or a `context` deadline the
  solution must meet; a `-race`-clean concurrency test for staff.
- Always keep a plain red→green correctness test too — the constraint is
  *additional*, never a replacement.

State the active constraint explicitly in the README **Task** ("must run in one
pass over the input", "must stay under N allocations", "must finish before the
context deadline") so the candidate knows what is graded.

Constraints must still obey §2: only use covered concepts to satisfy them. A
`middle` puzzle cannot require a senior-only technique to hit its limit.

## 5c. Cumulative rule — reuse earlier topics (spiral)

Scope (§2) says a puzzle *may* use covered concepts. This says it *should*.

> Every puzzle exercises its **target subtopic** and **weaves in earlier
> covered concepts** so the learner keeps practicing them. Knowledge is
> cumulative, not siloed.

- The **target subtopic** is the new thing under test — always central.
- On top of it, deliberately require **≥1–2 prior covered concepts** to reach
  the fix (more at higher levels). Later puzzles lean on more history.
- Never contrive it: the reused concept must be natural to the problem, not
  bolted on.
- The README **Topics to Master** table then lists the target **plus** the
  reused prior concepts — so the table itself shows the accumulation.

Get the exact allowed history for a slot with:

```bash
scripts/coverage.sh <level>/<NN-topic>/<MM-subtopic>   # full covered list
scripts/scaffold.sh <slot> <name> <Func>               # create slot + inject it
```

`scaffold.sh` copies `_template/` into the slot and drops the covered set into
the puzzle as a `COVERED:` reference block, so the author designs the bug
against real prior material.

## 5d. Depth ladder — how deep you must understand Go

Three independent axes now: **scope** (§2, which concepts), **difficulty dial**
(§5b, scale/resource pressure), and **depth** — how far *under* the language you
must see to explain the bug. Depth rises with level.

> As the level rises, the puzzle can only be understood by knowing Go **all the
> way down** — surface behaviour at junior, runtime/memory-model/compiler at
> staff.

| Level  | Depth the bug lives at | Model the solver must hold |
|--------|------------------------|----------------------------|
| junior | **surface / API** — visible language behaviour. | What the code *does*: slice header, map, error value, receiver. |
| middle | **mechanics** — how the construct is implemented. | Backing arrays + cap growth, interface = (type,word) pair, defer/panic order, goroutine + channel semantics. |
| senior | **runtime / cost model** — allocation, GC, escape, scheduling. | Escape analysis (heap vs stack), GC reachability + leaks, `sync` internals, allocation cost, `context` propagation. |
| staff  | **internals / memory model** — the guarantees themselves. | Go memory model (happens-before), scheduler + GOMAXPROCS, false sharing/cache lines, compiler/inlining, `unsafe`, atomics. |

Rules:

- A puzzle is at the **right depth** when the *shallow* reading of the bug looks
  correct, and only the deeper model reveals why it fails. Junior: obvious once
  seen. Staff: needs the memory model / runtime to even see it.
- The README **hint** and **Topics to Master** must pitch at that depth — a
  staff hint talks happens-before / escape, not "check your loop bound".
- Depth still obeys §2 scope: the deep concept must be a covered topic. Runtime,
  advanced-topics, performance-engineering topics naturally sit at senior/staff
  positions, which is what unlocks the deeper puzzles there.
- Pick the archetype from §7 whose natural depth matches the level (data race,
  false sharing, escape → staff; aliasing, error-wrap → junior/middle).

## 5e. Two authoring modes (pick by level)

A puzzle ships **red** (tests fail) and the learner drives it **green**. There
are two ways to stage that red state — pick by what the level should train:

- **Implement-from-scratch** — the default for **junior fundamentals** (and any
  puzzle whose point is *learning a concept*). Ship a stub: the real signature +
  a doc comment stating intended behaviour + `panic("not implemented")` (and a
  neutral placeholder like `const X = 0` where a value is needed to compile).
  No `CHANGE CODE` markers, no example solution in the code. The learner writes
  the concept themselves — deeper mastery through recall + synthesis. This is
  what the variables-and-constants exemplars use.

- **Planted-bug** — for **senior/staff** and anywhere the goal is *reading and
  debugging real code*. Ship working-looking code with exactly **one** bug
  between `// CHANGE CODE BELOW/ABOVE THIS LINE`. Trains recognition + debugging,
  closer to on-the-job work.

Both verify identically (§4.6): red before, green after. Choose one per puzzle;
do not mix markers into a from-scratch stub.

**Offering both on the site (optional).** A from-scratch puzzle may also ship a
planted-bug variant as a sibling file `<name>.debug.txt` (a compiles-but-buggy
copy of the same package, red against the same tests). `gen-problems.py` picks
it up into the problem's `debug` field, and the playground shows a **learn ⇄
debug** toggle. The real `.go` file stays the stub; `.debug.txt` is never
compiled by `go` (not a `.go` file), so it can define the same symbols without
clashing. Keep it in sync with the tests when you edit the puzzle.

## 5f. Warning guards — anti-hardcode & clean-code (non-blocking on Run)

Some tasks can be "passed" by cheating the *letter* of the tests — hand-typing
magic numbers instead of deriving them, or hardcoding outputs instead of
computing them. Catch this with a **warning guard test**: a test that inspects
the submitted source and, when the rule is broken, emits a line starting with
`WARN:` **without failing** (use `t.Logf`, not `t.Error`/`t.Fatal`).

The runner treats warnings specially:
- A **Run** ignores them (still shows PASSED) — they are advisory.
- A **Submit** is *not accepted* (does not count as solved) while any warning is
  present. The candidate must resolve them first.
- The runner also auto-adds a `not gofmt-clean` warning for an unformatted
  submission, so **clean formatting is required to submit** (the Format button
  fixes it).

Write a guard by parsing the candidate file with `go/parser` (pass `0` for the
mode so comments are *excluded* — a rule word mentioned in a comment must not
satisfy the check) and inspecting the AST:

```go
func TestNoHardcodedConstants(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "<pkg>.go", nil, 0) // 0 = skip comments
	if err != nil {
		return // parse trouble is not this check's concern
	}
	ok := false
	ast.Inspect(f, func(n ast.Node) bool {
		if id, ok2 := n.(*ast.Ident); ok2 && id.Name == "iota" { ok = true }
		return true
	})
	if !ok {
		t.Logf("WARN: derive the constants from iota, don't hand-type magic numbers")
	}
}
```

Guidance:
- The guard reads `<pkg>.go` by name; the runner materializes the submitted
  source under that name, so the check sees the candidate's code.
- Prefer a robust positive signal (e.g. "the input parameter is referenced",
  "`iota` is used") over blacklisting specific literals — blacklists both miss
  variants and false-positive on comments/strings.
- Keep the `WARN:` message actionable — say what to do, not just what's wrong.
- Exemplars: `byteunits` (must use `iota`), `temperature` (`CToF` must use its
  input `c`).

## 6. Conventions carried from the exemplar

- One authoring mode per puzzle (§5e): from-scratch stub *or* one planted bug
  between `// CHANGE CODE BELOW/ABOVE THIS LINE` — never both.
- Do not change the function signature or the tests in the task.
- README **Topics to Master** lists only covered-set + target concepts.
- No `SOLUTION.md`; no roadmap-path references in READMEs.
- Each puzzle is its own module (own `go.mod`), `go 1.26`.
