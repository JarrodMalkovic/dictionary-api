package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"

    "dictionary-api/api/handlers"
    "dictionary-api/internal/config"
    "dictionary-api/internal/database"
)

func main() {
    cfg := config.LoadConfig()

    // Initialize MongoDB client
    mongoClient := database.NewMongoClient(cfg.MongoURI)
    handlers.Init(mongoClient)

    // Initialize Elasticsearch client
    esClient := database.NewElasticsearchClient(cfg.ElasticsearchURI)
    handlers.InitElastic(esClient)

    // Check Elasticsearch connection
    res, err := esClient.Info()
    if err != nil {
        log.Fatalf("Error getting response: %s", err)
    }
    defer res.Body.Close()
    fmt.Println(res)

    r := mux.NewRouter()
    r.HandleFunc("/words/{word}", handlers.GetWordHandler).Methods("GET")
    r.HandleFunc("/words/search", handlers.SearchWordsHandler).Methods("GET")

    fmt.Println("Server is running on port 8000")
    log.Fatal(http.ListenAndServe(":8000", r))
}
