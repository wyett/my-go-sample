// @author     : wyettLei
// @date       : Created in 2021/3/1 11:27
// @description: TODO

package config

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	DefaultFormat    = "2005-01-02 15:04:05"
	DefaultSeperator = "="
)

type ConfigFile struct {
	config     *os.File
	dateFormat string
}

// new ConfigFile
func NewConfigFile(file *os.File) *ConfigFile {
	return &ConfigFile{file, DefaultFormat}
}

func (cf *ConfigFile) SetDateFormat(df string) {
	cf.dateFormat = df
}

// load configFile to config struct
func (cf *ConfigFile) Load(config interface{}) error {
	// 1.判断config不为空
	if config == nil {
		return errors.New("config is none")
	}

	// 2.获取config的 可用且为struct
	elements := reflect.ValueOf(config).Elem()
	if !elements.IsValid() || elements.Kind() != reflect.Struct {
		return errors.New("target config is not valid struct")
	}

	// 3.读取config
	configs := readConfigFile(*bufio.NewReader(cf.config))
	if len(configs) == 0 {
		return errors.New("config is none")
	}

	// 4.parse config
	for _, line := range configs {
		linePair := strings.SplitN(line, "=", 2)
		//fmt.Printf("left=%s,right=%s\n", linePair[0], linePair[1])
		//key, value := mapStringArr(linePair, strings.Trim(, " "))
		key := strings.Trim(linePair[0], " ")
		value := strings.Trim(linePair[0], " ")

		var fieldName string
		var err error
		// 根据key获取struct的fieldName
		if fieldName, err = getFieldByTag(config, key); err != nil {
			return errors.New("no fitted field of struct")
		}

		// get field by name
		field := elements.FieldByName(fieldName)
		if !field.IsValid() || !field.CanSet() {
			return fmt.Errorf("%s could not be used", field)
		}

		//
		switch field.Kind() {
		case reflect.String:
			field.SetString(value)
		case reflect.Bool:
			if v, err := strconv.ParseBool(value); err != nil {
				return fmt.Errorf("bool key %s has wrong value %s", key, value)
			} else {
				field.SetBool(v)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if v, err := strconv.ParseInt(value, 10, 64); err != nil {
				field.SetInt(v)
			} else {
				return fmt.Errorf("int key %s has wrong value %s", key, value)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if v, err := strconv.ParseUint(value, 10, 64); err != nil {
				return fmt.Errorf("Uint key %s has wrong value %s", key, value)
			} else {
				field.SetUint(v)
			}
		}
	}
	return nil
}

// read config file
func readConfigFile(file bufio.Reader) (res []string) {
	for {
		if buf, end, err := file.ReadLine(); end || err != nil {
			break
		} else {
			line := strings.Trim(string(buf), "\n")
			if line != "" && !strings.HasPrefix(line, "#") {
				res = append(res, line)
			}
		}
	}
	return res
}

// get filedName of struct by key name
func getFieldByTag(cs interface{}, tag string) (fieldName string, err error) {
	fieldCounts := reflect.ValueOf(cs).Elem().NumField()
	types := reflect.TypeOf(cs)
	for i := 0; i < fieldCounts; i++ {
		if types.Elem().Field(i).Tag.Get("config") == tag {
			return types.Elem().Field(i).Name, nil
		}
	}
	//return "", errors.New("no fitted element of struct")
	return "", fmt.Errorf("no fitted element of struct %s", tag)
}

func mapStringArr(arr []string, f func(...string) string) []string {
	var res []string
	for _, item := range arr {
		res = append(res, f(item))
	}
	return res
}
