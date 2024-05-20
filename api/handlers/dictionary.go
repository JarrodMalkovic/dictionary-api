package handlers

import (
    "context"
    "encoding/json"
    "net/http"
    "time"

    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"

    "dictionary-api/internal/dictionary"
    "dictionary-api/internal/config"
)

var client *mongo.Client

func Init(client *mongo.Client) {
    client = client
}

func GetWordHandler(w http.ResponseWriter, r *http.Request) {
    apiKey := r.URL.Query().Get("api_key")
    if !checkAPIKey(apiKey) {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    vars := mux.Vars(r)
    word := vars["word"]

    collection := client.Database("dictionary").Collection("words")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var entry dictionary.DictionaryEntry
    err := collection.FindOne(ctx, bson.M{"word": word}).Decode(&entry)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            http.Error(w, "Word not found", http.StatusNotFound)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(entry)
}

func checkAPIKey(apiKey string) bool {
    return apiKey == config.LoadConfig().APIKey
}
