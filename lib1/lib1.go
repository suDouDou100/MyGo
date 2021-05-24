package lib1

import "fmt"

func changeValue(p *int) { //*表示p是一个指针类型 *int表示指向整形的指针类型
	*p = 10 //*p表示改变当前p所指向的空间的指
}

func init() {
	var a int = 1

	changeValue(&a) //&a表示我传的是a的地址

	fmt.Println("a is ", a)
}
