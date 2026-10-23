# Snip indicator width regression

The figure compares measured output widths from `text.Snip("abcdef", limit, indicator)`.
`Plain` uses `"..."`; `ANSI` uses the same indicator wrapped in yellow SGR sequences.
Widths are measured with `text.StringWidthWithoutEscSequences`.

- Before: `9a59d895c26f170f10b0c16a5326c6e43c045fdc`.
- After: `a1351005aef385002c0118bea3655f44b68dd4d2`.
- Limits: 1 and 2 display columns.
- Figure: 16 cm wide, 420 dpi. Bars start at zero; hatching distinguishes revisions without relying on color alone.

The two JSON files contain the raw inputs, outputs, revision identifiers and measured widths.
To reproduce a measurement, place `measure-snip.go` next to a checkout of the relevant revision, then run this from that checkout:

```sh
go run ../measure-snip.go "$(git rev-parse HEAD)"
```

All plotted values are measured program outputs, not illustrative estimates.
