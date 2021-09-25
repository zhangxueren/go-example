package main

import (
	"fmt"
	"go-example/v0/arithmetic"
)

func main() {
	head := &arithmetic.ListNode{Val:1, Next:&arithmetic.ListNode{Val:2, Next:&arithmetic.ListNode{Val:3, Next:&arithmetic.ListNode{Val:4, Next:nil}}}}

	fmt.Println(arithmetic.MiddleNode(head))
}

