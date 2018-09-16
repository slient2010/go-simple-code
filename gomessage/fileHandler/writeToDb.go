package fileHandler

import "fmt"

type Writer interface{
	Write(wc chan string)
}


type WriteToInfluxDB struct {
	InfluxDBDsn string
}

func (w *WriteToInfluxDB) Write(wc chan string) {
	fmt.Println(<-wc)
}