package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	port = fmt.Sprintf(":%s", port)

	chn := make(chan struct{}, 10)

	for i := 1; i > 0; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()

			chn <- struct{}{}

			defer func() {
				<-chn
			}()

			REQ(port)
		}()
	}

	wg.Wait()
}

func REQ(port string) {

	resp, err := http.Get("http://localhost")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
