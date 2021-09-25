package arithmetic

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}
func MiddleNode(head *ListNode) *ListNode {
	mLen, tLen := 0, 0
	curr,middle := head, head
	for curr != nil {
		tLen ++
		curr = curr.Next
		if tLen / 2 > mLen {
			middle = middle.Next
			mLen ++
		}
	}
	return middle
}