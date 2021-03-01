// @author     : wyettLei
// @date       : Created in 2021/3/1 17:41
// @description: TODO

package flagsample

import (
	config "commonpkg/conf"
	"commonpkg/utils"
	"fmt"
	"os"
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
		utils.Crash(fmt.Sprintf("load %v failed %v", &config.Options, err), -1)
	}

	//fmt.Printf(reflect.TypeOf(configuration))
	fmt.Printf("%v", configuration)

}
