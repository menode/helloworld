package array

import "fmt"

//切片： 包含三个参数 地址、长度、容量
func TestMap() {
	//使用make创建了一个int切片   长度为3  容量为5
	slice := make([]int, 3, 5)
	fmt.Println(slice)
	//[]中写数字就是数组  不写就是切片 长度和容量一致
	slice1 := []int{10, 20, 30, 40, 50}
	fmt.Println(slice1)
	//新切片共享地址指向的值
	newSlice1 := slice1[1:3]
	fmt.Println(newSlice1)
	//修改新切片
	newSlice1[1] = 33
	fmt.Println(slice1)
	fmt.Println(newSlice1)
	//新切片叠加  老切片会往后替换
	newSlice1 = append(newSlice1, 60)
	fmt.Println(slice1)
	fmt.Println(newSlice1)
	//使用三个索引创建切片
	source := []string{"apple", "orange", "plum", "banana", "grape"}
	slice2 := source[2:3:4]
	fmt.Println(slice2)

}
