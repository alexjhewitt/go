package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{"http://google.com", "http://reddit.com", "http://golang.org", "http://amazon.com"}
	ch := make(chan string)
	for _, link := range links {
		go checkLink(link, ch)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-ch)
	}
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		ch <- "Might be down I think"
		return
	}
	fmt.Println(link, "is up!")
	ch <- "Yup it is up"
}
