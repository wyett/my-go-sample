package main

/**
 * @author     : wyettLei
 * @date       : Created in 2020/10/28 19:05
 * @description: TODO
 */

import (
	"basic_sample/mycontroll"
	"basic_sample/myfunc"
)

func main() {
	// 1. hello
	//basic.HelloMain()

	//// 2. const
	//basic.ConstMain()

	//// 3.variable
	//basic.VariableMain()

	//// 4.execute_order
	//basic.ExecuteOrderMain()
	//basic.ExecuteOrderMain2()

	//// 5.main package func
	//printMain()

	/////////////////////////////////////////
	// time
	//mytime.BasicTimeOps()

	////////////////////////////////////////////
	// array
	//mytype.BasicArrayOps()
	// slice
	//mytype.BasicSliceOps()
	// map
	//mytype.BasicMapOps()
	// method
	//mytype.BasicMethodOps()

	//////////////////////////////////////////
	mycontroll.BasicControllOps()

	//fmt.Printf("======================================\n")
	//////////////////////////////////////////
	// recover()
	//exception.BasicRecoverOps()

	//////////////////////////////////////
	// file
	//fmt.Printf("======================================\n")
	//myio.BasicStatOps()
	//myio.BasicDirOps()

	//////////////////////////////////////
	//concurrency.ConcurrencyMain()
	//concurrency.ChanMain()
	//concurrency.WorkerPoolMain()

	//////////////////////////////////////
	myfunc.SuperFuncMain()

}
