// @author     : wyettLei
// @date       : Created in 2021/2/25 15:32
// @description: TODO

package mylog

import (
	"log"
	"os"
)

var (
	f *os.File
)

func LogSample1Main() {

	// normal print
	log.Print("this is normal print")
	log.Printf("%s", "log printf")
	log.Println("println")

	// panic
	log.Panic("panic\n")
	log.Panicf("%s\n", "panicf")
	log.Panicln("panicln")

	// fatal
	log.Fatal("log fatal")
	log.Fatalf("%v", "fatal")
	log.Fatalln("fatalln")

	log.SetPrefix("print:")
	// normal print
	log.Print("this is normal print")
	log.Printf("%v", "log printf")
	log.Panicln("princln")
}

//------------------------------log to file---------------------------------//
var logDir = ""
var logFile = "mylog.log"

func LogSample2Main() {

	//os.OpenFile(logFile, )

}
