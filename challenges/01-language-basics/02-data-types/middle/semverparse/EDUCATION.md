# Field-by-field numeric parsing

## The idea

Parse a structured numeric string in one pass: accumulate digits with
`n = n*10 + (c-'0')`, and on a separator flush the field and move on. Validate
strictly — count fields, forbid empty ones and stray characters.

## Why it matters

Version tags, dotted-decimal IPs, and timestamps are parsed exactly this way in
real tooling. Doing it by hand (rather than `Split`+`Atoi`) shows the byte-level
mechanics and forces explicit validation.

## Watch out

- Reject an empty field (`1..2`) and a trailing dot.
- Enforce the exact field count; `1.2` and `1.2.3.4` are both invalid here.
- A leading `-` or letters must fail — only `'0'..'9'` and `.` are valid.
