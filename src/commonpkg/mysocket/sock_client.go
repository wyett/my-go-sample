// @author     : wyettLei
// @date       : Created in 2021/4/19 10:17
// @description: TODO

package mysocket

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func SockClientMain() {

	// 1.与客户端建立链接
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("connect to server error ", err)
		return
	}
	defer conn.Close()

	// 2.从终端读入
	reader := bufio.NewReader(os.Stdin)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read info err %v", err)
		}

		// 3. check exit
		line := strings.Trim(s, "\t\n")
		if line == "exit" {
			fmt.Println("输入完成，退出")
			return
		}

		_, err = conn.Write([]byte(s))
		if err != nil {
			fmt.Println("write line failed ", err)
		}

	}
	//fmt.Printf("client send %d byte data and exit", n)

}
