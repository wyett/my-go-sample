package first

import "fmt"

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/3 10:48
 * @description: TODO
 */

// interface full check
// 弱验证
type shape interface {
	sides() int
	area() int
}

type square struct {
	len int
}

func (s *square) sides() int {
	return 4
}

func CsInterface4Main() {
	s := square{5}
	fmt.Printf("sides=%d\n", s.sides())
}
