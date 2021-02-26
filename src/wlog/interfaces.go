// @author     : wyettLei
// @date       : Created in 2021/2/26 17:02
// @description: TODO

package wlog

import "os"

// write log file
type WFile interface {
	// create log file
	CreateLogFile() *os.File

	// create Channel
	CreateLogChannel() chan string

	// log rotate
	Rotate(*os.File) bool
}
