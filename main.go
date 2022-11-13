package main

import (
	"fmt"
	"strukdat/datastruct"
)

func main() {
	// Single List Example
	singleList := datastruct.NewSingleList().
		PushHead("Muruyung", 23).
		PushHead("Sir", 17).
		PushTail("Naufal", 22)

	singleList.Print()

	singleList = singleList.PopTail()
	fmt.Println()
	singleList.Print()

	fmt.Println()

	// Double List Example
	doubleList := datastruct.NewDoubleList().
		PushHead("Muruyung", 23).
		PushHead("Sir", 17).
		PushTail("Naufal", 22)

	doubleList.PrintFromHead()
	fmt.Println()
	doubleList.PrintFromTail()
	fmt.Println()

	doubleList = doubleList.PopTail()
	doubleList.PrintFromHead()
}
