package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"xorm.io/xorm"
)

type User struct {
	Id      int64
	Name    string
	Salt    string    `xorm:"varchar(32) notnull"`
	Passwd  string    `xorm:"varchar(200) notnull"`
	Created time.Time `xorm:"notnull"`
	Updated time.Time `xorm:"notnull"`
}

var engine *xorm.Engine
var err error

func main() {
	url := "user:passwd@tcp(ip:port)/dbname?charset=utf8"
	engine, err = xorm.NewEngine("mysql", url)
	if err != nil {
		log.Println("NewEngine Error: ", err.Error())
		return
	}
	defer engine.Close()

	//engine.ShowSQL(true)

	err = engine.Sync(new(User))
	if err != nil {
		log.Println("Sync Error: ", err.Error())
		return
	}

	//user1 := User{
	//	Name:    "张三",
	//	Salt:    "dadfe",
	//	Passwd:  "seeeee",
	//	Created: time.Now(),
	//	Updated: time.Now(),
	//}
	//cnt, err := engine.Insert(user1, user1)
	//if err != nil {
	//	log.Println("Insert Error: ", err.Error())
	//	return
	//}
	//
	//log.Println("Affected Count: ", cnt)

	info := User{}
	_, _ = engine.Desc("id").Get(&info)
	fmt.Printf("%v", info)
}
