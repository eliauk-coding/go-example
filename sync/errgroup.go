package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

func main() {
	startTime := time.Now()
	var g errgroup.Group
	var urls = []string{
		"https://www.hao123.com/?tn=48020221_37_hao_pg",
		"https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-sync-primitives/#errgroup",
		"https://www.hao123.com/?tn=48020221_37_hao_pg",
	}
	for i := range urls {
		url := urls[i]
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				fmt.Println(resp.Status)
				resp.Body.Close()
			}
			return err
		})
	}
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all urls.")
	} else {
		fmt.Println("Failed fetched all urls.")
	}
	duration := time.Now().Sub(startTime)
	fmt.Println("duration:", duration)
}
