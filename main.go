package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func getLinksByConsole() []string {
	links := []string{}
	link := ""
	stop := false
	response := ""

	for !stop {
		// get from the user the url link to check
		fmt.Println("Insert a link to check:")
		_, err := fmt.Scanln(&link)
		if err != nil {
			panic(err)
		}
		// check if the url is a valid one
		_, err = url.ParseRequestURI(link)
		if err != nil {
			fmt.Println("The inserted url: '", link, "' is not valid, make sure that it looks like this: 'https://www.example.com'")
		} else {
			links = append(links, link)
			// recap of the inserted urls
			fmt.Println()
			fmt.Println("You have inserted the current links:")
			for _, l := range links {
				fmt.Println(l)
			}
			fmt.Println("\nDo you want to insert one more link? Y/n")
			fmt.Println()
			// check if the user want add one more url or start the websites check
			_, err = fmt.Scanln(&response)
			if err != nil {
				panic(err)
			}
			if response == "No" || response == "N" || response == "n" {
				stop = true
			}
		}
	}

	return links
}

func main() {
	links := getLinksByConsole()
	channel := make(chan string)

	fmt.Println("\n___STARTING TO CHECK___")
	fmt.Println()

	for _, link := range links {
		// create a Goroutine for each link that check the website status
		go checkLink(link, channel)
	}

	// listen on the channel, when a checkLink() function got the website status it returns its link
	for linkFromPreviousCheck := range channel {
		// the returned link is used to re-check the website after 1 minute
		go func(link string) {
			time.Sleep(60 * time.Second)
			checkLink(link, channel)
		}(linkFromPreviousCheck)
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
