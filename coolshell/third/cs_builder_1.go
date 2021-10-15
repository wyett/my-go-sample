package third

import "fmt"

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/7 11:36
 * @description: TODO
 */

type hpPair struct {
	host string
	port uint32
}

type mysqlConfig struct {
	server  []hpPair
	user    string
	passwd  string
	dbname  string
	charset string
}

type mysqlConfigBuilder struct {
	mysqlConfig
}

func (m *mysqlConfigBuilder) addServer(hp []hpPair) *mysqlConfigBuilder {
	m.mysqlConfig.server = hp
	return m
}

func (m *mysqlConfigBuilder) addUser(user string) *mysqlConfigBuilder {
	m.mysqlConfig.user = user
	return m
}

func (m *mysqlConfigBuilder) addPasswd(passwd string) *mysqlConfigBuilder {
	m.mysqlConfig.passwd = passwd
	return m
}

func (m *mysqlConfigBuilder) addDbname(dbname string) *mysqlConfigBuilder {
	m.mysqlConfig.dbname = dbname
	return m
}

func (m *mysqlConfigBuilder) addCharset(charset string) *mysqlConfigBuilder {
	m.mysqlConfig.charset = charset
	return m
}

func (m *mysqlConfigBuilder) build() mysqlConfig {
	return m.mysqlConfig
}

func CsBuilder1Main() {
	mcb := mysqlConfigBuilder{}
	fmt.Println(mcb.addServer([]hpPair{{"127.0.0.1", 3306}, {"127.0.0.1", 3307}}).
		addUser("root").
		addPasswd("123456").
		addDbname("wyett").
		addCharset("utf-8").
		build())
}
