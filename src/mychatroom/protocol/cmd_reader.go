// @author     : wyettLei
// @date       : Created in 2021/4/20 12:08
// @description: TODO

package protocol

import (
	"bufio"
	"io"
	"log"
)

type CommandReader struct {
	reader *bufio.Reader
}

func NewCommandReader(reader io.Reader) *CommandReader {
	return &CommandReader{
		reader: bufio.NewReader(reader),
	}
}

func (r *CommandReader) Read() (interface{}, error) {

	command, err := r.reader.ReadString(' ')
	if err != nil {
		return nil, err
	}

	switch command {
	case "Message ":
		user, err := r.reader.ReadString(' ')
		if err != nil {
			return nil, err
		}
		message, err := r.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		return MessageCommand{
			user[:len(user)-1],
			message[:len(message)-1],
		}, nil
		// similar implementation for other commands
	default:
		log.Printf("Unknown command: %v", command)
	}

	return nil, UnKnownCommand
}
