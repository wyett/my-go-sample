// @author     : wyettLei
// @date       : Created in 2021/4/19 18:35
// @description: common struct

package simplechatroom

import (
	"errors"
	"fmt"
	"os"
)

type CommonResult struct {
	Code int
	Err  error
}

type CommonHandler interface {
	HandleExit()
	Crash()
}

func NewCommonResult(code int, err error) *CommonResult {
	return &CommonResult{Code: code, Err: err}
}

func (cr *CommonResult) SetResultCode(code int) {
	cr.Code = code
}

func (cr *CommonResult) SetResultErr(err error) {
	cr.Err = err
}

func (cr *CommonResult) HandleExit() error {
	if e := recover(); e != nil {
		if exit, ok := e.(CommonResult); ok == true {
			os.Exit(exit.Code)
		}
		panic(e)
	}
}

func (cr *CommonResult) Crash(msg string, errCode int) {
	fmt.Println(msg)
	panic(CommonResult{errCode, errors.New(msg)})
}
