package main

func removeElements(head *ListNode, val int) *ListNode {
	var res *ListNode

	for head != nil {

		if head.Val == val {
			head = head.Next // de cate ori se intalneste valoarea la inceput
		} else {
			res = head
			break
		}
	}

	current := res // se foloseste referinta la res ca sa avem un capat, current se inlocuieste

	for current != nil && current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return res
}
