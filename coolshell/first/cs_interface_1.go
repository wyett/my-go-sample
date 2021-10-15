package first

import "fmt"

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/2 14:44
 * @description: TODO
 */

type country struct {
	Name string
}

type city struct {
	Name string
}

type printable interface {
	printStr()
}

func (c country) printStr() {
	fmt.Println(c.Name)
}

func (c city) printStr() {
	fmt.Println(c.Name)
}

func CsInterface1Main() {
	c1 := country{"China"}
	c2 := city{"BeiJing"}
	c1.printStr()
	c2.printStr()
}
