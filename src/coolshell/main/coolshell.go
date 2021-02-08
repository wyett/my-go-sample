package main

import (
	"coolshell/fifth"
	"sync"
	"time"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/2 10:06
 * @description: TODO
 */

//var WG sync.WaitGroup
var x sync.Mutex

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//WG.Add(2)

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
	//go second.CsErr1Main()
	//go second.CsErr4Main()

	//---------------------function option----------------------//
	//go third.CsBuilder1Main()
	//go third.CsFuncOption2Main()

	//---------------------map&&reduce&&filter----------------------//
	go fifth.CsMapRedFilter1Main()

	time.Sleep(20 * time.Millisecond)
	//WG.Wait()
}
