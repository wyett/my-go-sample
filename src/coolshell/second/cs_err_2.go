package second

import (
	"encoding/binary"
	"io"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/3 18:06
 * @description: TODO
 */

//别名
//type Point types.Pointer
//type Point unsafe.Pointer
//type Point syscall.Pointer
type Point struct {
	Longitude     float32
	Latitude      float32
	Distance      int32
	ElevationGain float32
	ElevationLoss float32
}

func parser(f io.Reader) (*Point, error) {
	var p Point
	var err error
	read := func(data interface{}) {
		if err != nil {
			return
		}
		err = binary.Read(f, binary.BigEndian, data)
	}

	read(&p.Longitude)
	read(&p.Latitude)
	read(&p.Distance)
	read(&p.ElevationGain)
	read(&p.ElevationLoss)

	if err != nil {
		return &p, err
	}
	return &p, nil

}
