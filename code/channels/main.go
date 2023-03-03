package main

import (
	"fmt"
	"net/http"
)

var linksDown int = 0

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

	for i := 0; i < len(links); i++ {
		<-c
	}

	if linksDown == 0 {
		fmt.Println("All links are up!")
	}
}

func check(link string, channel chan string) {
	_, error := http.Get(link)

	if error != nil {
		fmt.Println(link, "might be down")
		channel <- link
		linksDown++
		return
	}

	fmt.Println(link, "is up")
	channel <- link
}
