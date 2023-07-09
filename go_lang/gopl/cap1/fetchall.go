package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func fetchallstart() {
	start := time.Now()
	ch := make(chan string)
	cache := make(map[string]int64)

	for _, url := range os.Args[1:] {
		go fetchall(url, ch, cache)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	for _, url := range os.Args[1:] {
		go fetchall(url, ch, cache)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchall(url string, ch chan<- string, cache map[string]int64) {
	start := time.Now()
	res, exist := cache[url]

	if exist {
	  secs := time.Since(start).Seconds()
	  ch <- fmt.Sprintf("%.2fs %7d %s", secs, res, url)
	  return
	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	defer resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	cache[url] = nbytes
	ch <- fmt.Sprintf("%.2fs %7d %s status: %d", secs, nbytes, url, resp.StatusCode)
}
