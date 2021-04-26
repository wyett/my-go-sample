// @author     : wyettLei
// @date       : Created in 2021/4/20 12:11
// @description: TODO

package protocol

import (
	"bufio"
	"fmt"
	"io"
)

type CommandWriter struct {
	//writer io.Writer
	writer *bufio.Writer
}

func NewCommandWriter(writer io.Writer) *CommandWriter {
	return &CommandWriter{
		writer: bufio.NewWriter(writer),
	}
}

func (w *CommandWriter) writeString(msg string) error {
	_, err := w.writer.Write([]byte(msg))
	return err
}

func (w *CommandWriter) Write(command interface{}) error {
	var err error
	switch v := command.(type) {
	case SendCommand:
		err = w.writeString(fmt.Sprintf("SEND %v\n", v.Message))
	case NameCommand:
		err = w.writeString(fmt.Sprintf("Name %v\n", v.Name))
	case MessageCommand:
		err = w.writeString(fmt.Sprintf("Message %v %v\n", v.Name, v.Message))
	default:
		err = UnKnownCommand
	}
	return err
}
