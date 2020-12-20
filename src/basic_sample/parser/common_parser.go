package parser

import (
	"flag"
	"fmt"
	"os"
)

type Exit struct {
	Code int
}

func BasicParserOps() {
	var err error
	defer handleExit()
	configuration := flag.String("config", "", "config file")

	if *configuration == "" {
		panic(Exit{0})
	}

	var file *os.File
	if file, err = os.Open(*configuration); err == nil {
		crash(fmt.Sprintf("read config file failed with %v", err), -1)
	}
	defer file.Close()
}

func crash(msg string, errCode int) {
	fmt.Println(msg)
	panic(Exit{errCode})
}

func handleExit() {
	if e := recover(); e != nil {
		if exit, ok := e.(Exit); ok == true {
			os.Exit(exit.Code)
		}
		panic(e)
	}
}
