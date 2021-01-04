package common

/**
 * @author     : wyettLei
 * @date       : Created in 2020/12/18 13:42
 * @description: TODO
 */

const (
	_logFileName     = "auto_partition.log"
	_logLevelInfo    = "info"
	_logLevelWarn    = "warn"
	_logLevelError   = "error"
	_logMaxSize      = 500
	_logMaxAge       = 30
	_logEncodingJson = "json"
)

// logger 配置
type LoggerConfig struct {
	Level    string `config:"level"`     //日志级别 debug|info|warn|error
	Store    string `config:"store"`     //日志目录
	FileName string `config:"file_name"` //日志文件名称
	MaxSize  int    `config:"max_size"`  //日志文件最大M字节
	MaxAge   int    `config:"max_age"`   //日志文件最大存活的天数
	Compress bool   `config:"compress"`  //是否启用压缩
	Encoding string `config:"encoding"`  //日志编码 console|json
}
