package config

import "time"

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/13 21:01
 * @description: TODO
 */

// hostPort
type HostPortPair struct {
	host string
	port uint32
}

// mysql connection
type MySQLConnection struct {
	HpPairs  []HostPortPair
	username string
	passwd   string
	charset  string
}

// information partition table
type MySQLPartition struct {
	TableSchema              string
	TableName                string
	PartitionName            string
	PartitionOrdinalPosition uint16
	PartitionMethod          string
	PartitionExpression      string
	PartitionDesc            string
	Ctime                    time.Duration
}

type PartitionStrategy struct {
	TableSchema string
	TableName   string
	Retation    uint32
	Unit        time.Duration
	//ticker time.Ticker
}

// mongo connection
type MongoConnection struct {
	HpPairs             []HostPortPair
	Replset             string
	Username            string
	Passwd              string
	SecondaryPreference string
}

type LogConfig struct {
	level       string `log.level`
	store       string `log.store`
	logName     string `log.file_name`
	logMaxSize  uint32 `log.max_size`
	logMaxAge   uint16 `log.max_age`
	logCompress bool   `log.compress`
	logEncoding string `log.encoding`
}
