package main

import (
	"fmt"
	"sync"
)

type urlMap struct {
	hmap map[string]bool
	mu   sync.Mutex
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, fetcher Fetcher, urlMap *urlMap) {
	urlMap.mu.Lock()
	if urlMap.hmap[url] {
		urlMap.mu.Unlock()
		return
	} else {
		urlMap.hmap[url] = true
	}
	urlMap.mu.Unlock()

	body, urls, err := fetcher.Fetch(url)
	var done sync.WaitGroup

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		done.Add(1)
		go func(u string) {
			defer done.Done()
			Crawl(u, fetcher, urlMap)
		}(u)
	}

	done.Wait()
	return
}

func main() {
	mapPtr := &urlMap{
		hmap: make(map[string]bool),
		mu:   sync.Mutex{},
	}
	Crawl("https://golang.org/", fetcher, mapPtr)
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
