package config

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2021/1/4 11:29
 * @description: test configure
 */

var (
	help        bool
	config_file string
)

func init() {
	flag.BoolVar(&help, "--help", false, "help")
	flag.StringVar(&config_file, "config", "autopartition.conf", "auto partition config file")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `version: 1.0.0
Usage: auto_partition [-c filename] [-s stock]

Options:
`)
	flag.PrintDefaults()

}

func TestReadConfig(t *testing.T) {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}

	return

}
