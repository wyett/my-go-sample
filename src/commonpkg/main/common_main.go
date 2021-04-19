package main

import "commonpkg/mysocket"

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/18 23:11
 * @description: common pkg main
 */

func main() {
	//--------------------------config----------------------/
	//conf.ParseConfigAndPrint("E://mygit//my-go-sample//conf//autopartition.conf", conf.Options)

	//mysql.MySQLQuery()

	//----------------------------mongo-------------------/
	//mongosample.MongoMain()

	//-------------------------net/http--------------------/
	//httpsample.HttpHello()
	//httpsample.HttpHello2()
	//httpsample.HttpHello3()
	//httpsample.HttpHello4()

	//-------------------------net/socket--------------------/
	//mysocket.SockServerMain()
	mysocket.SockClientMain()
}
