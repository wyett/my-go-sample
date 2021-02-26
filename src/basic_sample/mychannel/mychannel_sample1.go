// @author     : wyettLei
// @date       : Created in 2021/2/26 15:06
// @description: channel and pipeline

package mychannel

import "fmt"

// put arr into channel, and return channel
// v1
func putArr2Channel(input []int) chan int {
	size := len(input)
	c := make(chan int, size)
	for _, n := range input {
		c <- n
	}
	close(c)
	return c
}

// v2
func putArr2Channel2(input []int) <-chan int {
	size := len(input)
	c := make(chan int, size)
	go func() {
		for _, n := range input {
			c <- n
		}
		close(c)
	}()
	return c
}

// filter odd num
func filtOddNum(ch <-chan int) <-chan int {
	size := len(ch)
	c := make(chan int, size)
	go func() {
		for n := range ch {
			if n%2 == 0 {
				c <- n
			}
		}
		close(c)
	}()
	return c
}

// square
func squareChan(ch <-chan int) <-chan int {
	size := len(ch)
	c := make(chan int, size)
	go func() {
		for n := range ch {
			c <- n * n
		}
		close(c)
	}()
	return c
}

// sum
func sumChan(ch <-chan int) <-chan int {
	size := len(ch)
	out := make(chan int, size)
	go func() {
		sum := 0
		for n := range ch {
			sum += n
		}
		out <- sum
	}()
	return out
}

//-----------------------v1-------------------------/

func Pipeline1Main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range sumChan(squareChan(filtOddNum(putArr2Channel2(arr)))) {
		fmt.Println(n)
	}
}

//-----------------------v2-------------------------/

type a2c func([]int) <-chan int
type pf func(<-chan int) <-chan int

//var arr = []int{1,2,3,4,5,6,7,8,9}

func pipeInt(arr []int, func1 a2c, func2 ...pf) <-chan int {
	ch1 := func1(arr)
	for i := range func2 {
		ch1 = func2[i](ch1)
	}
	return ch1
}

func PipeLine2Main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range pipeInt(arr, putArr2Channel2, filtOddNum, squareChan, sumChan) {
		fmt.Println(n)
	}
}
