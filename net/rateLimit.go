package net

import (
	"sync"
	"time"
)

var (
	ipLastRequest = make(map[string]time.Time)
	ipMu          sync.Mutex
)

func IsRateLimited(ip string) bool {
	ipMu.Lock()
	defer ipMu.Unlock()

	last, ok := ipLastRequest[ip]
	now := time.Now()
	if ok && now.Sub(last) < 10*time.Second {
		return true
	}
	ipLastRequest[ip] = now
	return false
}
