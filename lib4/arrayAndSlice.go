package arrayAndSlice

import "fmt"

func init() {
	// 定义固定长度的数组
	// var myArray1 [10]int

	myArray2 := [10]int{1, 2, 3, 4}
	// for i := 0; i < 10; i++ {
	for i := 0; i < len(myArray2); i++ {
		fmt.Println(myArray2[i])
	}
}
