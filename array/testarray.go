package array

import "fmt"

func TestArray() {
	arr := [5]int{10, 20, 30, 40, 50}
	fmt.Println(arr)

	arr1 := [5]*int{0: new(int), 1: new(int)}
	*arr1[0] = 10
	*arr1[1] = 20
	fmt.Println(*arr1[0])

	var arr2_1 [5]string
	arr2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	arr2_1 = arr2
	fmt.Println(arr2_1)

}
