package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	client := &http.Client{}
	queue := make(chan int, 1000000)

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			queue <- i
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()

	for i := 0; i < 10000; i++ {
		go func(c chan int, client *http.Client) {
			for {
				i, ok := <-c
				if !ok {
					return
				}
				req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
				if err != nil {
					fmt.Println(fmt.Errorf("request error %d", i).Error())
					continue
				}
				resp, err := client.Do(req)
				if err != nil {
					fmt.Println(fmt.Errorf("resp error %d", i).Error())
					continue
				}
				fmt.Printf("Request %d %d \n", i, resp.StatusCode)
				resp.Body.Close()
				client.CloseIdleConnections()

			}
		}(queue, client)
	}

	for {
		time.Sleep(60 * time.Second)
	}
}
