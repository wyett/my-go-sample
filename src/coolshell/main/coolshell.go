package main

import (
	"coolshell/second"
	"sync"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/2 10:06
 * @description: TODO
 */

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	//----------------slice--------------------//
	//go first.SliceOne()
	//go first.SliceTwo()
	//go first.SliceThree()
	//go first.Deep_Compare()
	//go first.PersonPrint()

	//----------------interface--------------------//
	//go first.CsInterface1Main()
	//go first.CsInterface2Main()
	//go first.CsInterface3Main()
	//go first.CsInterface4Main()

	//----------------------err---------------------//
	go second.CsErr1Main()

	wg.Wait()
}
