package main

import (
	"sync"
	"time"
)

// * mutex -> mutual exclusion
// maps are  not thread-safe
// ! no one can read the mutex if it's being written to

type safeCounter struct {
	counts map[string]int
	mux    *sync.RWMutex
}

func (sc safeCounter) inc(key string) {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mux.RLock()
	defer sc.mux.RUnlock()
	return sc.counts[key]
}

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}
