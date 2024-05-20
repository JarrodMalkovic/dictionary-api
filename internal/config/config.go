package config

import (
    "os"
)

type Config struct {
    MongoURI         string
    ElasticsearchURI string
}

func LoadConfig() Config {
    return Config{
        MongoURI:         getEnv("MONGO_URI", "mongodb://localhost:27017"),
        ElasticsearchURI: getEnv("ELASTICSEARCH_URI", "http://elasticsearch:9200"),
    }
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
