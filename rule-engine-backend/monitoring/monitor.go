package monitoring

import (
	"log"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type SystemStats struct {
	CPUUsage    float64 `json:"cpuUsage"`
	MemoryUsage float64 `json:"memoryUsage"`
	Timestamp   int64   `json:"timestamp"`
}

type AlertThreshold struct {
	CPUThreshold    float64 `json:"cpuThreshold"`
	MemoryThreshold float64 `json:"memoryThreshold"`
}

var (
	currentStats SystemStats
	alertConfig  AlertThreshold
	statsMutex   sync.RWMutex
	alertMutex   sync.RWMutex
)

// StartMonitoring begins continuous monitoring of system resources
func StartMonitoring() {
	// Default alert thresholds
	alertConfig = AlertThreshold{
		CPUThreshold:    80.0, // 80% CPU usage
		MemoryThreshold: 80.0, // 80% memory usage
	}

	// Initialize with zeros
	currentStats = SystemStats{
		CPUUsage:    0,
		MemoryUsage: 0,
		Timestamp:   time.Now().Unix(),
	}

	// Start monitoring in background
	go monitorSystemResources()

	log.Println("System monitoring started")
}

// monitorSystemResources continuously updates system stats
func monitorSystemResources() {
	// Initial CPU measurement to establish baseline
	// The first call to cpu.Percent always returns 0, so we make an initial call
	_, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Printf("Error initializing CPU monitoring: %v", err)
	}

	for {
		updateSystemStats()
		checkAlerts()
		time.Sleep(1 * time.Second)
	}
}

// updateSystemStats collects current CPU and memory usage
func updateSystemStats() {
	// Get CPU usage - use a shorter interval for more responsive updates
	// The second parameter (false) means we get total CPU usage, not per-core
	cpuPercent, err := cpu.Percent(500*time.Millisecond, false)
	if err != nil {
		log.Printf("Error getting CPU stats: %v", err)
		return
	}

	// Get memory usage
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("Error getting memory stats: %v", err)
		return
	}

	// Update the current stats
	statsMutex.Lock()
	if len(cpuPercent) > 0 {
		currentStats.CPUUsage = cpuPercent[0]
	}
	currentStats.MemoryUsage = memInfo.UsedPercent
	currentStats.Timestamp = time.Now().Unix()
	statsMutex.Unlock()
}

// checkAlerts checks if current usage exceeds thresholds
func checkAlerts() {
	statsMutex.RLock()
	stats := currentStats
	statsMutex.RUnlock()

	alertMutex.RLock()
	thresholds := alertConfig
	alertMutex.RUnlock()

	if stats.CPUUsage > thresholds.CPUThreshold {
		log.Printf("ALERT: CPU usage (%.2f%%) exceeds threshold (%.2f%%)", stats.CPUUsage, thresholds.CPUThreshold)
		// Here you could trigger additional alert mechanisms (email, SMS, etc.)
	}

	if stats.MemoryUsage > thresholds.MemoryThreshold {
		log.Printf("ALERT: Memory usage (%.2f%%) exceeds threshold (%.2f%%)", stats.MemoryUsage, thresholds.MemoryThreshold)
		// Here you could trigger additional alert mechanisms (email, SMS, etc.)
	}
}

// GetCurrentStats returns the most recent system stats
func GetCurrentStats() SystemStats {
	statsMutex.RLock()
	defer statsMutex.RUnlock()
	return currentStats
}

// SetAlertThresholds updates the alert thresholds
func SetAlertThresholds(thresholds AlertThreshold) {
	alertMutex.Lock()
	defer alertMutex.Unlock()
	alertConfig = thresholds
	log.Printf("Alert thresholds updated: CPU=%.2f%%, Memory=%.2f%%", thresholds.CPUThreshold, thresholds.MemoryThreshold)
}

// GetAlertThresholds returns the current alert thresholds
func GetAlertThresholds() AlertThreshold {
	alertMutex.RLock()
	defer alertMutex.RUnlock()
	return alertConfig
}
