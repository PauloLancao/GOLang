package main

import (
	"ch8/du"
	"flag"
)

// go run .\du1.go C:\Users\paulo.lancao\Desktop\BJSS-Training\GOLang
func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree.
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			du.WalkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// Print the results.
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	du.PrintDiskUsage(nfiles, nbytes)
}
