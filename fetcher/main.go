package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

func Contains[T comparable](a []T, b T) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}

var links = make([]string, 0, 10)
var linksMutex sync.Mutex
var visitedCounter = 0
var ch = make(chan bool)
var visitedMutex sync.Mutex

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	if !Contains(links, url) {
		links = append(links, url)
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		visitedMutex.Lock()
		visitedCounter++
		if visitedCounter == len(links) {
			ch <- true
		}
		visitedMutex.Unlock()
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	linksMutex.Lock()
	for _, u := range urls {
		if !Contains(links, u) {
			links = append(links, u)
			go Crawl(u, depth-1, fetcher)
		}
	}

	linksMutex.Unlock()
	visitedMutex.Lock()
	visitedCounter++
	if visitedCounter == len(links) {
		ch <- true
	}
	visitedMutex.Unlock()
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
	<- ch
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
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
