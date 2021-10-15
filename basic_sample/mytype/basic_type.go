package mytype

import "fmt"

/**
 * @author     : wyettLei
 * @date       : Created in 2020/12/18 14:07
 * @description: TODO
 */

// array
func BasicArrayOps() {
	// 1.new
	arr1 := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("%v\n", arr1[i])
	}

	// 2....
	arr2 := [...]string{"aaa", "bbb", "ccc", "ddd"}
	fmt.Println("arr2 length: ", len(arr2))
	fmt.Printf("%T\n", arr2)

	// 3.modify
	arr2[0] = "eee"
	var arr3 [len(arr2)]string
	fmt.Println("arr3 is ", arr3)
	for id, value := range arr2 {
		arr3[id] = value
	}
	fmt.Println("arr2 is ", arr2)
	fmt.Println("arr3 is ", arr3)
}

func BasicSliceOps() {
	// 1.new
	slice1 := make([]int, 3, 5)
	fmt.Println(slice1)
	fmt.Println(len(slice1), cap(slice1))
	//fmt.PrintS

	// 2.slice with array
	arr1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := arr1[:]
	fmt.Println("whole slice2: ", slice2)
	slice3 := slice2[2:5]
	fmt.Println("whole slice3: ", slice3)

	// 3.modify slice
	for i, value := range slice3 {
		slice3[i] += value
	}
	fmt.Println("after:", slice3)
	fmt.Println("slice2 after modify:", slice2)
	fmt.Println("arr1 after modify:", arr1)

	// append() and copy()
	slice2 = append(slice2, 0)
	printSlice(slice2)

	slice3 = append(slice3, 2, 3, 4)
	printSlice(slice3)

	slice4 := make([]int, len(slice2), cap(slice2)*2)
	copy(slice4, slice2)
	printSlice(slice4)

}

func printSlice(arr []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(arr), cap(arr), arr)
}

func BasicMapOps() {
	// 1. create
	map1 := make(map[string]string)
	map2 := map[string]int{"a": 1, "b": 2, "c": 1}
	fmt.Println(map1)
	fmt.Println(len(map1))
	delete(map2, "c")
	fmt.Println(map2)
}

func BasicFuncOps() {
	// func arg

	// defer

}
