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
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// Loop through channels
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// Wait for the channel to return some value
	// After chanel return some value, assign it to "l" (link)
	// If we use range on slice, we take element from the slice and assign it to l
	for l := range c {
		// GO function literal aka anonymous function aka Lambda function
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
			// "}()" -> here is the function, call it!
		}(l)
	}

}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		//c <- "Might be down I think"
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	//c <- "Yep its up!"
	c <- link
}
