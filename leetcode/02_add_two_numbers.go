package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		root  = &ListNode{}
		cur   = root
		upper int
	)

	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + upper
		upper = val / 10
		val %= 10
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		val := l1.Val + upper
		upper = val / 10
		val %= 10
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
		l1 = l1.Next
	}
	for l2 != nil {
		val := l2.Val + upper
		upper = val / 10
		val %= 10
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
		l2 = l2.Next
	}
	if upper > 0 {
		cur.Next = &ListNode{Val: upper}
	}
	return root.Next
}
