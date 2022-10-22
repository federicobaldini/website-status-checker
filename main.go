package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

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
