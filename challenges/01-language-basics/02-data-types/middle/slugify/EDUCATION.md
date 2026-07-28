# Normalizing text to a slug

## The idea

Scan runes, keep alphanumerics (lowercased), and collapse every run of other
characters into a single `-`. A neat trick avoids trailing dashes: instead of
appending a dash on each separator, set a *pending* flag and only emit the dash
right before the next kept character.

## Why it matters

Slugs, filenames, and identifiers are generated from free text everywhere. The
"collapse runs + trim edges" shape recurs; the pending-separator technique keeps
it single-pass with no post-trim.

## Watch out

- Lowercase with `unicode.ToLower` (or ASCII arithmetic).
- Don't emit a leading dash: the pending flag naturally handles this if you only
  flush before real output.
- Decide how to treat non-ASCII letters (this spec keeps only ASCII alnum).
