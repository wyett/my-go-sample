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
/*mysql> select table_schema,table_name,partition_name,PARTITION_ORDINAL_POSITION,PARTITION_METHOD,PARTITION_EXPRESSION,PARTITION_DESCRIPTION,CREATE_TIME from partitions where TABLE_SCHEMA='wyett';
+--------------+-------------------+----------------+----------------------------+------------------+----------------------+-----------------------+---------------------+
| table_schema | table_name        | partition_name | PARTITION_ORDINAL_POSITION | PARTITION_METHOD | PARTITION_EXPRESSION | PARTITION_DESCRIPTION | CREATE_TIME         |
+--------------+-------------------+----------------+----------------------------+------------------+----------------------+-----------------------+---------------------+
| wyett        | w_partition_demo  | p202001        |                          1 | RANGE            | to_days(create_time) | 737821                | 2021-02-18 14:07:55 |
| wyett        | w_partition_demo  | p202002        |                          2 | RANGE            | to_days(create_time) | 737850                | 2021-02-18 14:07:55 |
| wyett        | w_partition_demo  | p202003        |                          3 | RANGE            | to_days(create_time) | 737881                | 2021-02-18 14:07:55 |
| wyett        | w_partition_demo  | p202004        |                          4 | RANGE            | to_days(create_time) | 737911                | 2021-02-18 14:07:55 |
| wyett        | w_partition_demo  | p202005        |                          5 | RANGE            | to_days(create_time) | 737942                | 2021-02-18 14:07:55 |
| wyett        | w_partition_demo  | p202006        |                          6 | RANGE            | to_days(create_time) | 737972                | 2021-02-18 14:07:55 |
| wyett        | w_partition_demo  | p202007        |                          7 | RANGE            | to_days(create_time) | 738003                | 2021-02-18 14:07:55 |
| wyett        | w_partition_demo  | p202008        |                          8 | RANGE            | to_days(create_time) | 738034                | 2021-02-18 14:07:55 |
| wyett        | w_partition_demo  | p202009        |                          9 | RANGE            | to_days(create_time) | 738064                | 2021-02-18 14:07:55 |
| wyett        | w_partition_demo  | p202010        |                         10 | RANGE            | to_days(create_time) | 738095                | 2021-02-18 14:07:55 |
| wyett        | w_partition_demo  | p202011        |                         11 | RANGE            | to_days(create_time) | 738125                | 2021-02-18 14:07:55 |
| wyett        | w_partition_demo  | p202012        |                         12 | RANGE            | to_days(create_time) | 738156                | 2021-02-18 14:07:55 |
| wyett        | w_partition_demo  | pother         |                         13 | RANGE            | to_days(create_time) | MAXVALUE              | 2021-02-18 14:07:55 |
| wyett        | w_partition_demo1 | p202001        |                          1 | RANGE            | to_days(create_time) | 737821                | 2021-02-18 14:12:19 |
| wyett        | w_partition_demo1 | p202002        |                          2 | RANGE            | to_days(create_time) | 737850                | 2021-02-18 14:12:19 |
| wyett        | w_partition_demo1 | p202003        |                          3 | RANGE            | to_days(create_time) | 737881                | 2021-02-18 14:12:19 |
| wyett        | w_partition_demo1 | p202004        |                          4 | RANGE            | to_days(create_time) | 737911                | 2021-02-18 14:12:19 |
| wyett        | w_partition_demo1 | p202005        |                          5 | RANGE            | to_days(create_time) | 737942                | 2021-02-18 14:12:19 |
| wyett        | w_partition_demo1 | p202006        |                          6 | RANGE            | to_days(create_time) | 737972                | 2021-02-18 14:12:19 |
| wyett        | w_partition_demo1 | p202007        |                          7 | RANGE            | to_days(create_time) | 738003                | 2021-02-18 14:12:19 |
| wyett        | w_partition_demo1 | p202008        |                          8 | RANGE            | to_days(create_time) | 738034                | 2021-02-18 14:12:19 |
| wyett        | w_partition_demo1 | p202009        |                          9 | RANGE            | to_days(create_time) | 738064                | 2021-02-18 14:12:19 |
| wyett        | w_partition_demo1 | p202010        |                         10 | RANGE            | to_days(create_time) | 738095                | 2021-02-18 14:12:19 |
| wyett        | w_partition_demo1 | p202011        |                         11 | RANGE            | to_days(create_time) | 738125                | 2021-02-18 14:12:19 |
| wyett        | w_partition_demo1 | p202012        |                         12 | RANGE            | to_days(create_time) | 738156                | 2021-02-18 14:12:19 |
| wyett        | w_partition_demo1 | pother         |                         13 | RANGE            | to_days(create_time) | MAXVALUE              | 2021-02-18 14:12:19 |
+--------------+-------------------+----------------+----------------------------+------------------+----------------------+-----------------------+---------------------+
*/

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
