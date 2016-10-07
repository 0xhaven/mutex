package mutex

import (
	"sync"
	"testing"
)

const numThreads = 100

func TestMutex(t *testing.T) {
	var m Mutex
	data := make(map[rune]int)

	var wg sync.WaitGroup
	wg.Add(numThreads)
	for i := 0; i < numThreads; i++ {
		go func() {
			for r := 'A'; r <= 'Z'; r++ {
				m.Lock()
				data[r]++
				m.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	for r, count := range data {
		if count != numThreads {
			t.Fatalf("wrong count (%d!=%d) for %v", count, numThreads, string(r))
		}
	}
}
