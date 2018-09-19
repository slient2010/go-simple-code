package main

import (
	"check-services-health/common"
	"check-services-health/res"
	"encoding/json"
	// "fmt"
	"log"
	"time"
)

type MonitorSrv struct {
	ChkSrvName string `json:"chksrv"`    // 服务健康检查名称
	ChkSrvUrl  string `json:"chksrvurl"` // 健康检查地址
	ReqstTimes int    `json:"int"`       // 健康检查次数
	ErrorTimes int    `json:"int"`       // 检查报错次数
}

func main() {
	// fmt.Println("vim-go")
	// get service health check url

	logurl := "https://dev-log.teyixing.com/health"

	for {
		// do http request
		res.ChkSrv(logurl)

		var m res.LogSrv
		a := common.HttpClientChkSrv(logurl)

		if err := json.Unmarshal([]byte(a), &m); err != nil {
			log.Fatal(err)
		}
		// fmt.Println(m.Data.Status)
		// fmt.Println(string(a))

		// save to influxdb

		common.SaveToInfluxDb(&m)

		// check this service status
		time.Sleep(30 * time.Second)
	}
}
