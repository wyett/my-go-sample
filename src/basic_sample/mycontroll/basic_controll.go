package mycontroll

import "fmt"

/**
 * @author     : wyettLei
 * @date       : Created in 2020/12/18 14:20
 * @description: TODO
 */

func BasicControllOps() {

	s := []int{1, 2, 3, 4, 5}

	for k, v := range s {
		fmt.Printf("k=%v,v=%v\n", k, v)
	}

}
