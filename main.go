package main

import (
	"fmt"
	"redis_sort/service"
)

func main() {
	// a := service.DataField{
	// 	VID:     400,
	// 	UpNum:   2,
	// 	CmtNum:  0,
	// 	AddTime: 1585456583,
	// }
	// _, err := service.AddData(&a)
	// fmt.Println(err)
	// _, err := service.DelData(&a)
	req := &service.ListReq{
		SortAct: "timeAsc",
		Page:    1,
		Size:    2,
	}
	li, err := service.ListData(req)
	fmt.Println(li, err)
}
