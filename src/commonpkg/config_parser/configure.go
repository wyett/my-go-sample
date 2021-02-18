package config

/**
 * @author     : wyettLei
 * @date       : Created in 2020/12/25 14:35
 * @description: config struct
 */

type PartitionConfig struct {
	// 0. version
	ConfVersion uint `config:"conf.version"` // do not modify the tag name

	// 1. common
	Id string `config:"id"`

	// 2. mysql
	MySQLHost     []string `config:"mysql.host"`
	MySQLUser     string   `config:"mysql.user"`
	MySQLPassword string   `config:"mysql.password"`
	MySQLCharset  string   `config:"mysql.charset"`

	// 3. partition
	PartitionPreserveRule []string `config:"partition.preserve.rule"`

	// 4. log
	level       string `config:log.level`
	store       string `config:log.store`
	logName     string `config:log.file_name`
	logMaxSize  uint32 `config:log.max_size`
	logMaxAge   uint16 `config:log.max_age`
	logCompress bool   `config:log.compress`
	logEncoding string `config:log.encoding`
}

var Option PartitionConfig
