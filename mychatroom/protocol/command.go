// @author     : wyettLei
// @date       : Created in 2021/4/20 12:11
// @description: TODO

package protocol

// sending message from client
type SendCommand struct {
	Message string
}

// setting client display name
type NameCommand struct {
	Name string
}

// broadcast message
type MessageCommand struct {
	Name    string
	Message string
}
