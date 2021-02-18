package utils

import (
	"fmt"
	"os"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/18 23:23
 * @description: TODO
 */

type Exit struct {
	Code int
}

func Crash(msg string, errCode int) {
	fmt.Println(msg)
	panic(Exit{errCode})
}

func HandleExit() {
	if e := recover(); e != nil {
		if exit, ok := e.(Exit); ok == true {
			os.Exit(exit.Code)
		}
		panic(e)
	}
}
