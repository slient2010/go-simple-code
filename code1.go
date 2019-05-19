// 封装例子
package main

import "fmt"

type Foo struct {
	baz string
}

func (f *Foo) echo() {
	fmt.Println(f.baz)
}

func main() {
	f := Foo{baz: "Hello, struct!"}
	f.echo()
	fmt.Println("vim-go")
}
