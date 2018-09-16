package main

import (
	"strings"
	"time"
	"logservices/fileHandler"
)





type LogProcess struct {
	rc chan string //读取channel
	wc chan string //写入channel
	read fileHandler.Reader
	write fileHandler.Writer
}


func (l *LogProcess) Process(){
	data := <- l.rc

	l.wc <-strings.ToUpper(data)
	// 处理模块
}

func main() {
	// 实例化
	r := &fileHandler.ReadFromFile{
		Path: "/tmp/access.log",
	}

	w := &fileHandler.WriteToInfluxDB{
		InfluxDBDsn:"username&password..",
	}

	// & 性能上的考虑，lp是引用类型的
	lp := &LogProcess{
		rc: make(chan string),
		wc: make(chan string),
		read: r,
		write: w,

	}
	//go (*lp).ReadFromFile() // 也可以，只是不宜读
	go lp.read.Read(lp.rc)
	go lp.Process()  // 处理文件
	go lp.write.Write(lp.wc)
	time.Sleep(1 * time.Second)
}
