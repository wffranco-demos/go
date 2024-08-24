package singleton

import (
	"sync"
	"testing"
)

func TestGetSingleDatabaseInstance(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetDatabaseInstance()
		}()
	}
	wg.Wait()
}
