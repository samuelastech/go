package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://amazon.com",
		"https://youtube.com",
		"https://golang.org",
		"https://facebook.com",
		"https://tophonetics.com",
	}

	c := make(chan string)

	for _, link := range links {
		go check(link, c)
	}

	for link := range c {
		go check(link, c)
	}
}

func check(link string, channel chan string) {
	_, error := http.Get(link)

	if error != nil {
		fmt.Println(link, "might be down")
		channel <- link
		return
	}

	fmt.Println(link, "is up")
	channel <- link
}
