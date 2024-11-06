package main

import (
	"fmt"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elk-golang-integration/internal/log"
)

func main() {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	logData := log.Log{
		Timestamp: time.Now(),
		Level:     "INFO",
		Message:   "Log message sent directly to Elasticsearch",
	}

	err = log.SendLogToElasticsearch(es, logData)
	if err != nil {
		log.Fatalf("Error sending log to Elasticsearch: %s", err)
	}

	fmt.Println("Log sent successfully")
}
