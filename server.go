package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/syndtr/goleveldb/leveldb"
)

// CreateStreamRequest represents a request to create a stream
type CreateStreamRequest struct {
	StreamName string
	ShardCount int
}

type EnhancedMonitoringData struct {
	ShardLevelMetrics []string
}

type ShardData struct {
	ShardId string
}

type StreamData struct {
	RetentionPeriodHours    string
	EnhancedMonitoring      []EnhancedMonitoringData
	EncryptionType          string
	HasMoreShards           bool
	Shards                  []ShardData
	StreamARN               string
	StreamName              string
	StreamStatus            string
	StreamCreationTimestamp int
	_seqIx                  []int    // Hidden data, remove when returning
	_tags                   []string // Hidden data, remove when returning
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

	db := initDB()
	defer db.Close()

	stream := StreamData{StreamName: req.StreamName}
	streamBytes, _ := json.Marshal(&stream)

	err := db.Put([]byte(req.StreamName), streamBytes, nil)
	if err != nil {
		log.Fatal(err)
	}

	val, err := db.Get([]byte(req.StreamName), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(val))
}

func initDB() *leveldb.DB {
	db, err := leveldb.OpenFile("/tmp/kinesis-lite.db", nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	fmt.Println("The server is starting...")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":4567", nil))
}
