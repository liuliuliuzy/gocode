package datastructures

import "fmt"

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func ToLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{nums[0], nil}
	tmp := head
	for _, num := range nums[1:] {
		node := ListNode{num, nil}
		tmp.Next = &node
		tmp = tmp.Next
	}
	return head
}

func ShowLinkedList(head *ListNode) {
	help := head
	for help != nil {
		if help.Next == nil {
			fmt.Printf("%d\n", help.Val)
			break
		}
		fmt.Printf("%d->", help.Val)
		help = help.Next
	}
}
