package main

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	SIZE = 5
)

func main() {
	arr := make([]int, SIZE)
	err := getArray(arr)
	if err != nil {
		panic(err)
	}

	printArray(arr)

	// slicing -> cap은 startIndex 에서 endIndex 까지의 길이가 된다
	arr2 := arr[1:4]
	arr3 := arr[2:5]

	printArray(arr2)
	printArray(arr3)

	// append시 cap안에 공간이 존재할 경우
	arrUsingAppendToInCap := append(arr[2:4], 100)
	printArray(arrUsingAppendToInCap)

	// 기존에 cap을 넘길경우 새로운 메모리공간에 자리를 잡는다
	arrUsingAppendToOverCap := append(arr, 500)
	printArray(arrUsingAppendToInCap)

	arr[1] = 1000
	arr[2] = 1000
	arr[3] = 1000
	arr[4] = 1000

	// 같은 메모리 주소값을 공유한다 -> shallow copy
	printArray(arr2)
	printArray(arr3)
	printArray(arrUsingAppendToInCap)
	printArray(arrUsingAppendToOverCap)

	err = getArray(arr)
	if err != nil {
		panic(err)
	}

}

func getArray(arr []int) error {

	if len(arr) != SIZE {
		msg := "size must be " + strconv.Itoa(SIZE)
		return errors.New(msg)
	}

	for i, v := range []int{1, 2, 3, 4, 5} {
		arr[i] = v
	}

	return nil
}

func printArray(arr []int) {
	fmt.Printf("\narr : %v, len : %d, cap : %d\n", arr, len(arr), cap(arr))
}
