package main

import (
	"fmt"
	"strings"
	"time"
)

type LogProcess struct {
	rc          chan string //读取channel
	wc          chan string //写入channel
	path        string      // 读取日志文件路劲
	influxDBDsn string      // influx data source
}

func (l *LogProcess) ReadFromFile() {
	line := "message"
	l.rc <- line
	// 读取模块

}

func (l *LogProcess) Process() {
	data := <-l.rc

	l.wc <- strings.ToUpper(data)
	// 处理模块
}

// * 1.是引用，结构体足够大，减少了copy操作，提升性能
//  2. 用l直接修改自身定义的参数
func (l *LogProcess) WriteToInfluxDB() {
	// 写入模块
	fmt.Println(<-l.wc)
}

func main() {
	// & 性能上的考虑，lp是引用类型的
	lp := &LogProcess{
		rc:          make(chan string),
		wc:          make(chan string),
		path:        "/tmp/access.log",
		influxDBDsn: "username&password..",
	}
	//go (*lp).ReadFromFile() // 也可以，只是不宜读
	go lp.ReadFromFile()    //读文件
	go lp.Process()         // 处理文件
	go lp.WriteToInfluxDB() //写入数据库
	time.Sleep(1 * time.Second)
}
