// author: Riyan Christy
// riyan [at] linux.com
// source problem: https://go.dev/tour/concurrency/10
// (the problematic code can also be found at that page)

package main

import (
	"fmt"
	"sync"
	"time"
)

// TODO: Don't fetch the same URL twice.
// With that instruction from the Go authors,
// we create a type here that will keep track of whether the URL has been fetchedd.
// but using map in concurrent programming isn't safe,
// so we use the synx.Mutex data structure to ensure that
// only one goroutine can access the map at a time.
type SafeTracker struct {
	mu sync.Mutex
	v  map[string]bool
}

type Fetcher interface { 
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string, c SafeTracker) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, c SafeTracker) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// [DONE]
	// This implementation doesn't do either one of them
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url, c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, c)
	}
	time.Sleep(time.Second * 2)
	return
}

func main() {
	c := SafeTracker{v: make(map[string]bool)}
	Crawl("https://golang.org/", 4, fetcher, c)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

// Value() returns the bool, whether the URL has been fetched.
func (c *SafeTracker) Value(key string) bool {
	// we're accesing the map concurrently here, so
	// we have to do lock/unlock for safety.
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.

	defer c.mu.Unlock()
	// You want to unlock after this surrounding func returns.
	return c.v[key]
}

func (c *SafeTracker) Setter(key string) {
	c.mu.Lock()
	c.v[key] = true
	c.mu.Unlock()
}

func (f fakeFetcher) Fetch(url string, c SafeTracker) (string, []string, error) {
	if res, ok := f[url]; ok {
		switch tester := c.Value(res.body); tester {
		case true:
			return "", nil, fmt.Errorf("URL already found: %s", url)
		case false:
			c.Setter(res.body)
			return res.body, res.urls, nil
		}
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

// so if it is in the string of map[string], it will be found.
// But if it's in the []string list, it won't be foundd.

