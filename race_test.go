package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

// what is race conditions ? access (Read OR Write) with several go routin with same variable
func TestDataRaceConditions(t *testing.T) {
	var state int32

	for i := 0; i < 10; i++ {
		go func(i int) {
			state += int32(i)
		}(i)
	}
}

// how fix it ?

// 1- mutexes : lock go routings for read or writing a variables (globally)
// in mutex if you have more than 1 state for lock and unlock we have a deadlock
func TestDataRaceConditionsFixByMutex(t *testing.T) {
	var state int32
	var mu sync.RWMutex //read and write lock
	//var mu sync.Mutex //only read lock
	for i := 0; i < 10; i++ {
		go func(i int) {
			mu.Lock()
			//mu.RLocker() // read lock only
			state += int32(i)
			mu.Unlock()
			//mu.RUnlock() // read unlock only
		}(i)
	}
}

// 2- atomic values :
func TestDataRaceConditionsFixByAtomic(t *testing.T) {
	var state int32

	//var state atomic.Value

	for i := 0; i < 10; i++ {
		go func(i int) {
			//state += int32(i)
			atomic.AddInt32(&state, int32(i))
		}(i)
	}
}
