# go-kvstore

RestAPI Task1 Develop an in-memory key value store. The store must support data storage and retrieval through HTTP REST interfaces. The store must cater for arbitrary data types.

## Go packages
go install .\logging\ .\response\ .\router\ .\storage\ .\tcp\

## Go run
go run -race .\main.go

## Go Curl e.g.
Curl HTTP server

#### Get by key
curl --location --request GET 'http://localhost:9000/kvs/k0'

#### Create record on kvs
curl --location --request POST 'http://localhost:9000/kvs/k0_' \
--header 'Content-Type: application/json' \
--data-raw '{
    "fullname": "fullname5",
    "firstname": "firstname5",
    "middlename": "middlename5",
    "lastname": "lastname5",
    "email": "test9@test.com",
    "age": "55",
    "phone": "555555555",
    "addresses": [
        {
            "id": 1,
            "addressline1": "line1",
            "addressline2": "line2",
            "postcode": "HG1 1AA",
            "city": "Harrogate",
            "county": "York",
            "country": "UK"
        }
    ]
}'

#### Post image
curl --location --request POST 'http://localhost:9000/kvs/k0_' \
--header 'Content-Type: image/png' \
--data-binary '@<replacewithabsoluteimagepath>'

#### Update key record on kvs
curl --location --request PUT 'http://localhost:9000/kvs/k0' \
--header 'Content-Type: application/json' \
--data-raw '{
	"KUpdated": "Updated"
}'

#### Delete record by key on kvs
curl --location --request DELETE 'http://localhost:9000/kvs/k0_'

## TCP server
Start TCP server
go run client.go 127.0.0.1:9001

### TCP client commands
cmd=get | key=1 -> <!-- command get data by key --> \
cmd=create | key=1 | body=<data> -> <!-- command create record on KVS with key / body --> \
cmd=update | key=1 | body=<data> -> \
cmd=delete | key=1 -> \
cmd=stop -> <!-- command stop tells server to stop receiving connections from this client -->

### TCP using bash script
kvs.sh 

## Run test from src directory
go test ./...

### Detect test with race conditions
go test -race -count=1 client_test.go
go test -race .\storage\

## Profiling
Url: https://golang.org/pkg/runtime/pprof/
The following command runs benchmarks in the current directory and writes the CPU and memory profiles to cpu.prof and mem.prof:

### Run in root folder that has benchmark tests e.g. ´router´
go test -cpuprofile cpu.prof -memprofile mem.prof -benchmem -run=^$ router -bench .
go tool pprof .\cpu.prof
pprof will open write web to see svg

### Include come cup - mem code to be able to run
go run .\main.go -cpuprofile cpu.prof -memprofile mem.prof

### Benchmark Storage Create record
Maps aren't thread safe, when running multiple goroutines to create
records on the KVS store, race conditions errors show up

## Without Mutex
goos: windows
goarch: amd64
pkg: storage
BenchmarkParallelWrites-8   	 2722153	       398 ns/op	     136 B/op	       2 allocs/op
PASS
ok  	storage	1.650s
Success: Benchmarks passed.

## With Mutex
goos: windows
goarch: amd64
pkg: storage
BenchmarkParallelWrites-8   	 2300593	       475 ns/op	     157 B/op	       2 allocs/op
PASS
ok  	storage	1.745s
Success: Benchmarks passed.