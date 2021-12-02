/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var num1 = 0
	var num2 = 0

	var lreturn *ListNode = &ListNode{0, nil}
	var lreturn_head = lreturn
	var total = 0

	for l1 != nil || l2 != nil {

		if l1 != nil {
			num1 += l1.Val
			if l1.Next != nil {
				num1 *= 10
			}
			l1 = l1.Next
		}

		if l2 != nil {
			num2 += l2.Val
			if l2.Next != nil {
				num2 *= 10
			}
			l2 = l2.Next
		}
	}
	total = num1 + num2

	for total != 0 {
		lreturn.Val = total % 10
		total /= 10

		if total != 0 {
			lreturn.Next = &ListNode{0, nil}
			lreturn = lreturn.Next
		}
	}
	return lreturn_head
}