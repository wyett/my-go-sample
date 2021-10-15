package first

import "fmt"

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/2 14:26
 * @description: TODO
 */

type Person struct {
	name   string
	sexual string
	age    int
}

func printPerson(p *Person) {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n", p.name, p.sexual, p.age)
}

func (p *Person) printP() {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n", p.name, p.sexual, p.age)
}

func PersonPrint() {
	var p = Person{
		name:   "wyett",
		sexual: "male",
		age:    33,
	}

	printPerson(&p)
	p.printP()
}
