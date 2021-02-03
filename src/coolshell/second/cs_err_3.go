package second

import (
	"encoding/binary"
	"io"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/3 18:58
 * @description: TODO
 */

type reader struct {
	r   io.Reader
	err error
}

func (r *reader) read(data interface{}) {
	if r.err == nil {
		r.err = binary.Read(r.r, binary.BigEndian, data)
	}
}

func parser3(f io.Reader) (*Point, error) {
	var p Point
	r := reader{r: f}

	r.read(&p.Longitude)
	r.read(&p.Latitude)
	r.read(&p.Distance)
	r.read(&p.ElevationGain)
	r.read(&p.ElevationLoss)

	if r.err != nil {
		return &p, r.err
	}
	return &p, nil

}
