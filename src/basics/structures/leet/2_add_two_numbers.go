/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var carry = 0

	var laux *ListNode = &ListNode{0, nil}
	var lreturn_head = laux

	for l1 != nil || l2 != nil {
		laux.Val = carry

		if l1 != nil {
			laux.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
			laux.Val += l2.Val
		}

		if laux.Val >= 10 {
			carry = laux.Val / 10
			laux.Val %= 10
		} else {
			carry = 0
		}

		if l1 != nil || l2 != nil {
			laux.Next = &ListNode{0, nil}
			laux = laux.Next
		}
	}

	if carry != 0 {
		laux.Next = &ListNode{carry, nil}
	}

	return lreturn_head
}