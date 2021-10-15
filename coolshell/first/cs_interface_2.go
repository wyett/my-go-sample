package first

import "fmt"

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/2 14:52
 * @description: 结构体嵌入
 */

type withName struct {
	name string
}

type country2 struct {
	withName
}

type city2 struct {
	withName
}

type printable2 interface {
	printstr()
}

func (w withName) printStr() {
	fmt.Println(w.name)
}

func CsInterface2Main() {
	c1 := country2{withName{"China"}}
	c2 := city2{withName{"BeiJing"}}
	c1.printStr()
	c2.printStr()
}
