package forth

import "errors"

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/7 18:16
 * @description: 不再由 控制逻辑 Undo 来依赖业务逻辑 stringSet，而是由业务逻辑 stringSet 来依赖 Undo
 */

// undo
type undo []func()

func (ud *undo) undo() error {
	functions := *ud
	if len(functions) == 0 {
		return errors.New("function does not exists")
	}
	index := len(functions) - 1
	if function := functions[index]; function != nil {
		function()
		functions[index] = nil
	}
	functions = functions[:index]
	return nil
}

func (ud *undo) add(function func()) {
	*ud = append(*ud, function)
}

// set
type stringSet struct {
	data map[string]bool
	ud   undo
}

func NewStringSet() stringSet {
	return stringSet{data: make(map[string]bool)}
}

// undo
func (ss *stringSet) undo() error {
	return ss.ud.undo()
}

// contains
func (ss *stringSet) contains(s string) bool {
	return ss.data[s]
}

// add
func (ss *stringSet) add(s string) {
	if !ss.contains(s) {
		ss.data[s] = true
		ss.ud.add(func() { ss.delete(s) })
	} else {
		ss.ud.add(nil)
	}
}

// delete
func (ss *stringSet) delete(s string) {
	if ss.contains(s) {
		delete(ss.data, s)
		ss.ud.add(func() { ss.add(s) })
	} else {
		ss.ud.add(nil)
	}
}
