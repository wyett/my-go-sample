package fifth

import "reflect"

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/8 19:22
 * @description: https://coolshell.cn/articles/21164.html
 */

func MyMap(data interface{}, fn interface{}) []interface{} {
	vfn := reflect.ValueOf(fn)
	vdata := reflect.ValueOf(data)
	result := make([]interface{}, vdata.Len())

	for i := 0; i < vdata.Len(); i++ {
		result[i] = vfn.Call([]reflect.Value{vdata.Index(i)})[0].Interface()
	}

	return result
}
