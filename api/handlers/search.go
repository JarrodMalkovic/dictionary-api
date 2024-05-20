package handlers

import (
    "context"
    "encoding/json"
    "net/http"
    "bytes"

    "github.com/elastic/go-elasticsearch/v7"
)

var es *elasticsearch.Client

func InitElastic(client *elasticsearch.Client) {
    es = client
}

func SearchWordsHandler(w http.ResponseWriter, r *http.Request) {
    apiKey := r.URL.Query().Get("api_key")
    if !checkAPIKey(apiKey) {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    query := r.URL.Query().Get("q")
    if query == "" {
        http.Error(w, "Missing query parameter 'q'", http.StatusBadRequest)
        return
    }

    var buf map[string]interface{}
    searchReq := map[string]interface{}{
        "query": map[string]interface{}{
            "match": map[string]interface{}{
                "word": query,
            },
        },
    }
    jsonStr, err := json.Marshal(searchReq)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    res, err := es.Search(
        es.Search.WithContext(context.Background()),
        es.Search.WithIndex("dictionary"),
        es.Search.WithBody(bytes.NewReader(jsonStr)),
        es.Search.WithTrackTotalHits(true),
        es.Search.WithPretty(),
    )
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer res.Body.Close()

    if err := json.NewDecoder(res.Body).Decode(&buf); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(buf)
}
