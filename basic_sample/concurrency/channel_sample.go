package concurrency

import (
	"fmt"
	"time"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2020/12/23 17:52
 * @description: TODO
 */

// chan type and address

func basicChanOps() {
	var ch chan int
	ch = make(chan int)

	fmt.Printf("%T %p\n", ch, ch)

}

func ChanMain() {
	//basicChanOps()
	//chanRWOps()
	//chanRWOps2()
	//chanRWOpsMain3()
	chanRMain()
}

// chan read/write
// 当子goroutine执行完毕前，主goroutine会因为读取ch1中的数据而阻塞。从而保证了子goroutine会先执行完毕。这就消除了对时间的需求
func chanRWOps() {
	var ch chan bool
	ch = make(chan bool)
	fmt.Printf("%T\n", ch)
	fmt.Println("chanRWOps begin")

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("chanRWOps 执行 i ", i)
		}
		//循环执行完成后设置标记
		ch <- true
	}()

	data := <-ch
	fmt.Println("执行结束 data: ", data)
}

// chan2
func chanRWOps2() {
	ch1 := make(chan int)
	done := make(chan bool)
	go func() {
		fmt.Println("子gorountin begin...")
		//ch1 := <- data
		data := <-ch1
		//time.Sleep(3*time.Second)
		fmt.Println("data:", data)
		done <- true
	}()

	ch1 <- 101
	fmt.Println("主gorountin 执行")
	<-done
	fmt.Println("主gorountin结束")
}

// chan3
func chanRWOpsMain3() {
	ch1 := make(chan int)
	go chanRWOpsChild3(ch1)
	for {
		time.Sleep(100 * time.Nanosecond)
		data, ok := <-ch1
		if !ok {
			fmt.Println("已取完所有元素", ok)
			break
		}
		fmt.Println("已读取第元素", data)
	}
	fmt.Println("主程序结束")
}

func chanRWOpsChild3(ch chan int) {
	fmt.Println("子程序开始...")
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("i ", i, "已写入ch")
	}
	close(ch)
	fmt.Println("写入完成")
}

// read chan
func chanRMain() {
	ch := make(chan int)
	go saveChan(ch)

	for v := range ch {
		fmt.Println("读取数据", v)
	}
	fmt.Println("main chan")
}

func saveChan(ch chan int) {
	fmt.Println("child goroutine")
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Nanosecond)
		ch <- i
	}
	close(ch)
}
