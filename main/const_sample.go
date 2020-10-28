package main

import "fmt"

/**
 * @author     : wyettLei
 * @date       : Created in 2020/10/28 15:37
 * @description: TODO
 */

//并行复制
//枚举
//赋值一个常量时，之后没赋值的常量都会应用上一行的赋值表达式
//赋值两个常量，iota 只会增长一次，而不会因为使用了两次就增长两次
//使用iota和位运算实现资源状态的使用
//使用_忽略不需要的iota
//简单地讲，每遇到一次 const 关键字，iota 就重置为 0
const (
	a = iota
	b
	c
	d = 5
	e
)

func main()  {
	fmt.Print(a, b, c, d, e)

}