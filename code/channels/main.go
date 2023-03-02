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

	for _, link := range links {
		check(link)
	}

	if linksDown == 0 {
		fmt.Println("All links are up!")
	}
}

func check(link string) {
	_, error := http.Get(link)

	if error != nil {
		fmt.Println(link, "might be down")
		linksDown++
		return
	}

	fmt.Println(link, "is up")
}
