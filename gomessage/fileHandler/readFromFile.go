package fileHandler


// 读取模块抽象出来
type Reader interface{
	Read(rc chan string) //Read方法

}

// 定义
type ReadFromFile struct {
	Path string // 读取日志文件路劲
}


// 实现接口
func (r *ReadFromFile) Read(rc chan string){
	// 读取模块
	line := "message"
	rc <- line

}

