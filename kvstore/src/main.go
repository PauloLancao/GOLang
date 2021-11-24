package main

import (
	"flag"
	"logging"
	"os"
	"router"
	"runtime"
	"runtime/pprof"
	"storage"
	"tcp"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	flag.Parse()
	cpuProfile()
	memProfile()

	// Start logging
	logging.CreateLogger(os.Args[1:])
	defer logging.Close()

	// Start storage
	store := storage.Start()
	defer store.Close()

	done := make(chan bool, 2)

	go router.Start(store, done)
	go tcp.Start(store, done)

	<-done
}

func cpuProfile() {
	if *cpuprofile != "" {
		uuid := logging.UUID()
		f, err := os.Create(*cpuprofile)
		if err != nil {
			logging.Fatalf(uuid, "cpuProfile", "cpuProfile", "could not create CPU profile: %+v", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			logging.Fatalf(uuid, "cpuProfile", "cpuProfile", "could not start CPU profile: %+v", err)
		}
		defer pprof.StopCPUProfile()
	}
}

func memProfile() {
	if *memprofile != "" {
		uuid := logging.UUID()
		f, err := os.Create(*memprofile)
		if err != nil {
			logging.Fatalf(uuid, "memProfile", "memProfile", "could not create memory profile: %+v", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			logging.Fatalf(uuid, "memProfile", "memProfile", "could not write memory profile: %+v", err)
		}
	}
}
