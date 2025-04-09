package monitoring

import (
	"encoding/json"
	"net/http"
)

// GetStatsHandler returns the current system stats
func GetStatsHandler(w http.ResponseWriter, r *http.Request) {
	stats := GetCurrentStats()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

// GetThresholdsHandler returns the current alert thresholds
func GetThresholdsHandler(w http.ResponseWriter, r *http.Request) {
	thresholds := GetAlertThresholds()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(thresholds)
}

// SetThresholdsHandler updates the alert thresholds
func SetThresholdsHandler(w http.ResponseWriter, r *http.Request) {
	var thresholds AlertThreshold
	if err := json.NewDecoder(r.Body).Decode(&thresholds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	// Validate thresholds
	if thresholds.CPUThreshold < 0 || thresholds.CPUThreshold > 100 ||
	   thresholds.MemoryThreshold < 0 || thresholds.MemoryThreshold > 100 {
		http.Error(w, "Thresholds must be between 0 and 100", http.StatusBadRequest)
		return
	}
	
	SetAlertThresholds(thresholds)
	w.WriteHeader(http.StatusOK)
}
