package mytime

import (
	"fmt"
	"time"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2020/12/18 10:10
 * @description: TODO
 */

func BasicTimeOps() {
	// 1.年月日
	t1 := time.Now()
	fmt.Printf("%T\n", t1)
	fmt.Println(t1)

	// 2.Date
	//t2 := time.Date(2020, 12, 1,2,3,4,100, time.UTC)
	t2 := t1.Format("2006年1月2日 3时4分5秒")
	fmt.Printf("%T\n", t1)
	fmt.Println(t2)
}
