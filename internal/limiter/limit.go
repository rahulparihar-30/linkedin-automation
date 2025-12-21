package limiter

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	LimitFile  = "limits.json"
	DailyLimit = 20
)

type Tracker struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

func CanSendRequest() bool {
	tracker := loadTracker()
	today := time.Now().Format("2006-01-02")

	if tracker.Date != today {
		tracker.Date = today
		tracker.Count = 0
		saveTracker(tracker)
		return true
	}

	if tracker.Count >= DailyLimit {
		fmt.Printf("⚠️ Daily limit reached (%d/%d). Stopping for today.\n", tracker.Count, DailyLimit)
		return false
	}

	fmt.Printf("✅ Daily Progress: %d/%d requests sent.\n", tracker.Count, DailyLimit)
	return true
}

func IncrementRequest() {
	tracker := loadTracker()
	today := time.Now().Format("2006-01-02")

	if tracker.Date != today {
		tracker.Date = today
		tracker.Count = 0
	}

	tracker.Count++
	saveTracker(tracker)
}

func loadTracker() Tracker {
	file, err := os.ReadFile(LimitFile)
	if err != nil {
		return Tracker{Date: time.Now().Format("2006-01-02"), Count: 0}
	}

	var tracker Tracker
	json.Unmarshal(file, &tracker)
	return tracker
}

func saveTracker(t Tracker) {
	data, _ := json.MarshalIndent(t, "", "  ")
	os.WriteFile(LimitFile, data, 0644)
}
