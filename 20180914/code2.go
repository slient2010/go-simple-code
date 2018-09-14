// 继承例子
package main

import (
	"fmt"
)

type Foo struct {
	baz string
}
type Bar struct {
	Foo
}

func (f *Foo) echo() {
	fmt.Println(f.baz)
}

func main() {
	f := Foo{baz: "Hello, struct"}
	b := Bar{Foo{baz: "Hello, struct"}}
	f.echo()
	b.echo()
}
