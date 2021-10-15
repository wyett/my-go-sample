package myerror

/**
 * @author     : wyettLei
 * @date       : Created in 2020/12/24 17:29
 * @description: TODO
 */

type PError struct {
	msg string
}

func New(msg string) error {
	return &PError{msg}
}

func (pe *PError) Error() string {
	return pe.msg
}
