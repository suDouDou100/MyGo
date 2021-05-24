//指针实例,交换两个变量的值

package pointer

import "fmt"

/*
func swap(a int, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}
*/

func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

func init() {
	var a int = 10
	var b int = 200

	swap(&a, &b)

	fmt.Println("a is", a, "\t b is", b)

	var p *int = &a

	fmt.Println("&a is", &a)
	fmt.Println("p is", p)

	var pp **int = &p //二级指针

	fmt.Println("pp is", pp)
}
