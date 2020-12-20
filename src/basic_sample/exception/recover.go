package exception

import "fmt"

// exit
type Exit struct {
	Code int
}

func BasicRecoverOps() {
	fmt.Println("begin BasicRecoverOps\n")
	defer handleExit()
	divNum()
	fmt.Printf("执行不到")

}

func handleExit() {
	if e := recover(); e != nil {
		fmt.Println("going wrong v%\n", e)
	}
}

func divNum() {
	a := 30
	b := 0
	var c = a / b
	fmt.Println("divNum() wrong %v", c)
	//panic(Exit{1})
}
