package scanner

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"time"

	"golang.org/x/sys/unix"
)

// FileInfo holds the relevant data for a file
type FileInfo struct {
	Path    string
	Created time.Time
}

// Scan recursively walks the directory and extracts creation dates
func Scan(root string) ([]FileInfo, error) {
	var files []FileInfo

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		created, err := getBirthTime(path)
		if err != nil {
			// Fallback to ModTime if birthtime is not available
			info, err := d.Info()
			if err != nil {
				return nil // Skip files we can't stat
			}
			created = info.ModTime()
		}

		files = append(files, FileInfo{
			Path:    path,
			Created: created,
		})

		return nil
	})

	return files, err
}

func getBirthTime(path string) (time.Time, error) {
	var statx unix.Statx_t

	// AT_FDCWD: Use current working directory for relative paths
	// AT_SYMLINK_NOFOLLOW: Don't follow symlinks
	err := unix.Statx(unix.AT_FDCWD, path, unix.AT_SYMLINK_NOFOLLOW, unix.STATX_BTIME, &statx)
	if err != nil {
		return time.Time{}, err
	}

	if statx.Mask&unix.STATX_BTIME == 0 {
		return time.Time{}, fmt.Errorf("birthtime not supported")
	}

	return time.Unix(statx.Btime.Sec, int64(statx.Btime.Nsec)), nil
}
