// @author     : wyettLei
// @date       : Created in 2021/3/4 16:53
// @description: TODO

package mysync

import "sync"

// 通过sync.once保证channel仅关闭一次
type wChannel struct {
	c    chan interface{}
	once sync.Once
}

func NewWChannel() *wChannel {
	return &wChannel{c: make(chan interface{})}
}

func (wc *wChannel) safeClose() {
	wc.once.Do(func() {
		close(wc.c)
	})
}
