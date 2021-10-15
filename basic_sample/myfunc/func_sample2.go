// @author     : wyettLei
// @date       : Created in 2021/4/15 9:03
// @description: TODO

package myfunc

import "fmt"

func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		fmt.Println("app inner:", &t)
		return t
	}
	fmt.Println("app outter:", &t)
	return c
}

func FuncSample2() {
	a := app()
	b := app()
	a("go")
	fmt.Println(b("All"))
}
