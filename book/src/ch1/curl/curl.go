package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// go run .\curl.go http://gopl.io http://bad.gopl.io
func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// Two ways of doing it
		// b, err := ioutil.ReadAll(resp.Body)
		b, err := io.Copy(os.Stdout, resp.Body)

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Println("Http response status: ", resp.Status)
		fmt.Printf("%s", b)
	}
}
