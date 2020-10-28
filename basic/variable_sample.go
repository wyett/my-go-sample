package basic

import (
	"fmt"
	"os"
	"runtime"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2020/10/28 16:02
 * @description: TODO
 */

//全局变量希望能够被外部包所使用，则需要将首个单词的首字母也大写
// 变量可以编译期间就被赋值，赋值给全局变量使用运算符等号 =
// Go 是在编译时就已经完成推断过程, 省略类型也可以
// 变量的类型也可以在运行时实现自动推断 os.Getenv("HOME")
// 赋值给局部变量，使用 :=
// 字
// 值类型: int、float、bool 和 string, 数组，结构。 值类型的变量的值存储在栈中

// 引用类型：更复杂的数据通常会需要使用多个字，这些数据一般使用引用类型保存。
// 同一个引用类型的指针指向的多个字可以是在连续的内存地址中（内存布局是连续的），这也是计算效率最高的一种存储形式；
// 也可以将这些字分散存放在内存中，每个字都指示了下一个字所在的内存地址。
// 引用类型：maps, slices, channel, 指针, 被引用的变量会存储在堆中，以便进行垃圾回收，且比栈拥有更大的内存空间
var (
	aa int
	bb bool
	cc string
	dd float64
	ee *int
)

func VariableMain() {
	fmt.Print(aa, bb, cc, dd, ee)
	var goos string = runtime.GOOS
	fmt.Printf("OS is: %s\n", goos)
	var path string = os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)

	// 打印
	printTest()

	// 赋值
	assignment()

}

func printTest() {
	// fmt.Sprintf与Printf, 不过前者将格式化后的字符串以返回值的形式返回给调用者
	var1 := fmt.Sprintf("%s:%s", "100", "年")
	fmt.Println("==========================")
	fmt.Printf("%s:%s", "100", "年")
	fmt.Println("==========================")

	// fmt.Print与fmt.Println, 会自动使用格式化标识符 %v 对字符串进行格式化，两者都会在每个参数之间自动增加空格，而后者还会在字
	fmt.Print(var1, var1)
	fmt.Println("==========================")
	fmt.Println(var1, var1)
	fmt.Println("==========================")

	// assignment
}

func assignment() {
	// := 局部变量，省略var，声明后必须被使用，不可重复赋值, 支持同时赋值
	a := 1

	// = 全局变量或者局部变量， 没什么限制
	a = 1

	// 并行赋值，被用在返回多个值, 多变量必须被提前声明
	var (
		aaa int
		bbb float64
		ccc string
	)
	aaa, bbb, ccc = 1, 1.1, "abc"
	//aa, bb, cc = 1, 1.1, "abc"
	fmt.Println(a, b, c)

	// 局部变量多变量赋值，不需要被声明
	aaaa, bbbb, cccc := 1, 1.1, "abc"

	// 交换值
	var h1, h2 int
	h1, h2 = h2, h1

	// 只写变量_来丢弃值
	_, h1 = 5, 7

}
