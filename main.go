package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://goland.org",
		"http://amazon.com",
	}
	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)
	}

	for l := range channel {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, channel)
		}(l)
	}
}

func checkLink(link string, channel chan string) {
	start := time.Now()
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down! response time:", time.Since(start))
		channel <- link
		return
	}
	fmt.Println(link, "is up! response time:", time.Since(start))
	channel <- link
}
