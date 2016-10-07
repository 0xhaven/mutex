package mutex

import (
	// https://golang.org/pkg/sync/atomic
	"sync/atomic"
)

type Mutex struct {
}

// Lock returns when the mutual exclusion lock is held.
func (m *Mutex) Lock() {
	swapped := atomic.CompareAndSwapInt32(p, testValue, newValue)
}

// Unlock releases the mutual exclusion.
// It should throw an panic if the lock isn't held.
func (m *Mutex) Unlock() {
	if mutexNotHeld {
		panic("attempting to unlock already unlocked mutex")
	}

	// else, unlock mutex...

}
