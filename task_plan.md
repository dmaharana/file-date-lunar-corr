# Task Plan - File Date Lunar Correlation CLI

Implement a Go CLI that recursively scans files, correlates their creation dates with lunar phases, and performs statistical analysis.

## Phases

- [x] **Phase 1: Research and Setup**
    - Search for Go libraries for lunar phases.
    - Search for Go libraries for statistics (chi-squared).
    - Investigate how to get file creation dates on Linux in Go.
    - Initialize Go module.
- [x] **Phase 2: Core Logic - File Traversal**
    - Implement recursive file walking.
    - Extract file creation/modification dates.
- [x] **Phase 3: Core Logic - Lunar Phase Mapping**
    - Integrate lunar phase calculation.
    - Bin phases into segments (8 segments as per spec).
- [x] **Phase 4: Core Logic - Statistical Analysis**
    - Implement/Integrate Chi-squared goodness-of-fit test.
- [x] **Phase 5: CLI and Output**
    - Implement CLI interface (flags for path, output CSV).
    - CSV export functionality.
    - Result reporting (counts, p-value).
- [x] **Phase 6: Testing and Validation**
    - Unit tests for phase mapping and statistics.
    - Integration test with sample directory.

## Decisions
- Used `github.com/dyindude/moonphase` for lunar phases.
- Used `gonum.org/v1/gonum/stat/distuv` for statistical distribution.
- Used `unix.Statx` for file birthtime on Linux.
- Implemented 8 binned phases centered around the standard phase points.

## Errors Encountered
| Error | Attempt | Resolution |
|-------|---------|------------|
| unused "os" import | 1 | Removed import |
