package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1000)
	for i:=0;i<1000;i++ {
		go sendRequest(i)
	}
	wg.Wait()
}

func sendRequest(threadNumber int) {
	defer wg.Done()
	url := "http://localhost:8080"
	response,err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strconv.Itoa(response.StatusCode))
}
