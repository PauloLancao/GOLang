package du

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
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

// WalkDirRec func
func WalkDirRec(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done() // release 1 wait count
	for _, entry := range DirEntsRec(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go WalkDirRec(subdir, n, fileSizes)
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

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// DirEntsRec returns the entries of directory dir.
func DirEntsRec(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
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
