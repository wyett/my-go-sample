// @author     : wyettLei
// @date       : Created in 2021/3/1 17:41
// @description: TODO

package conf

import (
	"commonpkg/utils"
	"fmt"
	"os"
	"reflect"
)

func ParseConfigAndPrint(f string, s struct{}) {
	var err error
	var file *os.File

	if file, err = os.Open(f); err != nil {
		utils.Crash(fmt.Sprintf("Configure file open failed. %v", err), -1)
	}
	defer file.Close()

	// load configFile
	configuration := NewConfigFile(file)
	if err := configuration.Load(&s); err != nil {
		utils.Crash(fmt.Sprintf("load autopartition config failed %v", err), -1)
	}

	fmt.Printf("configuration=>%T\n", configuration)
	fmt.Printf("configuration=>%T\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)
	fmt.Printf("%+v\n", s)

	values := reflect.ValueOf(s)
	types := reflect.TypeOf(s)
	for i := 0; i < values.NumField(); i++ {
		fmt.Println(values.Field(i))
	}

	fmt.Println("=======================================")
	for i := 0; i < types.NumField(); i++ {
		fmt.Println(types.Field(i))
	}
}
