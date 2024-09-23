package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{"http://google.com", "http://reddit.com", "http://golang.org", "http://amazon.com"}
	ch := make(chan string)
	for _, link := range links {
		go checkLink(link, ch)
	}

	for {
		go checkLink(<-ch, ch)
		time.Sleep(time.Second * 5)
	}
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		ch <- link
		return
	}
	fmt.Println(link, "is up!")
	ch <- link
}
