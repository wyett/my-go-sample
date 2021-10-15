// @author     : wyettLei
// @date       : Created in 2021/4/19 10:16
// @description: TODO

package mysocket

import (
	"fmt"
	"net"
)

func SockServerMain() {

	// 1.listen
	fmt.Printf("start listen 8888...")
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Printf("listen err %s\n", err)
		return
	}
	defer listener.Close()

	// 2.accept
	//conn, err = listener.Accept()
	for {
		fmt.Println("waiting for connecting...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept err %s\n", err)
		} else {
			fmt.Printf("accept connect %v, from client %s\n", conn, conn.RemoteAddr().String())
		}

		go process(conn)
	}
}

func process(conn net.Conn) {

	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("read info error %v\n", err)
			return
		}

		fmt.Print(string(buf[:n]))
	}

}
