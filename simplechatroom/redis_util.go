// @author     : wyettLei
// @date       : Created in 2021/4/20 9:16
// @description: redis pool

package simplechatroom

import (
	"errors"
	"github.com/go-redis/redis/v8"
)

// redis pool
type RedisPool struct {
	//conn *redis.Conn
	connPool    []redis.Conn
	poolSize    int32
	maxIdle     int32
	maxActive   int32
	idleTimeout int32
}

// interface for redis pool
type RedisPoolOps interface {
	addConn()
	getConn()
	releaseConn()
}

func NewRedisPool(conn []redis.Conn, poolSize int32, maxIdle int32, maxActive int32, idleTimeout int32) (*RedisPool,
	CommonResult) {
	cr := CommonResult{}
	return &RedisPool{
		connPool:    make([]redis.Conn, poolSize),
		maxIdle:     maxIdle,
		maxActive:   maxActive,
		idleTimeout: idleTimeout,
	}, cr
}

// setter
func (rp *RedisPool) SetPoolSize(poolSize int32) {
	rp.poolSize = poolSize
}

func (rp *RedisPool) SetMaxIdle(maxIdle int32) {
	rp.maxIdle = maxIdle
}

func (rp *RedisPool) SetMaxActive(maxActive int32) {
	rp.maxActive = maxActive
}

func (rp *RedisPool) SetIdleTimeout(idleTimeout int32) {
	rp.idleTimeout = idleTimeout
}

func (rp *RedisPool) addConn(conn redis.Conn) error {
	if cap(rp.connPool) > len(rp.connPool) {
		//append(rp.connPool, conn)
		return nil
	}

	return errors.New("连接池已满，请稍后创建新链接")
}
