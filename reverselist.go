package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		memoNext := head.Next
		head.Next = prev // leaga previous la head
		prev = head      // se salveaza in lista
		head = memoNext  // memoNext => trece pe urmatoarea pozitie
	}

	return prev

}
