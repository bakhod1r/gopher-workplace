#!/usr/bin/env python3
"""Generate 10-advanced-topics puzzles from compact specs.

Layout produced (mirrors HIERARCHY.md):
    challenges/10-advanced-topics/<subtopic>/<level>/<name>/
        go.mod Makefile <name>.go <name>_test.go README.md EDUCATION.md

Specs live in scripts/advanced_specs/<subtopic-key>_<level>.py, each exporting
SPECS: a list of dicts. Run:  python3 scripts/gen_advanced.py [module ...]
"""

from __future__ import annotations

import importlib.util
import os
import pathlib
import subprocess
import sys
import textwrap

ROOT = pathlib.Path(__file__).resolve().parent.parent
CH = ROOT / "challenges"
TOPIC = "10-advanced-topics"
SPEC_DIR = ROOT / "scripts" / "advanced_specs"

MAKEFILE = (CH / "_template" / "Makefile").read_text()


def gofile(spec: dict, body: str, markers: bool, solved: bool = False) -> str:
    name = spec["name"]
    parts = [f'// Package {name} — Gopher Workplace challenge.\npackage {name}\n']
    imports = spec.get("imports", [])
    if solved and "sol_imports" in spec:
        imports = spec["sol_imports"]
    if imports:
        if len(imports) == 1:
            parts.append(f'import {imports[0]}\n')
        else:
            inner = "\n".join("\t" + i for i in imports)
            parts.append("import (\n" + inner + "\n)\n")
    extra = spec.get("extra")
    if solved and "sol_extra" in spec:
        extra = spec["sol_extra"]
    if extra:
        parts.append(extra.strip() + "\n")
    doc = "\n".join("// " + l if l else "//" for l in spec["doc"].strip().split("\n"))
    inner = textwrap.indent(textwrap.dedent(body).strip(), "\t")
    if markers:
        inner = "\t// CHANGE CODE BELOW THIS LINE\n" + inner + "\n\t// CHANGE CODE ABOVE THIS LINE"
    parts.append(f'{doc}\n{spec["sig"]} {{\n{inner}\n}}\n')
    return "\n".join(parts)


def readme(spec: dict) -> str:
    ex = []
    for i, e in enumerate(spec["examples"], 1):
        block = f'**Example {i}:**\n\n```\nInput:  {e[0]}\nOutput: {e[1]}\n```\n'
        if len(e) > 2 and e[2]:
            block += f'\n_Explanation:_ {e[2]}\n'
        ex.append(block)
    topics = "\n".join(
        f'| {i} | **{t[0]}** | {t[1]} |' for i, t in enumerate(spec["topics"], 1)
    )
    tasks = "\n".join(f'{i}. {t}' for i, t in enumerate(spec["task"], 1))
    mode_note = (
        f'Change only the code between the `CHANGE CODE` markers.'
        if spec["mode"] == "bug"
        else f'Replace the stub body in [{spec["name"]}.go]({spec["name"]}.go) with a working implementation.'
    )
    return f"""# {spec["title"]}

**Level:** {spec["level"]}
**Topic:** {TOPIC} / {spec["sub"]}

## Context

{spec["context"].strip()}

## Task

{"Fix the single planted bug in" if spec["mode"] == "bug" else "Implement"} [{spec["name"]}.go]({spec["name"]}.go):

{tasks}

{mode_note}

## Examples

{chr(10).join(ex)}
## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
{topics}

## Hint

{spec["hint"].strip()}

## Validate

```bash
make verify
```
"""


def education(spec: dict) -> str:
    approach = "\n".join(f'{i}. {s}' for i, s in enumerate(spec["approach"], 1))
    pitfalls = "\n".join(f'- {p}' for p in spec["pitfalls"])
    sol = gofile(spec, spec["solution"], markers=False, solved=True)
    sol = sol.split("\n\n", 1)[1] if "\n\n" in sol else sol
    return f"""# {spec["title"]}

## Intuition

{spec["intuition"].strip()}

## Approach

{approach}

## Solution

```go
{sol.strip()}
```

## Walkthrough

{spec["walkthrough"].strip()}

## Pitfalls

{pitfalls}
"""


def emit(spec: dict) -> pathlib.Path:
    slot = CH / TOPIC / spec["level"] / spec["name"]
    slot.mkdir(parents=True, exist_ok=True)
    mod = f'github.com/gopher-workplace/challenges/{TOPIC}/{spec["level"]}/{spec["name"]}'
    (slot / "go.mod").write_text(f"module {mod}\n\ngo 1.26\n")
    (slot / "Makefile").write_text(MAKEFILE)
    if os.environ.get("GEN_SOLVED") == "1":
        body = spec["solution"]
    else:
        body = spec["buggy"] if spec["mode"] == "bug" else spec["stub"]
    solved = os.environ.get("GEN_SOLVED") == "1"
    (slot / f'{spec["name"]}.go').write_text(
        gofile(spec, body, spec["mode"] == "bug" and not solved, solved)
    )
    (slot / f'{spec["name"]}_test.go').write_text(
        f'package {spec["name"]}\n\n' + textwrap.dedent(spec["tests"]).strip() + "\n"
    )
    (slot / "README.md").write_text(readme(spec))
    (slot / "EDUCATION.md").write_text(education(spec))
    subprocess.run(["gofmt", "-w", str(slot)], check=True)
    return slot


def load(mod_path: pathlib.Path) -> list:
    spec = importlib.util.spec_from_file_location(mod_path.stem, mod_path)
    m = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(m)
    return m.SPECS


def main() -> None:
    args = sys.argv[1:]
    files = (
        [SPEC_DIR / f"{a}.py" for a in args]
        if args
        else sorted(SPEC_DIR.glob("*.py"))
    )
    n = 0
    for f in files:
        for s in load(f):
            emit(s)
            n += 1
        print(f"{f.stem}: ok")
    print(f"generated {n} puzzles")


if __name__ == "__main__":
    main()
