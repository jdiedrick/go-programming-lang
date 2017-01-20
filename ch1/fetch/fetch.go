// fetch prints the content found at each specified url
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		protocol := "http://"
		if !strings.HasPrefix(url, protocol) {
			url = protocol + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:  %v\n", err)
			os.Exit(1)
		}

		/*
			b, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			}

			s := resp.Status

			fmt.Printf("Status:\n%s\n\nBody:\n%s\n", s, b)
		*/
		io.Copy(os.Stdout, resp.Body)
	}
}
