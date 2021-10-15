package forth

import "fmt"

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/7 15:35
 * @description: TODO
 */

type weight struct {
	x, y int
}

type label struct {
	weight
	text string
}

type button struct {
	label
}

type listBox struct {
	weight
	texts []string
	index int
}

//-------------------------

type painter interface {
	paint()
}

type clicker interface {
	click()
}

//------------------------

func (label label) paint() {
	fmt.Printf("label==>%d * %d", label.weight.x, label.y)
}

func (button button) paint() {
	fmt.Printf("button==>%s", button.text)
}

func (button button) click() {
	fmt.Printf("button==>%s", button.text)
}

func (listBox listBox) paint() {
	fmt.Printf("listBox==>%s", listBox.texts)
}

func (listBox listBox) click() {
	fmt.Printf("listBox==>%s", listBox.texts)
}

func NewButton(x int, y int, text string) button {
	return button{label{weight{x, y}, text}}
}

//--------------------

func CsPloy1Main() {
	b1 := button{label{weight{10, 2}, "button1"}}
	b2 := NewButton(10, 2, "button2")
	lb := listBox{weight{10, 2}, []string{"aaa", "bb"}, 0}

	for _, painter := range []painter{b1, b2, lb} {
		painter.paint()
	}

	for _, weight := range []interface{}{b1, b2, lb} {
		weight.(painter).paint()
		if clicker, ok := weight.(clicker); ok {
			clicker.click()
		}
	}
}
