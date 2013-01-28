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
func Crawl(url string, depth int, fetcher Fetcher) {
    // Fetch URLs in parallel.
    // Don't fetch the same URL twice.
    visited := struct {
        m   map[string]bool
        sync.Mutex
    }{m: make(map[string]bool)}
    var crawl func(url string, depth int)
    crawl = func(url string, depth int) {
        if depth <= 0 {
            return
        }
        visited.Lock()
        if visited.m[url] {
            fmt.Println("skipping", url)
            visited.Unlock()
            return
        }
        visited.m[url] = true
        visited.Unlock()
        body, urls, err := fetcher.Fetch(url)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("found: %s %q\n", url, body)
        done := make(chan bool)
        for _, u := range urls {
            go func(url string) {
                crawl(url, depth-1)
                done <- true
            }(u)
        }
        for _ = range urls {
            <-done
        }
    }
    crawl(url, depth)
}

func main() {
    s := time.Now()
    Crawl("http://golang.org/", 4, fetcher)
    fmt.Println(time.Now().Sub(s))
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
    if res, ok := (*f)[url]; ok {
        time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
        return res.body, res.urls, nil
    }
    return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
    "http://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "http://golang.org/pkg/",
            "http://golang.org/cmd/",
        },
    },
    "http://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "http://golang.org/",
            "http://golang.org/cmd/",
            "http://golang.org/pkg/fmt/",
            "http://golang.org/pkg/os/",
        },
    },
    "http://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
    "http://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
}
