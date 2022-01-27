package solutions

//node是要被删除的节点
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func DeleteNode(node *ListNode) {
	deleteNode(node)
}
