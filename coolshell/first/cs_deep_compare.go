package first

import (
	"fmt"
	"reflect"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/2 11:36
 * @description: TODO
 */

type data struct {
	name string
	age  int8
}

func Deep_Compare() {
	v1 := data{}
	v2 := data{}
	fmt.Println("v1==v2:", reflect.DeepEqual(v1, v2))

	m1 := map[string]string{"name": "wyett1", "age": "10"}
	m2 := map[string]string{"name": "wyett2", "age": "10"}
	fmt.Println("m1==m2", reflect.DeepEqual(m1, m2))
}
