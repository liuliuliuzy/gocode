package solutions

/*
给你一个链表的头节点 head ，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

如果链表中存在环，则返回 true 。 否则，返回 false 。
*/

// method1: hash table
// func hasCycle(head *ListNode) bool {
// 	visited := make(map[*ListNode]bool) // key type is *ListNode
// 	for head != nil {
// 		if visited[head] {
// 			return true
// 		}
// 		visited[head] = true
// 		head = head.Next
// 	}
// 	return false
// }

// method2: two pointers
// aka floyd cycle detection algorithm: https://en.wikipedia.org/wiki/Cycle_detection
func hasCycle(head *ListNode) bool {
	// special situation
	if head == nil || head.Next == nil {
		return false
	}
	// define two pointers
	slow, fast := head, head.Next
	// now let them move in different speed
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func HasCycle(head *ListNode) bool {
	return hasCycle(head)
}
