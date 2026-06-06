# Findings - File Date Lunar Correlation CLI

## Lunar Phase Calculation
- **Library:** `github.com/dyindude/moonphase`
- **Method:** `m.Phase()` returns 0.0 - 1.0 (0=New, 0.5=Full).
- **Segments:** Can be easily binned into 8 segments (0-0.125, 0.125-0.25, etc.).

## Statistical Analysis
- **Library:** `gonum.org/v1/gonum/stat`
- **Chi-squared test:**
    - Formula: `Σ (O - E)^2 / E`
    - p-value: `1.0 - distuv.ChiSquared{K: df}.CDF(chi2)`
    - DF: `len(bins) - 1` (usually 7 for 8 phases).

## File Creation Date on Linux
- **System Call:** `unix.Statx` from `golang.org/x/sys/unix`.
- **Mask:** `unix.STATX_BTIME`.
- **Fallback:** Modification time (`mtime`) if `BTIME` is not supported by the filesystem or kernel.
- **Note:** Requires Kernel 4.11+.

## Architecture
- Recursive walker using `path/filepath.WalkDir`.
- Worker pool or simple sequential processing (given file I/O is the bottleneck). Sequential is fine for a start.
- CSV output using `encoding/csv`.
- CLI using standard `flag` package or `spf13/cobra`. Standard `flag` is sufficient for this simple tool.
