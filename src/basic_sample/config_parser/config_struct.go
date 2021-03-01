package config_parser

/**
 * @author     : wyettLei
 * @date       : Created in 2021/1/15 14:27
 * @description: TODO
 */

type Mongo_Conn struct {
}

// log configuration
type LogConfig struct {
	level       string `config:log.level`
	store       string `config:log.store`
	logName     string `config:log.file_name`
	logMaxSize  uint32 `config:log.max_size`
	logMaxAge   uint16 `config:log.max_age`
	logCompress bool   `config:log.compress`
	logEncoding string `config:log.encoding`
}
