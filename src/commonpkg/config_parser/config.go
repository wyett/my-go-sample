package config

import (
	"errors"
	"os"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/18 23:36
 * @description: config
 */

type ConfigLoader struct {
	file      *os.File
	separator string
}

func (loader *ConfigLoader) setSeparator(sep string) {
	loader.separator = sep
}

func (loader *ConfigLoader) Load(f interface{}) error {
	// 1.check f
	if f == nil {
		return errors.New("load f nil error")
	}

	// 2.

	// 3.readLine

}
