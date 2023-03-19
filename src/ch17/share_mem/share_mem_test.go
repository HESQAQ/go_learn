package share_mem

import (
	"sync"
	"testing"
	"time"
)

// counter 在不同线程中自增, 导致并发竞争, 需要对共享的内存进行保护
func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter: %d", counter)
}

// 对 counter 加锁
func TestCounterSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	// 对 main 线程延迟执行,等协程先执行完
	time.Sleep(1 * time.Second)
	t.Logf("counter: %d", counter)
}

// 对 counter 加锁  waitgroup 替换 sleep
func TestCounterSafe2(t *testing.T) {
	var mut sync.Mutex
	//var rw sync.RWMutex  读写专用锁
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	// 对 main 线程延迟执行,等协程先执行完
	//time.Sleep(1 * time.Second)
	wg.Wait()

	t.Logf("counter: %d", counter)
}
