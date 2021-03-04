// @author     : wyettLei
// @date       : Created in 2021/3/4 16:30
// @description: TODO

package mysync

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// add 10000 times with mutex
func mutexFunc() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	x := int64(0)
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			x++
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("mutexFunc()=====>%d", x)

}

// add 10000 times with atomic
func atomicFunc() {
	var wg sync.WaitGroup
	x := int64(0)
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&x, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("mutexFunc()=====>%d", x)
}

func MAAMain() {
	go mutexFunc()
	go atomicFunc()

	time.Sleep(1000 * time.Millisecond)
}
