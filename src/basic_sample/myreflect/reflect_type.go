package myreflect

import (
	"fmt"
	"reflect"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2020/12/25 9:06
 * @description: TODO
 */

type user struct {
	name string
}

type admin struct {
	user
	passwd string
}

func ReflectMain1() {
	var ad admin
	t := reflect.TypeOf(ad)

	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Printf("name==>%v, type==>%v, index==>%v, anonymous==>%v, offset==>%v, pkgPath==>%v, tag==>%v\n",
			f.Name, f.Type, f.Index, f.Anonymous, f.Offset, f.PkgPath, f.Tag)
	}

}

// 如果成员变量为指针类型，先用Elem获取其类型
func ReflectMain2() {
	var ad *admin

	t := reflect.TypeOf(ad)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Printf("Name==>%v, Type==>%v\n", f.Name, f.Type)
	}

}

//--------------------type interface and pointer interface---------------/

func (*user) toString() {}
func (admin) test()     {}

func ReflectMain3() {
	var ad admin

	methods := func(p reflect.Type) {
		for i, n := 0, p.NumMethod(); i < n; i++ {
			m := p.Method(i)
			//fmt.Printf("name==>%v, type==>%v, func==>%v\n", m.Name, m.Type, m.Func)
			fmt.Println(m.Name)
		}
	}

	fmt.Println("==============type interface===============")
	methods(reflect.TypeOf(ad))
	fmt.Println("==============pointer interface===============")
	methods(reflect.TypeOf(&ad))

}
