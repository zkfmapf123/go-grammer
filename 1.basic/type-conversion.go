package main

import "fmt"

func main() {

	speed, force := valueInit()
	// speed = speed * force // Error

	// method 1
	speed = speed * int(force) // conversion float64 -> int
	fmt.Println(speed)         // 200

	// method 2
	speed, force = valueInit()
	speed = int(float64(speed) * force)
	fmt.Println(speed)

	/*
		타입 캐스팅의 경우, float64 -> int가 될때 올바른 값이 안나올 수 있다
	*/
}

func valueInit() (int, float64) {
	return 100, 2.5
}
