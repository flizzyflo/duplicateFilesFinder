package finder

import (
	"fmt"
	"time"
)

type PerformanceTracker struct {
	startTime          time.Time
	endTime            time.Time
	elapsedTime        time.Duration
	foldersVisited     int
	filesVisited       int
	hiddenFilesSkipped int
}

func (pt *PerformanceTracker) Start() {
	pt.startTime = time.Now()
}

func (pt *PerformanceTracker) AddFile() {
	pt.filesVisited++
}

func (pt *PerformanceTracker) AddFolder() {
	pt.foldersVisited++
}

func (pt *PerformanceTracker) AddHiddenFile() {
	pt.hiddenFilesSkipped++
}

func (pt *PerformanceTracker) Finish() {
	pt.endTime = time.Now()
	pt.elapsedTime = pt.endTime.Sub(pt.startTime)
}

func (pt *PerformanceTracker) Print() {
	fmt.Printf("-----------------------------------------------------------\n")
	fmt.Printf("[PERF] - Duration: %v\n", pt.elapsedTime)
	fmt.Printf("[PERF] - Folders checked: %v\n", pt.foldersVisited)
	fmt.Printf("[PERF] - Files checked: %v\n", pt.filesVisited)
	fmt.Printf("[PERF] - Hidden files skipped: %v\n", pt.hiddenFilesSkipped)
	fmt.Printf("-----------------------------------------------------------\n")
}
