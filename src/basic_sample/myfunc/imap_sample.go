package myfunc

import "fmt"

/**
 * @author     : wyettLei
 * @date       : Created in 2020/12/24 18:16
 * @description: TODO
 */

func myMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func SuperFuncMain() {
	arr := []int{1, 2, 3, 4, 5}
	r := myMap(arr, func(x int) int {
		return x * x
	})
	fmt.Println(r)
}
