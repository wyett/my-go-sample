package mytype

import "fmt"

/**
 * @author     : wyettLei
 * @date       : Created in 2020/12/18 17:38
 * @description: TODO
 */

type Rectangle struct {
	width, height int
}

func (r *Rectangle) setVal() {
	r.height = 20
}

func BasicMethodOps() {
	p := Rectangle{1, 2}
	s := p
	p.setVal()
	fmt.Println(p.height, s.height)

	// inhert
	em1 := Employee{Human{"aaa", 30, "100000000"}, "ccc"}
	st2 := Student{Human{"bbb", 12, "100000001"}, "ddd"}
	st2.sayHi()
	em1.sayHi()

}

///////////////////////////////////inhert and rewrite

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	compony string
}

func (h *Human) sayHi() {
	fmt.Println("你好我是Human: ", h.name)
}

// rewrite sayHi
func (h *Employee) sayHi() {
	fmt.Println("你好我是Human: ", h.name, " and Compony is ", h.compony)
}
