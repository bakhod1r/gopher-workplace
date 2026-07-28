#!/usr/bin/env python3
"""Generate web/assets/js/problems.js from the challenges/ tree.

Walks every puzzle (a dir containing go.mod), pulls the starter code from its
non-test .go file and the description from README.md, and emits window.PROBLEMS
+ window.CATALOG. Puzzles with a wasm runner (see RUNNERS) are runnable; the
rest are shown but locked ("coming soon").
"""
import json
import os
import re
import sys

ROOT = os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
CH = os.path.join(ROOT, "challenges")
OUT = os.path.join(ROOT, "site", "web", "assets", "js", "problems.js")

# Local-runner-only build: the in-browser wasm path is gone. Every puzzle runs
# against the real Go toolchain via the localhost backend (site/cmd/localrunner),
# so there are no per-puzzle wasm functions to register.
RUNNERS = {}

LEVELS = ["junior", "middle", "senior", "staff"]

# Puzzles hidden from the site (files kept on disk, just not published).
EXCLUDE = set()


def find_puzzles():
    out = []
    for dirpath, _dirs, files in os.walk(CH):
        if "go.mod" in files:
            out.append(dirpath)
    return sorted(out)


def starter_code(pdir):
    """The stub the candidate edits: the non-test, non-cmd .go file."""
    cands = []
    for f in os.listdir(pdir):
        if f.endswith(".go") and not f.endswith("_test.go"):
            cands.append(f)
    if not cands:
        return ""
    cands.sort(key=len)  # shortest name = the main file
    with open(os.path.join(pdir, cands[0])) as fh:
        return fh.read()


def debug_code(pdir):
    """Optional planted-bug variant for Debug mode: `<name>.debug.txt`.

    A puzzle ships an implement-from-scratch stub as its real .go file; an
    optional sibling `*.debug.txt` holds a compiles-but-buggy version so the
    site can offer a debugging exercise on the same tests. Empty if absent.
    """
    for f in sorted(os.listdir(pdir)):
        if f.endswith(".debug.txt"):
            with open(os.path.join(pdir, f)) as fh:
                return fh.read()
    return ""


def md_inline(s):
    s = (s.replace("&", "&amp;").replace("<", "&lt;").replace(">", "&gt;"))
    s = re.sub(r"`([^`]+)`", r"<code>\1</code>", s)
    s = re.sub(r"\*\*([^*]+)\*\*", r"<strong>\1</strong>", s)
    s = re.sub(r"\[([^\]]+)\]\(([^)]+)\)", r"\1", s)  # drop links, keep text
    return s


def md_to_html(md):
    """Minimal markdown -> HTML. Raw HTML lines (e.g. <details>) pass through."""
    html, i, lines = [], 0, md.split("\n")
    in_code, in_list = False, False

    def close_list():
        nonlocal in_list
        if in_list:
            html.append("</ul>")
            in_list = False

    while i < len(lines):
        ln = lines[i]
        if ln.startswith("```"):
            if in_code:
                html.append("</code></pre>")
                in_code = False
            else:
                close_list()
                html.append('<pre class="md"><code>')
                in_code = True
            i += 1
            continue
        if in_code:
            html.append(ln.replace("&", "&amp;").replace("<", "&lt;").replace(">", "&gt;"))
            i += 1
            continue
        # raw HTML block (details/summary/etc): pass through untouched
        if ln.lstrip().startswith("<"):
            close_list()
            html.append(ln)
            i += 1
            continue
        m = re.match(r"(#{1,6})\s+(.*)", ln)
        if m:
            close_list()
            lvl = len(m.group(1))
            html.append("<h%d>%s</h%d>" % (lvl, md_inline(m.group(2)), lvl))
            i += 1
            continue
        m = re.match(r"[-*]\s+(.*)", ln)
        if m:
            if not in_list:
                html.append("<ul>")
                in_list = True
            html.append("<li>%s</li>" % md_inline(m.group(1)))
            i += 1
            continue
        if ln.strip() == "":
            close_list()
            i += 1
            continue
        close_list()
        html.append("<p>%s</p>" % md_inline(ln))
        i += 1
    if in_code:
        html.append("</code></pre>")
    close_list()
    return "\n".join(html)


def header_html(md, title):
    """LeetCode-style header: title + level pill + estimated time."""
    level = field(md, "Level")
    est = field(md, "Estimated time")
    lv_cls = level.lower() if level else ""
    meta = ""
    if level:
        meta += '<span class="pill lv-%s">%s</span>' % (lv_cls, md_inline(level))
    if est:
        meta += '<span class="p-time">%s</span>' % md_inline(est)
    return ('<div class="p-head"><h1 class="p-title">%s</h1>'
            '<div class="p-meta">%s</div></div>' % (md_inline(title), meta))


def strip_head_fields(md):
    """Remove the H1 + Level/Topic/Estimated meta lines; header_html renders them."""
    md = re.sub(r"^#\s+.*\n", "", md, count=1)
    md = re.sub(r"^\*\*(Level|Topic|Estimated time):\*\*.*\n?", "", md, flags=re.M)
    return md.lstrip("\n")


def strip_topics_section(md):
    """Drop a trailing `## Topics` section — the styled footer replaces it."""
    return re.sub(r"\n##\s+Topics\b.*\Z", "\n", md, flags=re.S)


def topics_from_section(md):
    """Concept names from the `## Topics` section, in order.

    The section is usually a table whose first column bolds the concept
    (`**iota**`) and whose later columns contain example code in backticks.
    Only the bolded concept names are chips — otherwise hint/solution snippets
    like `_ = iota` leak out. Falls back to backticked items when nothing is
    bolded (older list-style sections).
    """
    m = re.search(r"\n##\s+Topics\b(.*?)(?=\n##\s|\Z)", md, flags=re.S)
    if not m:
        return []
    body = m.group(1)
    bold = re.findall(r"\*\*([^*]+)\*\*", body)
    if bold:
        return [b.strip() for b in bold]
    return re.findall(r"`([^`]+)`", body)


def topic_tags(md):
    """Chips for the Topics footer / Learn row.

    Prefer the `## Topics` section; fall back to a `**Topic:**` meta line.
    """
    sect = topics_from_section(md)
    if sect:
        return sect
    raw = field(md, "Topic")
    if not raw:
        return []
    tags = []
    m = re.match(r"([^(]+)(?:\(([^)]*)\))?", raw)
    if m:
        main = m.group(1).strip().rstrip("→").strip()
        # split a "A → B" main into both sides
        for part in re.split(r"[→›>]", main):
            part = part.strip()
            if part:
                tags.append(part)
        if m.group(2):
            for t in m.group(2).split(","):
                t = t.strip().strip("`")
                if t:
                    tags.append(t)
    # de-dupe, keep order
    seen, out = set(), []
    for t in tags:
        k = t.lower()
        if k not in seen:
            seen.add(k)
            out.append(t)
    return out


def topics_footer_html(md):
    tags = topic_tags(md)
    if not tags:
        return ""
    chips = "".join('<span class="tchip">%s</span>' % md_inline(t) for t in tags)
    return ('\n<details class="topics-foot"><summary class="topics-lbl">'
            'Topics <span class="topics-n">%d</span></summary>'
            '<div class="tchips">%s</div></details>' % (len(tags), chips))


def mark_solution_nocopy(html):
    """Tag the Solution <details> so the frontend can block selection/copy."""
    return re.sub(
        r"<details>(\s*<summary>Solution</summary>)",
        r'<details class="nocopy">\1',
        html,
    )


def title_of(md, slug):
    m = re.search(r"^#\s+(.*)", md, re.M)
    if m:
        return m.group(1).strip()
    return slug.rsplit("/", 1)[-1]


def field(md, name):
    """Read a `**Name:** value` line from the README, else ''."""
    m = re.search(r"\*\*%s:\*\*\s*(.+)" % re.escape(name), md)
    return m.group(1).strip() if m else ""


def topic_label(md, slug):
    """Group label from README **Level** + **Topic** (not the path).

    New layout slug: <NN-topic>/<MM-subtopic>/<level>/<name>.
    """
    level = field(md, "Level")
    tags = topic_tags(md)
    topic = tags[0] if tags else ""
    parts = slug.split("/")
    if not topic:  # fallback: derive from the topic dir
        topic = clean_dir_name(parts[0])
    if not level:
        level = parts[2].capitalize()
    return "%s · %s" % (level, topic)


LV_ORDER = {"junior": 0, "middle": 1, "senior": 2, "staff": 3}
LV_SET = set(LEVELS)


def clean_dir_name(part):
    return re.sub(r"^\d+-", "", part).replace("-", " ").title()


def level_index(slug):
    """Index of the level segment in the slug, whatever the nesting depth.

    Most puzzles are <NN-topic>/<MM-subtopic>/<level>/<name> (level at parts[2]),
    but some subtopics nest one deeper (…/<subtopic>/<level>/<name>). Scan for
    the segment that names a level instead of hardcoding an index.
    """
    parts = slug.split("/")
    for i, p in enumerate(parts):
        if p in LV_SET:
            return i
    return 2  # fallback to the flat layout


def level_of(slug):
    parts = slug.split("/")
    i = level_index(slug)
    return parts[i] if i < len(parts) else "junior"


def group_key(slug):
    """Sidebar group = the topic dir only, so a topic is ONE collapsible group.

    Level is shown per-item (as a flag), not by splitting the topic into a group
    per level.
    """
    return clean_dir_name(slug.split("/")[0])


def clean_title(title, slug):
    """Drop a 'Challenge NN — ' / 'Challenge M01 — ' prefix; keep the name."""
    t = re.sub(r"^Challenge\s+[A-Za-z]*\d+\s*[—\-:]\s*", "", title).strip()
    return t or clean_dir_name(slug.split("/")[-1])


def main():
    problems, catalog = {}, {}
    for pdir in find_puzzles():
        slug = os.path.relpath(pdir, CH)
        if slug in EXCLUDE:
            continue
        starter = starter_code(pdir)
        debug = debug_code(pdir)
        readme_path = os.path.join(pdir, "README.md")
        md = ""
        if os.path.exists(readme_path):
            with open(readme_path) as fh:
                md = fh.read()
        title = title_of(md, slug)
        gofile = ""
        for f in sorted(os.listdir(pdir)):
            if f.endswith(".go") and not f.endswith("_test.go"):
                gofile = f
                break
        runner = RUNNERS.get(slug)
        level = (field(md, "Level") or level_of(slug)).lower()
        # Every runnable puzzle now goes through the local Go toolchain backend.
        has_tests = any(f.endswith("_test.go") for f in os.listdir(pdir))
        is_backend = has_tests
        problems[slug] = {
            "id": slug,
            "title": title,
            "level": level,
            "tag": "backend" if is_backend else "",
            "description": mark_solution_nocopy(
                            header_html(md, title)
                            + md_to_html(strip_head_fields(strip_topics_section(md)))
                            + topics_footer_html(md)),
            "starter": starter,
            "debug": debug,
            "file": gofile,
            "run": runner["run"] if runner else "",
            "runCustom": runner["runCustom"] if runner else "",
            "customDefault": runner.get("customDefault", "") if runner else "",
        }
        grp = group_key(slug)
        catalog.setdefault(grp, [])
        parts = slug.split("/")
        li = level_index(slug)
        catalog[grp].append({
            "id": slug,
            "name": clean_title(title, slug),
            "sub": clean_dir_name(parts[li - 1]) if li >= 1 else "",
            "lv": {"junior": "J", "middle": "M", "senior": "S",
                   "staff": "T"}.get(level_of(slug), "J"),
            "level": level,
            # Not locked: the local runner serves every puzzle. Playability is
            # gated at run time by whether the backend is connected.
            "locked": False,
            "tag": "backend" if is_backend else "",
            "_sort": slug,
        })

    def group_order(g):
        slug0 = catalog[g][0]["id"]  # any member shares the topic dir
        return slug0.split("/")[0]   # NN-topic prefix orders the topics

    order = sorted(catalog, key=group_order)
    for g in catalog:
        # Within a topic group: cluster by level (junior→staff), then path.
        catalog[g].sort(key=lambda it: (LV_ORDER.get(it["level"], 9), it["_sort"]))
        for it in catalog[g]:
            del it["_sort"]
    cat_list = [{"topic": g, "items": catalog[g]} for g in order]

    with open(OUT, "w") as fh:
        fh.write("/* AUTO-GENERATED by site/scripts/gen-problems.py — do not edit by hand.\n")
        fh.write("   Source of truth: the challenges/ tree. Re-run build.sh to refresh. */\n")
        fh.write("window.PROBLEMS = " + json.dumps(problems, indent=2) + ";\n\n")
        fh.write("window.CATALOG = " + json.dumps(cat_list, indent=2) + ";\n")

    print("wrote %s: %d puzzles (all run on the local Go toolchain backend)"
          % (os.path.relpath(OUT, ROOT), len(problems)))


if __name__ == "__main__":
    sys.exit(main())
