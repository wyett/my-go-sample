package myio

import (
	"fmt"
	"os"
	"path/filepath"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2020/12/18 14:21
 * @description: TODO
 */

func BasicStatOps() {
	fileInfo, err := os.Stat("E:\\mygit\\my-go-sample\\conf\\aa.txt")
	if err != nil {
		fmt.Println("stat file failed", err)
		return
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.Sys())
}

func BasicDirOps() {
	aa := "E:\\mygit\\my-go-sample\\conf\\aa.txt"
	bb := "bb.txt"

	fmt.Println(filepath.IsAbs(aa))
	fmt.Println(filepath.IsAbs(bb))
	fmt.Println(filepath.Abs(aa))
	fmt.Println(filepath.Abs(bb))

	fmt.Println(filepath.Join(aa, ".."))
}
