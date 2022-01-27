package solutions

// 我自己的解法，击败go 100% + 100%
// 还行吧 =。=
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// find the first helper pointer
	// we need a dumb node here
	helper := ListNode{Val: 0, Next: head}
	iterator := &helper
	for iterator.Next != nil && iterator.Next.Val < x {
		iterator = iterator.Next
	}
	// now move on to find out the node with smaller value
	h2 := iterator.Next
	if h2 == nil {
		return head
	}
	h3 := h2.Next
	for h3 != nil {
		if h3.Val < x {
			h2.Next = h3.Next
			h3.Next = iterator.Next
			iterator.Next = h3
			iterator = h3
			h3 = h2.Next
		} else {
			h2 = h3
			h3 = h3.Next
		}
	}
	return helper.Next
}

func partitionII(head *ListNode, x int) *ListNode {
	// store the small & large nodes
	small := &ListNode{}
	large := &ListNode{}
	smallHead := small
	largeHead := large
	for head != nil {
		// iterate every node of the linked list
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
		// now head is always faster than small and large pointer
	}
	// cut out the large.Next
	large.Next = nil
	small.Next = largeHead.Next
	return smallHead.Next
}

func Partition(head *ListNode, x int) *ListNode {
	return partition(head, x)
}
