package main

import (
	"fmt"
	"strings"
	"time"
)

type Reader interface {
	Read(rf chan string)
}

type ReadFile struct {
	FileRead string
}

func (r *ReadFile) Read(rf chan string) {
	msg := "safdsaf"
	rf <- msg
	// fmt.Println(r.FileRead)
}

type Writer interface {
	Write(wf chan string)
}

type WriteFile struct {
	FileWrite string
}

func (w *WriteFile) Write(wf chan string) {
	fmt.Println(<-wf)
}

type LogProcess struct {
	r     chan string
	w     chan string
	read  Reader
	write Writer
}

func (l *LogProcess) Process() {
	data := <-l.r
	l.w <- strings.ToUpper(data)
}

func main() {
	r := &ReadFile{
		FileRead: "sssss",
	}
	w := &WriteFile{
		FileWrite: "wwww",
	}
	p := &LogProcess{
		r:     make(chan string),
		w:     make(chan string),
		read:  r,
		write: w,
	}
	go p.read.Read(p.r)
	go p.Process()
	go p.write.Write(p.w)
	time.Sleep(3 * time.Second)
	fmt.Println("vim-go")
}
