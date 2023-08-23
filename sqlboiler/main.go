package main

import (
	"database/sql"
	"fmt"
	"github.com/daobin/go-practice/sqlboiler/tables"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		log.Println("打开数据库失败：", err.Error())
		return
	}

	boil.SetDB(db)

	userList, err := tables.Users(
		//tables.UserWhere.Name.EQ(null.StringFrom("张三")),
		Offset(0),
		Limit(1),
		OrderBy("id desc"),
	).AllG()

	if err != nil {
		log.Println("获取用户列表失败：", err.Error())
		return
	}

	if len(userList) == 0 {
		log.Println("获取用户列表为空")
		return
	}

	for _, user := range userList {
		fmt.Printf("User Item[%v]; Age[%v]\n", user.Name, user.Age)
	}

	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		log.Println("开启事务失败：", err.Error())
		return
	}

	// 更新所有用户年龄
	_, err = tables.Users().UpdateAll(tx, tables.M{"age": 28})
	if err != nil {
		log.Println("更新用户AGE失败：", err.Error())
		return
	}

	userInfo, err := tables.FindUserG(1)
	if err != nil {
		log.Println("获取用户信息失败：", err.Error())
		return
	}
	fmt.Printf("User Info[%v]; Age[%v]\n", userInfo.Name, userInfo.Age)

	// 永远18
	userInfo.Age = 18
	_, err = userInfo.Update(tx, boil.Infer())

	// 事务的提交&回滚
	_ = tx.Commit()
	//_ = tx.Rollback()
}
