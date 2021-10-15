package basic

/**
 * @author     : wyettLei
 * @date       : Created in 2020/10/29 14:39
 * @description: TODO
 */

var a1 = "G"

func ExecuteOrderMain() {
	println("execute_order_sample.ExecuteOrderMain")
	n()
	m()
	n()
}

func n() {
	println("execute_order_sample.n()")
	println(a1)
}

func m() {
	println("execute_order_sample.m()")
	a1 := "O"
	println(a1)
}

////////////////////////////////////////
var a2 string

func ExecuteOrderMain2() {
	println("execute_order_sample.ExecuteOrderMain2")
	a2 = "G"
	println(a2)
	f1()
}

func f1() {
	println("execute_order_sample.f1()")
	a2 := "O"
	println(a2)
	f2()
}

func f2() {
	println("execute_order_sample.f2()")
	print(a2)
}
