package database

import (
    "log"
    "time"

    "github.com/elastic/go-elasticsearch/v7"
)

func NewElasticsearchClient(elasticsearchURI string) *elasticsearch.Client {
    var es *elasticsearch.Client
    var err error
    retryCount := 5
    retryInterval := 5 * time.Second

    for i := 0; i < retryCount; i++ {
        cfg := elasticsearch.Config{
            Addresses: []string{elasticsearchURI},
        }
        es, err = elasticsearch.NewClient(cfg)
        if err == nil {
            res, err := es.Info()
            if err == nil {
                defer res.Body.Close()
                return es
            }
        }

        log.Printf("Failed to connect to Elasticsearch (%d/%d): %v. Retrying in %s...", i+1, retryCount, err, retryInterval)
        time.Sleep(retryInterval)
    }

    log.Fatalf("Failed to connect to Elasticsearch after %d attempts: %v", retryCount, err)
    return nil
}
