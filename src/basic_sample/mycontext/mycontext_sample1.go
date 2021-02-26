// @author     : wyettLei
// @date       : Created in 2021/2/26 9:58
// @description: TODO

package mycontext

import "fmt"

//var c = make(chan int, 2)

func gofunc1(c chan int) {
	fmt.Println("go func 1")
	c <- 1
}

func MyContextSample1Main() {
	c := make(chan int, 2)

	go func() {
		fmt.Println("go func2")
		c <- 2
	}()

	go gofunc1(c)

	fmt.Println("main func")
	fmt.Println("read chan......")
	first := <-c
	fmt.Printf("read chan......%d", first)

	<-c

}

//---------------------select----------------------/

var stop chan struct{}

func MyContextSelectMain() {
	go func(stop *chan struct{}) {
		for {
			select {
			case <-*stop:
				fmt.Println("stop")
				return
			default:
				fmt.Println("default")
			}
		}
	}(&stop)
}
