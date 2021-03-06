package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello, web!")
	_, err := writer.Write(message)
	fmt.Println("http.Request :", request.GetBody)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	sizes := make(chan int)

	go responseSize("https://example.com", sizes)
	go responseSize("https://golang.org", sizes)
	go responseSize("https://golang.org/doc", sizes)
	//time.Sleep(5 * time.Second)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)

	/*http.HandleFunc("/hello", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}*/
}

func responseSize(url string, channel chan int) {
	start := time.Now()
	fmt.Println("Getting :", url)
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	body, err1 := ioutil.ReadAll(response.Body)

	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Printf("Length of body for %s after GET : %d", url, len(body))
	fmt.Println()
	fmt.Println(fmt.Sprintf("%s took %s", url, time.Since(start)))
	channel <- len(body)
}
