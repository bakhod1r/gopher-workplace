package main

// Pulling structured bits out of a puzzle README: title, `**Name:** value` meta
// lines, and the `## Topics` section that becomes the topic chips.

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var (
	reTitle       = regexp.MustCompile(`(?m)^#\s+(.*)`)
	reH1Line      = regexp.MustCompile(`^#\s+.*\n`)
	reMetaLine    = regexp.MustCompile(`(?m)^\*\*(Level|Topic|Estimated time):\*\*.*\n?`)
	reTopicsTail  = regexp.MustCompile(`(?s)\n##\s+Topics\b.*$`)
	reNumPrefix   = regexp.MustCompile(`^\d+-`)
	reChallengeNo = regexp.MustCompile(`^Challenge\s+[A-Za-z]*\d+\s*[—\-:]\s*`)
	reTopicMain   = regexp.MustCompile(`^([^(]+)(?:\(([^)]*)\))?`)
	reSolutionDet = regexp.MustCompile(`<details>(\s*<summary>Solution</summary>)`)
	reTopicsHead  = regexp.MustCompile(`\n##\s+Topics\b`)
	reNextH2      = regexp.MustCompile(`\n##\s`)
	reArrowSplit  = regexp.MustCompile(`[→›>]`)
)

// field reads a `**Name:** value` line from the README, else "".
func field(md, name string) string {
	re := regexp.MustCompile(`\*\*` + regexp.QuoteMeta(name) + `:\*\*\s*(.+)`)
	if m := re.FindStringSubmatch(md); m != nil {
		return strings.TrimSpace(m[1])
	}
	return ""
}

func titleOf(md, slug string) string {
	if m := reTitle.FindStringSubmatch(md); m != nil {
		return strings.TrimSpace(m[1])
	}
	parts := strings.Split(slug, "/")
	return parts[len(parts)-1]
}

// stripHeadFields removes the H1 + Level/Topic/Estimated meta lines; headerHTML
// renders them instead.
func stripHeadFields(md string) string {
	if loc := reH1Line.FindStringIndex(md); loc != nil && loc[0] == 0 {
		md = md[loc[1]:]
	}
	md = reMetaLine.ReplaceAllString(md, "")
	return strings.TrimLeft(md, "\n")
}

// stripTopicsSection drops a trailing `## Topics` section — the styled footer
// replaces it.
func stripTopicsSection(md string) string {
	return reTopicsTail.ReplaceAllString(md, "\n")
}

// topicsSection returns the body of the `## Topics` section, or "".
func topicsSection(md string) string {
	head := reTopicsHead.FindStringIndex(md)
	if head == nil {
		return ""
	}
	rest := md[head[1]:]
	if next := reNextH2.FindStringIndex(rest); next != nil {
		return rest[:next[0]]
	}
	return rest
}

// topicsFromSection lists concept names from `## Topics`, in order.
//
// The section is usually a table whose first column bolds the concept
// (`**iota**`) and whose later columns hold example code in backticks. Only the
// bolded names become chips — otherwise hint/solution snippets like `_ = iota`
// leak out. Falls back to backticked items when nothing is bolded (older
// list-style sections).
func topicsFromSection(md string) []string {
	body := topicsSection(md)
	if body == "" {
		return nil
	}
	var out []string
	for _, m := range reBold.FindAllStringSubmatch(body, -1) {
		out = append(out, strings.TrimSpace(m[1]))
	}
	if len(out) > 0 {
		return out
	}
	for _, m := range reInlineCode.FindAllStringSubmatch(body, -1) {
		out = append(out, m[1])
	}
	return out
}

// topicTags are the chips for the Topics footer. Prefer the `## Topics`
// section; fall back to a `**Topic:**` meta line.
func topicTags(md string) []string {
	if sect := topicsFromSection(md); len(sect) > 0 {
		return sect
	}
	raw := field(md, "Topic")
	if raw == "" {
		return nil
	}
	var tags []string
	if m := reTopicMain.FindStringSubmatch(raw); m != nil {
		main := strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(m[1]), "→"))
		for _, part := range reArrowSplit.Split(main, -1) { // "A → B" → both sides
			if p := strings.TrimSpace(part); p != "" {
				tags = append(tags, p)
			}
		}
		if m[2] != "" {
			for _, t := range strings.Split(m[2], ",") {
				if t = strings.Trim(strings.TrimSpace(t), "`"); t != "" {
					tags = append(tags, t)
				}
			}
		}
	}
	seen := map[string]bool{}
	var out []string
	for _, t := range tags {
		k := strings.ToLower(t)
		if !seen[k] {
			seen[k] = true
			out = append(out, t)
		}
	}
	return out
}

func topicsFooterHTML(md string) string {
	tags := topicTags(md)
	if len(tags) == 0 {
		return ""
	}
	var chips strings.Builder
	for _, t := range tags {
		fmt.Fprintf(&chips, `<span class="tchip">%s</span>`, mdInline(t))
	}
	return fmt.Sprintf("\n<details class=\"topics-foot\"><summary class=\"topics-lbl\">"+
		"Topics <span class=\"topics-n\">%d</span></summary>"+
		"<div class=\"tchips\">%s</div></details>", len(tags), chips.String())
}

// headerHTML is the LeetCode-style header: title + level pill + estimated time.
func headerHTML(md, title string) string {
	level := field(md, "Level")
	est := field(md, "Estimated time")
	meta := ""
	if level != "" {
		meta += fmt.Sprintf(`<span class="pill lv-%s">%s</span>`, strings.ToLower(level), mdInline(level))
	}
	if est != "" {
		meta += fmt.Sprintf(`<span class="p-time">%s</span>`, mdInline(est))
	}
	return fmt.Sprintf(`<div class="p-head"><h1 class="p-title">%s</h1>`+
		`<div class="p-meta">%s</div></div>`, mdInline(title), meta)
}

// markSolutionNocopy tags the Solution <details> so the frontend can block
// selection/copy.
func markSolutionNocopy(html string) string {
	return reSolutionDet.ReplaceAllString(html, `<details class="nocopy">$1`)
}

// cleanDirName turns "03-composite-types" into "Composite Types".
func cleanDirName(part string) string {
	return titleCase(strings.ReplaceAll(reNumPrefix.ReplaceAllString(part, ""), "-", " "))
}

// titleCase mirrors Python's str.title(): upper after any non-letter, lower
// elsewhere.
func titleCase(s string) string {
	prevLetter := false
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) {
			if prevLetter {
				b.WriteRune(unicode.ToLower(r))
			} else {
				b.WriteRune(unicode.ToUpper(r))
			}
			prevLetter = true
		} else {
			b.WriteRune(r)
			prevLetter = false
		}
	}
	return b.String()
}

// cleanTitle drops a "Challenge NN — " prefix; keeps the name.
func cleanTitle(title, slug string) string {
	t := strings.TrimSpace(reChallengeNo.ReplaceAllString(title, ""))
	if t != "" {
		return t
	}
	parts := strings.Split(slug, "/")
	return cleanDirName(parts[len(parts)-1])
}

// groupKey is the stable sidebar group from the path: "<Level> · <Topic dir>".
func groupKey(slug string) string {
	parts := strings.Split(slug, "/")
	return capitalize(parts[0]) + " · " + cleanDirName(parts[1])
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}
