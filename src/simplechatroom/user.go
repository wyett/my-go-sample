// @author     : wyettLei
// @date       : Created in 2021/4/19 16:13
// @description: TODO

package simplechatroom

type UserInfo struct {
	userId   UUID
	ipAddr   string
	username string
}

type UserLoginInfo struct {
	username string
	password string
}

type UserRegistryInfo struct {
	userId UUID
	UserLoginInfo
}

type UserOperation interface {
	Login()
}

type RegistryOperation interface {
	Register()
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
	// todo
	return false, cr
}

//------------------------UserRegistryInfo---------------------/
func NewUserRegistryInfo(username string, password string) *UserRegistryInfo {
	return &UserRegistryInfo{
		userId:        UUID{},
		UserLoginInfo: UserLoginInfo{username, password},
	}
}

func (uri *UserRegistryInfo) SetUserId(userId UUID) {
	uri.userId = userId
}

func (uri *UserRegistryInfo) SetUserLoginInfo(uli UserLoginInfo) {
	uri.UserLoginInfo = uli
}

func (uri *UserRegistryInfo) Register(info UserRegistryInfo) (bool, CommonResult) {
	cr := CommonResult{}
	// todo
	return true, cr
}
