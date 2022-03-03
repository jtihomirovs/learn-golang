package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// resp.Body implements reader interface
	// os.Stdout implenets writer interface. If we hover over os.Stdout we would see *File which implements writer interface
	// func Copy(dst Writer, src Reader)
	io.Copy(os.Stdout, resp.Body)
}
