package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	arg := os.Args[1]
	if arg == "server" {
		doServer()
	} else if arg == "client" {
		doClient()
	} else {
		fmt.Println("set command args `server` or `client`")
	}
}

func doServer() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8088", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	body := r.Body
	b, _ := ioutil.ReadAll(body)
	responseMessage := fmt.Sprintf("METHOD : %s\nBody: %+v\n", method, string(b))
	fmt.Println("message\n", responseMessage)
	w.WriteHeader(200)
	w.Write([]byte(responseMessage))
}

func doClient() {
	ctx := context.Background()
	body := []byte("body sentence.")
	request, _ := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:8088/", bytes.NewReader(body))
	response, _ := http.DefaultClient.Do(request)
	resBody, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("first response body : %v", string(resBody))
	time.Sleep(5 * time.Second)
	response, _ = http.DefaultClient.Do(request)
	resBody, _ = ioutil.ReadAll(response.Body)
	fmt.Printf("second response body : %v", string(resBody))
}
