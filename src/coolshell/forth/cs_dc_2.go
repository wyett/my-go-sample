package forth

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/7 15:36
 * @description: TODO
 */

type IntSet struct {
	data map[int]bool
}

func NewIntSet() IntSet {
	return IntSet{make(map[int]bool)}
}

// add
func (set *IntSet) add(x int) {
	set.data[x] = true
}

// delete
func (set *IntSet) delete(x int) {
	delete(set.data, x)
}

// contain
func (set *IntSet) contain(x int) bool {
	return set.data[x]
}
