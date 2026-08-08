package finder

import "time"

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
