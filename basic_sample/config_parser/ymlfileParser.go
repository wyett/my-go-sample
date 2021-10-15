package config_parser

import (
	"fmt"
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2020/12/25 14:44
 * @description: test yml config parser
 */

type FirstChild struct {
	Aaaa string
}

type SecondChild struct {
	Aaaa int
	Bbbb int
}

type ThridChild struct {
	dddd int
}

func YmlConfigParser(c *FirstChild) {
	//testConfig := SecondChild{}
	data, err := ioutil.ReadFile("E:\\mygit\\my-go-sample\\src\\basic_sample\\config_parser\\config.yml")
	if err != nil {
		fmt.Println("文件不存在")
		return
	}
	fmt.Println("read file from config.yml", data)

	//
	e := yaml.Unmarshal(data, &c)
	if e != nil {
		fmt.Println("parser First failed")
		return
	}
	fmt.Println("c: ", &c)

	d, _ := yaml.Marshal(c)
	fmt.Println(string(d))

}

func MapParser() {
	data, err := ioutil.ReadFile("E:\\mygit\\my-go-sample\\src\\basic_sample\\config_parser\\config.yml")
	if err != nil {
		fmt.Println("文件不存在")
		return
	}
	fmt.Println("read file from config.yml", data)

	m := make(map[interface{}]interface{})
	e := yaml.Unmarshal(data, &m)
	if e != nil {
		fmt.Println("parser SecondChild failed")
		return
	}
	fmt.Println("m:", m)

	d, _ := yaml.Marshal(m)
	fmt.Println(string(d))

}
