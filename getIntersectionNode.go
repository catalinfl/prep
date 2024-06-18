package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	seen := make(map[*ListNode]bool)

	for n := headA; n != nil; n = n.Next {
		seen[n] = true
	}

	for n := headB; n != nil; n = n.Next {
		if seen[n] {
			return n
		}
	}

	return nil

}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	tmpA := headA
	tmpB := headB

	var len int

	for tmpA != nil {
		tmpA = tmpA.Next
		len++
	}

	for tmpB != nil {
		tmpB = tmpB.Next
		len--
	}

	if len > 0 {
		tmpA = headA
		tmpB = headB
	} else {
		tmpA = headB
		tmpB = headA
		len = -len
	}

	for i := 0; i < len; i++ {
		tmpA = tmpA.Next
	}

	for tmpA != tmpB {
		tmpA = tmpA.Next
		tmpB = tmpB.Next
	}

	return tmpA
}
