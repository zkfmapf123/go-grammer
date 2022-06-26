package main

import "fmt"

func main() {

	basicArray()
	sliceArray(10)
	shallowCpy()
	deepCpy()
}

// 명시적 선언
func basicArray() {
	var arr [4]int
	for i, v := range []int{1, 2, 3, 4} {
		arr[i] = v
	}

	fmt.Println(arr)
}

// 가변길이 배열
func sliceArray(len int) {
	// type, len, cap
	arr := make([]int, len, len) // exists default value

	for i, v := range []int{1, 2, 3, 4, 5} {
		arr[i] = v
	}

	fmt.Println(arr)
}

// shallow copy
func shallowCpy() {
	fmt.Println("\nShallow copy")
	arr := make([]int, 5, 5)

	for i, v := range []int{1, 2, 3, 4, 5} {
		arr[i] = v
	}

	arrCpy1 := arr[1:3]
	arrCpy2 := arr[2:5]

	fmt.Println(arr)
	fmt.Println(arrCpy1)
	fmt.Println(arrCpy2)

	arr[1] = 100
	arr[2] = 200
	arr[3] = 300

	fmt.Println(arr)
	fmt.Println(arrCpy1)
	fmt.Println(arrCpy2)
}

// deep copy
func deepCpy() {
	fmt.Println("\nDepp copy")
	arr := make([]int, 5, 5)

	for i, v := range []int{1, 2, 3, 4, 5} {
		arr[i] = v
	}

	arrCpy1 := make([]int, 2, 2)
	arrCpy2 := make([]int, 3, 3)
	copy(arrCpy1, arr[1:3])
	copy(arrCpy2, arr[2:5])

	fmt.Println(arr)
	fmt.Println(arrCpy1)
	fmt.Println(arrCpy2)

	arr[1] = 100
	arr[2] = 200
	arr[3] = 300

	fmt.Println(arr)
	fmt.Println(arrCpy1)
	fmt.Println(arrCpy2)

}
