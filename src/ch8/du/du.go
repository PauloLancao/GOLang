package du

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// WalkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func WalkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range DirEnts(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			WalkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// DirEnts returns the entries of directory dir.
func DirEnts(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

// PrintDiskUsage func
func PrintDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
