package main

import (
	"fmt"
	"miao/backend/common"
	"miao/datamodels"
)

func  main()  {
	data := map[string]string{
		"name" : "jack",
		//"id" : "2",
		"num" : "0",
		"image" : "http://",
		"url" : "了我确认",
		"bo" : "true",
	}

	result := &datamodels.Product{}
	common.DataToStructByTagSql(data,result)
	fmt.Printf("+%v",*result)
}
