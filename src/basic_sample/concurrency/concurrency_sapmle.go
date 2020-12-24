package concurrency

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2020/12/23 15:11
 * @description: TODO
 */

func basicRuntimeInfoOps() {
	println("GOROOT: ", runtime.GOROOT())
	println("GOOS", runtime.GOOS)
	println("cpu: ", runtime.NumCPU())
	println(runtime.GOMAXPROCS(runtime.NumCPU()))

}

// Gosched
func goschedOps() {
	go func() {
		for i := 0; i < 5; i++ {
			println("goroutine...", i)
		}
	}()

	for i := 0; i < 4; i++ {
		//runtime.Gosched()
		println("goschedOps...", i)
	}
}

// runtime.Goexit
func goexitMain() {
	go func() {
		println("goexitMain begin...")

		goexitChild1()

		println("goexitMain end...")
	}()

	// sleep 3
	time.Sleep(3 * time.Second)
}

func goexitChild1() {
	println("goexitChild1 begin...")
	defer println("goexitChild1 defer...")
	// exit current goroutin, 不影响defer
	runtime.Goexit()
	println("goexitChild1 end...")
}

// sync
var ticket = 100
var s sync.WaitGroup  // 同步等待组
var lock sync.RWMutex // lock header

func syncMain() {
	// init
	s.Add(4) //设置同步等待组数量
	go saleTicket("窗口1")
	go saleTicket("窗口2")
	go saleTicket("窗口3")
	go saleTicket("窗口4")
	//表示让当前的goroutine等待，进入阻塞状态。一直到WaitGroup的计数器为零。才能解除阻塞
	s.Wait()
}

func saleTicket(name string) {
	rand.Seed(time.Now().UnixNano())
	//就是当WaitGroup同步等待组中的某个goroutine执行完毕后，设置这个WaitGroup的counter数值减1
	// 像latchdown
	defer s.Done()
	//for i := 1; i < ticket; i++ {
	//    println(name, "卖出了",i)
	//}

	// 加锁模式
	for {
		lock.Lock()
		if ticket > 0 {
			//加与不加差别很大
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Nanosecond)
			println(name, "卖出了", ticket)
			ticket--
		} else {
			lock.Unlock()
			println("票已售完")
			break
		}
		//一次售一张
		lock.Unlock()
	}
}

func ConcurrencyMain() {
	//basicRuntimeInfoOps()
	//goschedOps()
	//goexitMain()
	//syncMain()
	mutexMain()
}

// mutex address or copy
var x = 0

func incr(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

func mutexMain() {
	var wg sync.WaitGroup
	var m sync.Mutex
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go incr(&wg, &m)
	}
	wg.Wait()
	fmt.Println("mutex Main over ", x)
}
