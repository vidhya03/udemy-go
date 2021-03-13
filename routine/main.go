package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"https://facebook.com",
		"https://golang.org",
	}

	C := make(chan string)

	for _, link := range links {
		go checkLink(link, C)
	}

	for l := range C {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, C)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		c <- link
		return
	}

	fmt.Println(link, " is up!")
	c <- link
}
