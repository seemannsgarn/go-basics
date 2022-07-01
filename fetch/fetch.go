// fetch output answer from url, like curl

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

		// check prefix
		prefix := "http://"
		if strings.HasPrefix(url, prefix) {
			continue
		} else {
			url = prefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// check status site
		// fmt.Println(resp.Status)

		// just print variable from request
		// b, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		// 	os.Exit(1)
		// }
		// fmt.Printf("%s", b)

		// Stdout output
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: stdout %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
