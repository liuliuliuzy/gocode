package solutions

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// first step: calculate the length of the linked list
	n := 1
	tail := head
	for tail.Next != nil {
		n++
		tail = tail.Next
	}
	// fmt.Println(n)
	tail.Next = head

	k = n - (k % n) - 1
	h1 := head
	for k > 0 {
		h1 = h1.Next
		k--
	}
	h2 := h1.Next
	h1.Next = nil
	return h2
}
func RotateRight(head *ListNode, k int) *ListNode {
	return rotateRight(head, k)
}
