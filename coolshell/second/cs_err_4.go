package second

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/4 10:33
 * @description: interface fluence program/Fluent Interface
 */

var b = []byte{0x48, 0x61, 0x6f, 0x20, 0x43, 0x68, 0x65, 0x6e, 0x00, 0x00, 0x2c}
var r = bytes.NewReader(b)

type person struct {
	name   []byte
	age    uint8
	weight uint8
	err    error
}

func (p *person) read(data interface{}) {
	if p.err == nil {
		p.err = binary.Read(r, binary.BigEndian, data)
	}
}

func (p *person) readName() *person {
	p.read(&p.name)
	return p
}

func (p *person) readAge() *person {
	p.read(&p.age)
	return p
}

func (p *person) readWeight() *person {
	p.read(&p.weight)
	return p
}

func (p *person) print() *person {
	if p.err == nil {
		fmt.Printf("name=%s, age=%d, weight=%d", p.name, p.age, p.weight)
	}
	return p
}

func CsErr4Main() {
	p := person{}
	p.readName().readAge().readWeight().print()
	fmt.Println(p.err)
}
