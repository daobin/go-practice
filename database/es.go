package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type product struct {
	Name      string `json:"name"`
	CreatedAt int    `json:"created_at"`
}

func main() {
	esClient, err := elastic.NewClient(
		elastic.SetURL("http://192.168.190.134:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		fmt.Println("elastic.NewClient Err: ", err.Error())
		return
	}

	idxName := "product.sku"
	ctx := context.Background()

	// 判断索引是否存在
	idxExist, err := esClient.IndexExists(idxName).Do(ctx)
	if err != nil {
		fmt.Println("elastic.IndexExists Err: ", err.Error())
		return
	}

	fmt.Println("索引判断结果：", idxExist)
	//if idxExist {
	//	// 删除索引
	//	_, err = esClient.DeleteIndex(idxName).Do(ctx)
	//	if err != nil {
	//		fmt.Println("elastic.DeleteIndex Err: ", err.Error())
	//		return
	//	}
	//}
	//
	//// 索引映射
	//setting := map[string]any{
	//	"mappings": map[string]any{
	//		"properties": map[string]any{
	//			"name": map[string]any{
	//				"type": "text",
	//			},
	//			"created_at": map[string]any{
	//				"type": "integer",
	//			},
	//		},
	//	},
	//}
	//
	//// 创建索引
	//idxRes, err := esClient.CreateIndex(idxName).BodyJson(setting).Do(ctx)
	//if err != nil {
	//	fmt.Println("esClient.CreateIndex Err: ", err.Error())
	//	return
	//}
	//
	//fmt.Println("创建索引结果：", idxRes)

	//sku := "SKU10002"
	//
	//// 添加、更新记录
	//prod := product{
	//	Name:      "儿童玩具2",
	//	CreatedAt: 1695097004,
	//}
	//
	//addRes, err := esClient.Index().Index(idxName).Id(sku).BodyJson(prod).Do(ctx)
	//if err != nil {
	//	fmt.Println("esClient.Insert Err: ", err.Error())
	//	return
	//}
	//fmt.Println("添加文档结果：", addRes)

	//// 更新，允许添加新字段
	//updateRes, err := esClient.Update().Index(idxName).Id(sku).Doc(map[string]any{"name": "儿童玩具", "price": 9.9}).Do(ctx)
	//if err != nil {
	//	if elastic.IsNotFound(err) {
	//		fmt.Println("esClient.Update Not Found")
	//	} else {
	//		fmt.Println("esClient.Update Err: ", err.Error())
	//	}
	//	return
	//}
	//fmt.Println("更新文档结果：", updateRes)
	//
	//// 删除
	//delRes, err := esClient.Delete().Index(idxName).Id(sku).Do(ctx)
	//if err != nil {
	//	fmt.Println("esClient.Delete Err: ", err.Error())
	//	return
	//}
	//fmt.Println("删除文档结果：", delRes)

	// 查询
	//infoRes, err := esClient.Get().Index(idxName).Id(sku).Do(ctx)
	//if err != nil {
	//	if elastic.IsNotFound(err) {
	//		fmt.Println("esClient.Get Not Found")
	//	} else {
	//		fmt.Println("esClient.Get Err: ", err.Error())
	//	}
	//	return
	//}
	//fmt.Println("获取文档结果：", infoRes)
	//
	//prodInfo := product{}
	//data, _ := infoRes.Source.MarshalJSON()
	//_ = json.Unmarshal(data, &prodInfo)
	//fmt.Println("获取文档结果2：", prodInfo)

	//query := elastic.NewMatchQuery("_id", sku)
	//listRes, err := esClient.Search(idxName).Query(query).Do(ctx)

	page := 1
	limit := 10
	offset := (page - 1) * limit

	// 获取字段
	fsc := elastic.NewFetchSourceContext(true)
	fsc.Include("name", "created_at")

	// 查询条件
	query1 := elastic.NewMatchQuery("name", "玩具")
	searchReq1 := elastic.NewSearchRequest().
		Source(
			elastic.NewSearchSource().
				Query(query1).
				FetchSourceContext(fsc).
				Sort("_id", false).
				From(offset).
				Size(limit),
		)

	listRes, err := esClient.MultiSearch().
		Add(searchReq1).Do(ctx)
	if err != nil {
		fmt.Println("esClient.Search Err: ", err.Error())
		return
	}

	fmt.Println("查询文档结果：", len(listRes.Responses))
	fmt.Println("查询文档结果：", listRes.Responses[0])

	for _, infoRes := range listRes.Responses[0].Hits.Hits {
		data, _ := infoRes.Source.MarshalJSON()
		fmt.Println("查询文档结果：", string(data))
	}

}
