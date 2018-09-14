// 多态
package main

import (
	"fmt"
)

// interface 一组方法的集合
type Foo interface {
	qux()
}

type Bar struct {
	m string
}
type Baz struct {
	n string
}

func (b Bar) qux() {
	fmt.Println(b.m)
	fmt.Println("ddd")
}

func (b Baz) qux() {
	fmt.Println("baz")
}

func main() {
	var f Foo
	f = Bar{m: "sss"}
	f.qux()
	f = Baz{}
	f.qux()
	// fmt.Println(f.qux())

}
