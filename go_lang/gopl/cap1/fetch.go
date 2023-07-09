package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func fetch() {
	for _, url := range os.Args[1:] {
		prefix := "http://"
		hasPrefix := strings.HasPrefix(url, prefix)
		if !hasPrefix {
			url = prefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		defer resp.Body.Close()
		// b, err := io.ReadAll(resp.Body)
		// resp.Body.Close()
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		// 	os.Exit(1)
		// }

		fmt.Printf("Status code: %d\n", resp.StatusCode)
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		}

		// fmt.Printf("RESULT: %v", os.Stdout)
	}
}
