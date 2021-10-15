// @author     : wyettLei
// @date       : Created in 2021/4/19 16:13
// @description: TODO

package simplechatroom

const (
	LoginMsgType    = "LoginMsg"
	LoginResMsgType = "LoginMsgType"
)

type Message struct {
	Type string
	Data string
}

type LoginMsg struct {
	UserId   string
	Username string
	Userpwd  string
}

type LoginResMsg struct {
	ErrCode string
	Err     error
}
