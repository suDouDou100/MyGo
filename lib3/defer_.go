package defer_

import "fmt"

func deferFunc() int {
	fmt.Println("defer func called......")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called......")
	return 0
}

func returnAndDefer() int {
	//写入defer关键字
	defer deferFunc()
	return returnFunc()
}

func init() {
	returnAndDefer()

	fmt.Println("init()......")
}
