package first

import "fmt"

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/2 15:20
 * @description: 我们使用了一个叫Stringable 的接口，我们用这个接口把“业务类型”
 *               Country 和 City 和“控制逻辑” Print() 给解耦了。于是，只要实现了Stringable 接口，都可以传给 PrintStr() 来使用。
 *               Program to an interface not an implementation
 */

type country3 struct {
	name string
}

type city3 struct {
	name string
}

type printable3 interface {
	toString() string
}

func (c country3) toString() string {
	return c.name
}

func (c city3) toString() string {
	return c.name
}

func printStr3(p printable3) {
	fmt.Println(p.toString())
}

func CsInterface3Main() {
	c1 := country3{"China"}
	c2 := city3{"BeiJing"}
	printStr3(c1)
	printStr3(c2)

}
