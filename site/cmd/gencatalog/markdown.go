package main

// Minimal markdown → HTML, matching what the site's CSS expects. Not a general
// markdown implementation: it handles exactly the constructs the puzzle READMEs
// use (headings, lists, fenced code, inline code/bold/links) and passes raw HTML
// lines (e.g. <details>) through untouched.

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	reInlineCode = regexp.MustCompile("`([^`]+)`")
	reBold       = regexp.MustCompile(`\*\*([^*]+)\*\*`)
	reLink       = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
	reHeading    = regexp.MustCompile(`^(#{1,6})\s+(.*)`)
	reBullet     = regexp.MustCompile(`^[-*]\s+(.*)`)
)

func escapeHTML(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	return strings.ReplaceAll(s, ">", "&gt;")
}

func mdInline(s string) string {
	s = escapeHTML(s)
	s = reInlineCode.ReplaceAllString(s, "<code>$1</code>")
	s = reBold.ReplaceAllString(s, "<strong>$1</strong>")
	s = reLink.ReplaceAllString(s, "$1") // drop links, keep text
	return s
}

func mdToHTML(md string) string {
	var out []string
	inCode, inList := false, false

	closeList := func() {
		if inList {
			out = append(out, "</ul>")
			inList = false
		}
	}

	for _, ln := range strings.Split(md, "\n") {
		switch {
		case strings.HasPrefix(ln, "```"):
			if inCode {
				out = append(out, "</code></pre>")
				inCode = false
			} else {
				closeList()
				out = append(out, `<pre class="md"><code>`)
				inCode = true
			}
		case inCode:
			out = append(out, escapeHTML(ln))
		case strings.HasPrefix(strings.TrimLeft(ln, " \t"), "<"):
			// raw HTML block (details/summary/etc): pass through untouched
			closeList()
			out = append(out, ln)
		case reHeading.MatchString(ln):
			closeList()
			m := reHeading.FindStringSubmatch(ln)
			lvl := len(m[1])
			out = append(out, fmt.Sprintf("<h%d>%s</h%d>", lvl, mdInline(m[2]), lvl))
		case reBullet.MatchString(ln):
			if !inList {
				out = append(out, "<ul>")
				inList = true
			}
			out = append(out, "<li>"+mdInline(reBullet.FindStringSubmatch(ln)[1])+"</li>")
		case strings.TrimSpace(ln) == "":
			closeList()
		default:
			closeList()
			out = append(out, "<p>"+mdInline(ln)+"</p>")
		}
	}
	if inCode {
		out = append(out, "</code></pre>")
	}
	closeList()
	return strings.Join(out, "\n")
}
