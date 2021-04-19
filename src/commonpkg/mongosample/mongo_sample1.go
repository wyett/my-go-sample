// @author     : wyettLei
// @date       : Created in 2021/4/14 17:51
// @description: TODO

package mongosample

import (
	"fmt"
	"github.com/wyett/mgo"
)

type Person struct {
	NAME  string
	PHONE string
}
type Men struct {
	Persons []Person
}

const (
	//URL = "mongodb://wyett_rw:jumpjump@10.16.17.103:10020,10.16.17.107:10021"
	URL = "wyett_rw:jumpjump@10.16.17.103:10020"
)

func MongoMain() {

	session, err := mgo.Dial(URL) //连接数据库
	if err != nil {
		panic(err)
	}
	defer session.Close()
	//session.SetMode(mgo.Monotonic, true)

	db := session.DB("wyett")
	collection := db.C("person")

	//*****集合中元素数目********
	countNum, err := collection.Count()
	if err != nil {
		panic(err)
	}
	fmt.Println("Things objects count: ", countNum)

	////*******插入元素*******
	//temp := &Person{
	//    PHONE: "18811577546",
	//    NAME:  "zhangzheHero"}
	////一次可以插入多个对象 插入两个Person对象
	//err = collection.Insert(&Person{"Ale", "+55 53 8116 9639"}, temp)
	//if err != nil {
	//    panic(err)
	//}

	////*****查询单条数据*******
	//result := Person{}
	//err = collection.Find(bson.M{"phone": "456"}).One(&result)
	//fmt.Println("Phone:", result.NAME, result.PHONE)

	////*****查询多条数据*******
	//var personAll Men  //存放结果
	//iter := collection.Find(nil).Iter()
	//for iter.Next(&result) {
	//    fmt.Printf("Result: %v\n", result.NAME)
	//    personAll.Persons = append(personAll.Persons, result)
	//}

	////*******更新数据**********
	//err = collection.Update(bson.M{"name": "ccc"}, bson.M{"$set": bson.M{"name": "ddd"}})
	//err = collection.Update(bson.M{"name": "ddd"}, bson.M{"$set": bson.M{"phone": "12345678"}})
	//err = collection.Update(bson.M{"name": "aaa"}, bson.M{"phone": "1245", "name": "bbb"})

	////******删除数据************
	//_, err = collection.RemoveAll(bson.M{"name": "Ale”})

}
