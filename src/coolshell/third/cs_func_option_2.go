package third

import (
	"fmt"
	"time"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/7 12:07
 * @description: TODO
 */

type mongoConfig struct {
	//必须
	host   []hpPair
	user   string
	passwd string
	authdb string
	// 非必须
	preference string
	timeout    time.Duration
}

// 定义函数类型
type option func(*mongoConfig)

// 定义函数
func host(hp []hpPair) option {
	return func(mc *mongoConfig) {
		mc.host = hp
	}
}

func user(user string) option {
	return func(mc *mongoConfig) {
		mc.user = user
	}
}

func passwd(passwd string) option {
	return func(mc *mongoConfig) {
		mc.passwd = passwd
	}
}

func authdb(authdb string) option {
	return func(mc *mongoConfig) {
		mc.authdb = authdb
	}
}

func preference(preference string) option {
	return func(mc *mongoConfig) {
		mc.preference = preference
	}
}

func timeout(timeout time.Duration) option {
	return func(mc *mongoConfig) {
		mc.timeout = timeout
	}
}

func NewMongoConfig(hp []hpPair, user string, passwd string, authdb string,
	options ...func(*mongoConfig)) (*mongoConfig, error) {

	// mongoConfig
	mc := mongoConfig{
		host:       hp,
		user:       user,
		passwd:     passwd,
		authdb:     authdb,
		preference: "secondary",
		timeout:    30 * time.Second,
	}

	for _, option := range options {
		option(&mc)
	}

	return &mc, nil
}

// tostring
func (mc *mongoConfig) print() {
	/*
	   fmt.Printf("%v:host=%s, user=%s, passwd=%s, authdb=%s, preference=%s, timeout=%d",
	       mc, mc.host, mc.user, mc.passwd, mc.authdb, mc.preference, mc.timeout)
	*/
	fmt.Printf("%T==>%+v\n", mc, mc)
}

func CsFuncOption2Main() {
	mc1, _ := NewMongoConfig([]hpPair{{"127.0.0.1", 27017}, {"127.0.0.1", 27018}},
		"wyett_rw", "123456", "wyett")
	print(mc1)

	mc2, _ := NewMongoConfig([]hpPair{{"127.0.0.1", 27017}, {"127.0.0.1", 27018}},
		"wyett_rw", "123456", "wyett", preference("secondary"))
	print(mc2)

	mc3, _ := NewMongoConfig([]hpPair{{"127.0.0.1", 27017}, {"127.0.0.1", 27018}},
		"wyett_rw", "123456", "wyett", preference("secondary"), timeout(3000))
	print(mc3)

}
