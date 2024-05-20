package dictionary

type DictionaryEntry struct {
    Word       string `json:"word"`
    Definition string `json:"definition"`
    APIKey     string `json:"api_key"`
}
