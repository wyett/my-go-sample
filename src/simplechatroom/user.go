// @author     : wyettLei
// @date       : Created in 2021/4/19 16:13
// @description: TODO

package simplechatroom

type UserInfo struct {
	userId   int32
	ipAddr   string
	username string
}

type UserLoginInfo struct {
	username string
	password string
}

type UserRegistryInfo struct {
	userId int32
	//userLoginInfo UserLoginInfo
	UserLoginInfo
}

type UserOperation interface {
	login()
	register()
}

//------------------------UserLoginInfo---------------------/

func NewUserLoginInfo(username string, password string) *UserLoginInfo {
	return &UserLoginInfo{
		username: username,
		password: password,
	}
}

func (uli *UserLoginInfo) SetUsername(username string) {
	uli.username = username
}

func (uli *UserLoginInfo) SetPassword(password string) {
	uli.password = password
}

func (uli *UserLoginInfo) Login(username string, password string) (bool, CommonResult) {
	cr := CommonResult{}
	// 1.
	return false, cr
}

//------------------------UserRegistryInfo---------------------/
func NewUserRegistryInfo(username string, password string) *UserRegistryInfo {
	return &UserRegistryInfo{
		userId: getUserId(),
		//{username:username, password:password},
		UserLoginInfo: {username, password},
	}
}
