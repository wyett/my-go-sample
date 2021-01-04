package main

import (
	"flag"
	"fmt"
	"os"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2020/12/18 12:14
 * @description: TODO
 */

/**
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


*/

//-------------------------------------------------------

var (
	help        bool
	config_file string
)

func init() {
	flag.BoolVar(&help, "--help", false, "help")
	flag.StringVar(&config_file, "config", "conf/autopartition.conf", "auto partition config file")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `version: 1.0.0
Usage: auto_partition [-c filename] [-s stock]

Options:
`)
	flag.PrintDefaults()

}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}

	return

}
