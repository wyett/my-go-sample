package fifth

import (
	"fmt"
	"strings"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/8 10:27
 * @description: TODO
 */

func mapStrToStr(arr []string, f func(string) string) []string {
	resArr := []string{}
	for _, item := range arr {
		resArr = append(resArr, f(item))
	}
	return resArr
}

func reduceStrToSize(arr []string, f func(string) int) int {
	size := 0
	for _, item := range arr {
		size += f(item)
	}
	return size
}

func filterStr(arr []string, f func(string) bool) []string {
	resArr := []string{}
	for _, item := range arr {
		if f(item) {
			resArr = append(resArr, item)
		}
	}
	return resArr
}

func CsMapRedFilter1Main() {
	var fruits = []string{"banana", "apple", "peach", "pair", "orange"}

	// map
	map1 := mapStrToStr(fruits, func(fruit string) string {
		return strings.ToUpper(fruit)
	})
	fmt.Printf("map==>%v\n", map1)

	red1 := reduceStrToSize(fruits, func(fruit string) int {
		return len(fruit)
	})
	fmt.Printf("reduce==>%v\n", red1)

	filter1 := filterStr(fruits, func(fruit string) bool {
		return len(fruit) > 5
	})
	fmt.Printf("filter==>%v\n", filter1)
}
