package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateStreamRequest represents a request to create a stream
type CreateStreamRequest struct {
	StreamName string
	ShardCount int
}

func handler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Body: %s\n", b)
	// fmt.Fprintf(w, "%s", b)

	fmt.Printf("Method:%s \nURL:%s \nProto:%s \n", r.Method, r.URL, r.Proto)
	//Iterate over all header fields
	for k, v := range r.Header {
		fmt.Printf("Header field %q, Value %q\n", k, v)
	}

	fmt.Printf("Host = %q\n", r.Host)
	fmt.Printf("RemoteAddr= %q\n", r.RemoteAddr)
	//Get value for a specified token
	fmt.Printf("\n\nFinding value of \"Accept\" %q\n", r.Header["Accept"])

	createStream(b)
}

func createStream(data []byte) {
	var req CreateStreamRequest
	json.Unmarshal(data, &req)
	fmt.Printf("StreamName: %s\n", req.StreamName)
	fmt.Printf("ShardCount: %d\n", req.ShardCount)
}

func main() {
	fmt.Println("The server is starting...")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":4567", nil))
}
