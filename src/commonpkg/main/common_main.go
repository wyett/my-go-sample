package main

import (
	"commonpkg/utils"
	"fmt"
	"os"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/18 23:11
 * @description: common pkg main
 */

var err error

func ConfigParser() {
	configFile := "conf/autopartition.conf"

	// open configFile
	var file *os.File
	if file, err = os.Open(configFile); err != nil {
		utils.Crash(fmt.Sprintf("Configure file open failed. %v", err), -1)
	}
	defer file.Close()

	// load configFile
	configure := nimo

}

func main() {
}
