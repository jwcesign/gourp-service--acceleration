package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var nextService = os.Getenv("NEXT_SERVICE")
var nowService = os.Getenv("NOW_SERVICE")

func handlerReq(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get requests from:", r.Host)
	if nextService == "" {
		w.Write([]byte(nowService))
		return
	}

	resp, err := http.Get("http://" + nextService)
	if err != nil {
		fmt.Println("Get next service error:", err)
		w.Write([]byte("error"))
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	respTotal := fmt.Sprintf("%s -> %s", nowService, string(body))
	w.Write([]byte(respTotal))
}

func main() {
	handler := http.HandlerFunc(handlerReq)
	http.Handle("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
