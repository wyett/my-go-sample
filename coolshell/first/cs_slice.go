package first

import (
	"bytes"
	"fmt"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/2 10:02
 * @description: TODO
 */

/**
slice结构有unsafe.pointer，实际是个struct类型，会导致共享内存
slice共享内存验证
[0 0 0 42 100]
[0 99 42]
[0 0 99 42 100]
*/
func SliceOne() {
	foo := make([]int, 5)
	foo[3] = 42
	foo[4] = 100

	fmt.Println(foo)

	bar := foo[1:4]
	bar[1] = 99

	fmt.Println(bar)
	fmt.Println(foo)

}

// append()导致slice共享内存终结
// append()再cap不够时会重新分配内存，cap够用则不会
func SliceTwo() {
	a := make([]int, 32)
	b := a[1:16]
	// cap够用则不会重新分配内存
	//b = append(b, 1)
	a = append(a, 1)
	a[2] = 10

	fmt.Println(b)
	fmt.Println(a)
}

//arr[:divPos]
//arr[:divPos:divPos], Full Slice Expression，其最后一个参数叫“Limited Capacity”
func SliceThree() {
	arr := []byte("AAAAAAA/BBBBBBB")
	//divPos := bytes.IndexAny(arr, "/")
	divPos := bytes.IndexByte(arr, '/')

	div1 := arr[:divPos]
	div2 := arr[divPos+1:]

	fmt.Println("%v,div1=>", string(div1))
	fmt.Println("div2=>", string(div2))

	div1 = append(div1, "aaaa"...)

	fmt.Println("div1=>", string(div1))
	fmt.Println("div2=>", string(div2))
	fmt.Println("whole=>", string(arr))

}
