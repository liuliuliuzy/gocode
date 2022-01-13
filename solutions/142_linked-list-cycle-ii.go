package solutions

/*
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
*/

// method1: still hash table
// func detectCycle(head *ListNode) *ListNode {
// 	visited := make(map[*ListNode]bool) // key type is *ListNode
// 	for head != nil {
// 		if visited[head] {
// 			return head
// 		}
// 		visited[head] = true
// 		head = head.Next
// 	}
// 	return nil
// }

// two pointers
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		// check if has cycle
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		// add the third pointer to find out the meet node
		if slow == fast {
			walker := head
			for walker != slow {
				walker = walker.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}

func DetectCycle(head *ListNode) *ListNode {
	return detectCycle(head)
}
