package html

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			a := r.FindStringSubmatch(string(html))
			if len(a) > 0 {
				c <- r.FindStringSubmatch(string(html))[1]
			} else {
				c <- fmt.Sprintf("Title não enconrado -> url %s", url)
			}
		}(url)
	}
	return c
}
