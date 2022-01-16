package solutions

import "math/rand"

type Solution struct {
	length   int
	elements []int
}

func Constructor(head *ListNode) Solution {
	s := Solution{length: 0, elements: []int{}}
	for head != nil {
		s.elements = append(s.elements, head.Val)
		s.length++
		head = head.Next
	}
	return s
}

func (s *Solution) GetRandom() int {
	// generate random step
	sample := rand.Intn(s.length)
	return s.elements[sample]
}
