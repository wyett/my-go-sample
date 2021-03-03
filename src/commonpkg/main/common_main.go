package main

import (
	"commonpkg/conf"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/18 23:11
 * @description: common pkg main
 */

func main() {
	//--------------------------config----------------------/
	conf.ParseConfigAndPrint("E://mygit//my-go-sample//conf//autopartition.conf", conf.Options)

	//mysql.MySQLQuery()

}
