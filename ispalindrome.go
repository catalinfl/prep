package main

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// if fast != nil {
	// 	slow = slow.Next
	// }

	list2 := reverse(slow)
	list1 := head

	for list2 != nil {
		if list1.Val != list2.Val {
			return false
		}

		list1 = list1.Next
		list2 = list2.Next
	}

	return true

}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}

	return prev
}
