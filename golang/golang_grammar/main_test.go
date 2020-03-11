package test

import (
	"fmt"
	"testing"
)

func TestBreak(t *testing.T) {
	i := 0
JLoop:

	fmt.Println("i: ", i)
	for ; i < 10; i++ {
		if i > 5 {
			goto JLoop
		}
		fmt.Println("i: ", i)
	}
	//fmt.Println("i: ", i)
}

func TestFunc(t *testing.T) {

	func() {
		fmt.Println("匿名函数")
	}()

	n := 10
	numberFunc := func(number *int) {
		fmt.Println("number: ", *number)
		*number += 20
	}
	numberFunc(&n)
	fmt.Println(n)

	i := 20
	func() {
		fmt.Println("i: ", i)
		i += 20
	}()

	fmt.Println("i: ", i)
}
