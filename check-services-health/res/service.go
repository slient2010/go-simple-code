// get srv urls
package res

import (
	// "encoding/json"
	"fmt"
)

type LogSrv struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    struct {
		Status string `json:"status"`
	} `json:"data"`
}

// 从数据库中获取数据
func GetSrvHlthUrl() (string, error) {
	data := GetData()

	for _, v := range data {
		// fmt.Println(v)
		fmt.Println(v.AppUrl)
	}

	return "", nil
}

// func ChkSrv(url string) (string, error) {
func ChkSrv(url string) {
	// 获取健康检查地址
	data, _ := GetSrvHlthUrl()
	fmt.Println(data)
	// return "", nil
}
