// @author     : wyettLei
// @date       : Created in 2021/2/26 12:09
// @description: TODO

package mycontext

import (
	"context"
	log2 "log"
	"os"
	"time"
)

var contextLog *log2.Logger

func MyContextSample2Main() {
	//log2.SetPrefix("MyContext Sample2")
	contextLog = log2.New(os.Stdout, "MyContext Sample2:", log2.Ltime)
	myContextControl()

}

func myContextControl() {

	contextLog.Println("main begin")
	// context
	context, canalFunc := context.WithCancel(context.Background())
	go doQuery2(context)
	doQuery1()
	canalFunc()
}

func doQuery1() {
	contextLog.Println("doQuery 1 begin")
	time.Sleep(time.Second * 3)
}

func doQuery2(ctx context.Context) {
	contextLog.Println("doQuery 2 begin")
	for {
		select {
		case <-ctx.Done():
			contextLog.Println("ctx.done called")
			return
		default:
			time.Sleep(time.Second)
			contextLog.Println("query2 wait for query1")
		}
	}
}
