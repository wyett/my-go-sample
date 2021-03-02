// @author     : wyettLei
// @date       : Created in 2021/3/1 17:41
// @description: TODO

package config

import (
	"commonpkg/utils"
	"fmt"
	"os"
	"reflect"
)

func ParseAPConfig(f string) {
	//configFile := "conf/autopartition.conf"

	// open configFile
	var err error
	var file *os.File
	if file, err = os.Open(f); err != nil {
		utils.Crash(fmt.Sprintf("Configure file open failed. %v", err), -1)
	}
	defer file.Close()

	// load configFile
	configuration := config.NewConfigFile(file)
	if err := configuration.Load(&config.Options); err != nil {
		utils.Crash(fmt.Sprintf("load autopartition config failed %v", err), -1)
	}

	//fmt.Printf(reflect.TypeOf(configuration))
	fmt.Printf("configuration=>%T\n", configuration)
	fmt.Printf("configuration=>%T\n", config.Options)
	fmt.Printf("%v\n", config.Options)
	fmt.Printf("%#v\n", config.Options)
	fmt.Printf("%+v\n", config.Options)

	values := reflect.ValueOf(config.Options)
	types := reflect.TypeOf(config.Options)
	for i := 0; i < values.NumField(); i++ {
		fmt.Println(values.Field(i))
	}

	fmt.Println("=======================================")
	for i := 0; i < types.NumField(); i++ {
		fmt.Println(types.Field(i))
	}
	//for v := range config.Options {
	//    fmt.Println(v)
	//}
}
