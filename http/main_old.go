package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	//fmt.Println(resp)
	// This bite slices purpose is to hold all the data that gets written into it by that read function
	// bs := []byte{}

	// Create an empty byte slice with 999999 empty elements
	bs := make([]byte, 999999)

	// Take the byte slice and fill it with content
	resp.Body.Read(bs)

	// Now the byte slice is filled with the html response body data and we can try to print it
	fmt.Println(string(bs))
}
