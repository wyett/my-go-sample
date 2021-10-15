// @author     : wyettLei
// @date       : Created in 2021/3/1 11:27
// @description: TODO

package config_parser

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestReadConfig(t *testing.T) {
	//var res []string
	var file *os.File
	var err error
	if file, err = os.Open("conf/log_config.conf"); err != nil {
		errors.New("open log_config failed")
	}
	defer file.Close()

	res := readConfigFile(*bufio.NewReader(file))

	for _, elem := range res {
		fmt.Println(elem)
	}
}

func TestLoadLogConfig(t *testing.T) {
	var file *os.File
	var err error
	if file, err = os.Open("conf/log_config.conf"); err != nil {
		fmt.Errorf("open config %s failed", file.Name())
	}
	defer file.Close()

	cf := ConfigFile{file, ";"}
	lc := new(LogConfig)
	cf.load(lc)
	//for k, v := range lc {
	//    fmt.Printf("")
	//}

}
