# File Date Lunar Correlation CLI

A Go-based command-line tool that recursively scans directories, extracts file creation dates, and correlates them with lunar phases to perform statistical analysis.

## Overview

This tool investigates the potential (though likely coincidental) relationship between file creation activity and the lunar cycle. It extracts the "birthtime" of files (on supported Linux filesystems) and maps each timestamp to one of eight lunar phases. Finally, it performs a Chi-squared goodness-of-fit test to determine if the distribution of file creations is statistically non-uniform.

## Features

- **Recursive Scanning**: Efficiently walks through subdirectories to gather file metadata.
- **Linux Birthtime Support**: Uses `unix.Statx` to retrieve the actual file creation date on supported filesystems (ext4, Btrfs, XFS, etc.).
- **Graceful Fallback**: Automatically falls back to modification time (`mtime`) if birthtime is not available.
- **Lunar Mapping**: Correlates dates with 8 binned lunar phases:
  - New Moon
  - Waxing Crescent
  - First Quarter
  - Waxing Gibbous
  - Full Moon
  - Waning Gibbous
  - Last Quarter
  - Waning Crescent
- **Statistical Significance**: Reports the Chi-squared statistic and p-value to test for non-random patterns.
- **Data Export**: Generates a detailed CSV report for further manual analysis.

## Installation

Ensure you have [Go](https://golang.org/doc/install) installed (1.21+ recommended).

```bash
# Clone the repository (or navigate to the project folder)
# Install dependencies
go mod download

# Build the binary
go build -o lunar-corr main.go
```

## Usage

Run the tool by specifying a directory to scan. By default, it scans the current directory and saves results to `file_dates.csv`.

```bash
./lunar-corr -dir /path/to/your/files -output my_analysis.csv
```

### Options

| Flag | Description | Default |
|------|-------------|---------|
| `-dir` | The root directory to start the recursive scan. | `.` |
| `-output` | The filename for the exported CSV data. | `file_dates.csv` |

## Understanding the Results

### Terminal Output
The tool prints a summary of file counts per lunar phase and the result of the Chi-squared test:
- **p > 0.05**: The distribution looks random (no significant correlation).
- **p ≤ 0.05**: The distribution is statistically non-uniform, suggesting a possible pattern.

### CSV Structure
The exported CSV contains:
- `Path`: Full path to the file.
- `Created Date`: The extracted creation or modification timestamp.
- `Lunar Phase Value`: The raw phase value (0.0 - 1.0).
- `Lunar Phase Name`: The binned phase category.

## Technical Details

- **Lunar Logic**: Powered by `github.com/dyindude/moonphase`, which implements standard astronomical algorithms.
- **Statistics**: Utilizes `gonum.org/v1/gonum` for precise Chi-squared distribution calculations.
- **System Access**: Leverages `golang.org/x/sys/unix` for low-level `statx` calls on Linux.

## License

MIT
