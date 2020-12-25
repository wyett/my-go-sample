package main

import "flag"

/**
 * @author     : wyettLei
 * @date       : Created in 2020/12/18 12:14
 * @description: TODO
 */

type Exit struct {
	Code int
}

func main() {
	var err error
	defer handleExit()

	configuration := flag.String("config", "", "configration file")

	if *configuration == "" {

	}

}

func handleExit() {
	if e := recover(); e == nil {

	}
}
