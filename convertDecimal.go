package main

import (
	"fmt"
	"strconv"
)

func getDecimalValue(head *ListNode) int {
	var decString string = ""

	for head != nil {
		decString += fmt.Sprintf("%d", head.Val)
		head = head.Next
	}

	number, _ := strconv.ParseInt(decString, 10, 64)

	return int(number)
}
