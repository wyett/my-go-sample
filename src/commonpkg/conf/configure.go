package conf

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
	MySQLProto    string   `config:"mysql.proto"`
	MySQLHost     []string `config:"mysql.host"`
	MySQLUser     string   `config:"mysql.user"`
	MySQLPassword string   `config:"mysql.password"`
	MySQLCharset  string   `config:"mysql.charset"`

	// 3. partition
	PartitionPreserveRule []string `config:"partition.preserve.rule"`

	// 4. log
	LogLevel    string `config:"log.level"`
	LogStore    string `config:"log.store"`
	LogName     string `config:"log.file_name"`
	LogMaxSize  uint32 `config:"log.max_size"`
	LogMaxAge   uint16 `config:"log.max_age"`
	LogCompress bool   `config:"log.compress"`
	LogEncoding string `config:"log.encoding"`
}

var Options PartitionConfig
