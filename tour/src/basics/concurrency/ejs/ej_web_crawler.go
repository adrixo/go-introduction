package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cp *CrawledPlaces, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("\nCrawling ", url)
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	if cp.CrawlPlace(url) {
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			cp.SetBody(url, err.Error())

			return
		}

		fmt.Printf("found: %s %q\n", url, body)
		cp.SetBody(url, body)

		for _, u := range urls {
			wg.Add(1)
			go Crawl(u, depth-1, fetcher, cp, wg)
		}
	}

	return
}

func main() {
	crawledMap := make(map[string]string)
	cp := &CrawledPlaces{crawled: crawledMap}

	var wg sync.WaitGroup

	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher, cp, &wg)
	wg.Wait()
	fmt.Println("Completed!")

	// time.Sleep(1000 * time.Millisecond)
	fmt.Println(len(cp.crawled), ": ", cp.crawled)
}

type CrawledPlaces struct {
	mu      sync.Mutex
	crawled map[string]string
}

func (cp *CrawledPlaces) CrawlPlace(place string) (canBeCrawled bool) {
	cp.mu.Lock()
	canBeCrawled = false

	if _, exists := cp.crawled[place]; !exists {
		fmt.Println("\tCan be crawled: ", place)
		cp.crawled[place] = ""
		canBeCrawled = true
	} else {
		fmt.Println("\tWas crawled before: ", place)
	}
	cp.mu.Unlock()

	return canBeCrawled
}

func (cp *CrawledPlaces) SetBody(place string, body string) {
	cp.mu.Lock()
	cp.crawled[place] = body
	cp.mu.Unlock()

	return
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	random := rand.Intn(2000-1500) + 1500
	time.Sleep(time.Duration(random) * time.Millisecond)

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
