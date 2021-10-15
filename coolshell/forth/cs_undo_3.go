package forth

import "errors"

/**
 * @author     : wyettLei
 * @date       : Created in 2021/2/7 15:36
 * @description: code from https://coolshell.cn/articles/21214.html
 */

type undoableIntSet struct {
	IntSet
	functions []func()
}

func NewUndoableIntSet() undoableIntSet {
	return undoableIntSet{NewIntSet(), nil}
}

// add中记录del操作
func (uis *undoableIntSet) add(x int) {
	if !uis.data[x] {
		uis.data[x] = true
		uis.functions = append(uis.functions, func() {
			uis.delete(x)
		})
	} else {
		uis.functions = append(uis.functions, nil)
	}
}

// del中记录add操作
func (uis *undoableIntSet) del(x int) {
	if uis.data[x] {
		uis.delete(x)
		uis.functions = append(uis.functions, func() {
			uis.add(x)
		})
	} else {
		uis.functions = append(uis.functions, nil)
	}
}

//Undo操作其实是一种控制逻辑，并不是业务逻辑，所以，在复用 Undo这个功能上是有问题。因为其中加入了大量跟 IntSet 相关的业务逻辑。
// undo操作中记录undo
// 获取最后一个functions并更新functions数组
func (uis *undoableIntSet) undo(x int) error {
	// 判断functions是否为空
	if len(uis.functions) == 0 {
		return errors.New("No function to undo")
	}
	//获取最后一个function的index
	index := len(uis.functions) - 1
	if f := uis.functions[index]; f != nil {
		// 执行获取到的f
		f()
		// 更新index位置为nil
		uis.functions[index] = nil
	}
	//去掉nil元素
	uis.functions = uis.functions[:index]
	return nil
}
