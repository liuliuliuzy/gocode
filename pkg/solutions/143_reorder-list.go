package solutions

func reorderList(head *ListNode) {
	// define three function closures

	// reverse linkded list
	var reverseList = func(head *ListNode) *ListNode {
		var previous *ListNode = nil
		current := head
		for current != nil {
			tmp := current.Next
			current.Next = previous
			previous = current
			current = tmp
		}
		return previous
	}
	// find the middle node of a linked list
	var middleList = func(head *ListNode) *ListNode {
		slow, fast := head, head
		for fast.Next != nil && fast.Next.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		return slow
	}
	// merge two linkded list
	var mergeList = func(h1, h2 *ListNode) {
		h1Tmp, h2Tmp := h1, h2
		for h1 != nil && h2 != nil {
			h1Tmp = h1.Next
			h2Tmp = h2.Next
			h1.Next = h2
			h1 = h1Tmp
			h2.Next = h1
			h2 = h2Tmp
		}
	}
	if head == nil {
		return
	}
	mid := middleList(head)
	h1 := head
	h2 := mid.Next
	mid.Next = nil
	h2 = reverseList(h2)
	mergeList(h1, h2)
}

func ReorderList(head *ListNode) {
	reorderList(head)
}
