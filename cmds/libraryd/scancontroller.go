package main

import "sync"

var scans map[int64]*Scan
var scansSync sync.RWMutex

type Scan struct {
	id      int64
	channel chan struct{}
}

// Adds a scan to the scan pool and waits for it to finish.
// Once it's finished it's removed from the pool.
func AddScanToPoolAndWatch(scan *Scan) {
	scansSync.Lock()
	defer scansSync.Unlock()

	if scans == nil {
		scans = make(map[int64]*Scan)
	}
	scans[scan.id] = scan

	go waitForScanToFinishAndRemove(scan)
}

// Removes a scan from the scan pool.
func RemoveScanFromPool(scan *Scan) {
	scansSync.Lock()
	defer scansSync.Unlock()
	delete(scans, scan.id)
}

// Returns true if any active scans are in the pool.
func ActiveScans() bool {
	scansSync.RLock()
	defer scansSync.RUnlock()
	return len(scans) > 0
}

func waitForScanToFinishAndRemove(scan *Scan) {
	<-scan.channel
	RemoveScanFromPool(scan)
}
