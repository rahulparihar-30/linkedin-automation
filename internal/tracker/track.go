package tracker

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

const HistoryFile = "history.json"

type HistoryEntry struct {
	ProfileURL string `json:"profile_url"`
	Status     string `json:"status"`
	Date       string `json:"date"`
}

type History map[string]HistoryEntry

var (
	data  History
	mutex sync.Mutex
)

func LoadHistory() {
	data = make(History)
	file, err := os.ReadFile(HistoryFile)
	if err == nil {
		json.Unmarshal(file, &data)
	}
}

func IsMessaged(url string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	if entry, exists := data[url]; exists {
		return entry.Status == "messaged"
	}
	return false
}

func MarkAsMessaged(url string) {
	mutex.Lock()
	defer mutex.Unlock()

	data[url] = HistoryEntry{
		ProfileURL: url,
		Status:     "messaged",
		Date:       time.Now().Format("2024-12-20"),
	}

	jsonData, _ := json.MarshalIndent(data, "", "  ")
	os.WriteFile(HistoryFile, jsonData, 0644)
}
