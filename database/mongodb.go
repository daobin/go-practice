package main

import (
	"fmt"
	"github.com/daobin/gotools"
	"gopkg.in/mgo.v2/bson"
	"os"
	"path/filepath"
)

type testData struct {
	Name  string `bson:"name"`  // 名称
	Age   int    `bson:"age"`   // 年龄
	Class string `bson:"class"` // 班级
}

// getMgoYmlFile 获取配置文件
func getMgoYmlFile() (string, error) {
	ymlFile, err := filepath.Abs("./database/mongo.yml")
	fmt.Println("YML File: ", ymlFile)

	if err != nil {
		fmt.Println("配置文件获取错误：", err.Error())
		return "", err
	}

	_, err = os.Stat(ymlFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("配置文件不存在")
			return "", err
		}

		fmt.Println("配置文件无效：", err.Error())
		return "", err
	}

	return ymlFile, nil
}

// initMgoYmlFile 初始化配置
func initMgoYmlFile(ymlFile string) error {
	err := gotools.DB.Mongo.Init(ymlFile)
	if err != nil {
		fmt.Println("初始化错误：", err.Error())
		return err
	}

	return nil
}

func getListMgoData(find bson.M, page, limit int) {
	fmt.Println("开始查询**************************")
	conn, err := gotools.DB.Mongo.Get("dev")
	if err != nil {
		fmt.Println("获取连接失败：", err.Error())
		return
	}
	defer gotools.DB.Mongo.CloseCurrent(conn)

	fmt.Println("查询条件：", find)
	fmt.Println("页码：", page)
	fmt.Println("每页数量：", limit)

	count, err := conn.C("test_data").Find(find).Count()
	if err != nil {
		fmt.Println("获取数据量失败：", err.Error())
		return
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	list := make([]testData, 0)
	err = conn.C("test_data").Find(find).Skip(offset).Limit(limit).All(&list)
	if err != nil {
		fmt.Println("获取数据失败：", err.Error())
		return
	}

	fmt.Println("数据量：", count)
	fmt.Println("数据列表：", list)
}

func createMgoData(tdList []interface{}) {
	fmt.Println("开始新增**************************")
	conn, err := gotools.DB.Mongo.Get("dev")
	if err != nil {
		fmt.Println("获取连接失败：", err.Error())
		return
	}
	defer gotools.DB.Mongo.CloseCurrent(conn)

	fmt.Println("新增数据：", tdList)
	if len(tdList) == 0 {
		fmt.Println("新增数据为空")
		return
	}

	err = conn.C("test_data").Insert(tdList...)
	if err != nil {
		fmt.Println("新增失败：", err.Error())
	}

	fmt.Println("新增成功")
}

func updateMgoData(find, update bson.M) {
	fmt.Println("开始修改**************************")
	conn, err := gotools.DB.Mongo.Get("dev")
	if err != nil {
		fmt.Println("获取连接失败：", err.Error())
		return
	}
	defer gotools.DB.Mongo.CloseCurrent(conn)

	fmt.Println("修改条件：", find)
	fmt.Println("修改数据：", update)

	_, err = conn.C("test_data").UpdateAll(find, bson.M{"$set": update})
	if err != nil {
		fmt.Println("修改失败：", err.Error())
	}

	fmt.Println("修改成功")
}

func deleteMgoData(find bson.M) {
	fmt.Println("开始删除**************************")
	conn, err := gotools.DB.Mongo.Get("dev")
	if err != nil {
		fmt.Println("获取连接失败：", err.Error())
		return
	}
	defer gotools.DB.Mongo.CloseCurrent(conn)

	fmt.Println("删除条件：", find)

	_, err = conn.C("test_data").RemoveAll(find)
	if err != nil {
		fmt.Println("删除数据失败：", err.Error())
		return
	}

	fmt.Println("删除数据成功")
}

func main() {
	// 获取配置文件
	ymlFile, err := getMgoYmlFile()
	if err != nil {
		return
	}

	// 初始化配置
	err = initMgoYmlFile(ymlFile)
	if err != nil {
		return
	}

	// 查
	getListMgoData(bson.M{}, 1, 10)

	// 增
	tdList := make([]interface{}, 2)
	tdList[0] = testData{
		Name:  "张三",
		Age:   10,
		Class: "3年级",
	}
	tdList[1] = testData{
		Name:  "张三2",
		Age:   10,
		Class: "3年级",
	}
	createMgoData(tdList)
	getListMgoData(bson.M{}, 1, 10)

	// 改
	updateMgoData(bson.M{"name": "张三2"}, bson.M{"name": "李四"})
	getListMgoData(bson.M{}, 1, 10)

	// 删
	deleteMgoData(bson.M{"name": "李四"})
	getListMgoData(bson.M{}, 1, 10)
	getListMgoData(bson.M{"name": "李四"}, 1, 10)

	fmt.Println("Success")
}
