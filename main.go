package main

import (
	"fmt"
	doublelist "strukdat/double_list"
	singlelist "strukdat/single_list"
)

func main() {
	// Single List Example
	singlelist.PushHead(112233, "Jakarta")
	singlelist.PushHead(123321, "Bandung")
	singlelist.PushTail(332113, "Cianjur")

	singlelist.Print()

	singlelist.PopTail()
	fmt.Println()
	singlelist.Print()

	fmt.Println()
	// Double List Example
	doublelist.PushHead(112233, "Jakarta")
	doublelist.PushHead(123321, "Bandung")
	doublelist.PushTail(332113, "Cianjur")

	doublelist.PrintFromHead()
	fmt.Println()
	doublelist.PrintFromTail()
	fmt.Println()

	doublelist.PopTail()
	doublelist.PrintFromHead()
}
