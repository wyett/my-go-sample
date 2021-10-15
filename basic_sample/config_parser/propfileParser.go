package config_parser

/**
 * @author     : wyettLei
 * @date       : Created in 2020/12/25 16:53
 * @description: TODO
 */

type PropertyConfig struct {

	// module1
	Aaaaa string `config:"module1.a"`
	Bbbbb string `config:"module1.b"`
	Ccccc string `config:"module1.c"`

	// module2
	Ddddd string `config:"module1.a"`
	Eeeee string `config:"module1.b"`

	// module3
	Fffff string `config:"module1.a"`
}

//func LoadProperties(f string) {
//	// open file
//	data, err := os.Open(f)
//	if err != nil {
//		fmt.Println("file %v not exists", f)
//	}
//	fmt.Println(data)
//
//}
