package main

import "commonpkg/flagsample"

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/18 23:11
 * @description: common pkg main
 */

func main() {
	//--------------------------config----------------------/
	flagsample.ParseAPConfig("E://mygit//my-go-sample//conf//autopartition.conf")

	//mysql.MySQLQuery()

}
