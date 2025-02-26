package AddTwoNumbers

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order,
and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list. You may assume
the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]


Constraints:

* The number of nodes in each linked list is in the range [1, 100].
* 0 <= Node.val <= 9
* It is guaranteed that the list represents a number that does not have leading zeros.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersWithCarry(l1, l2, 0)
}

func addTwoNumbersWithCarry(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	var answer ListNode
	var sum int
	var l1Next *ListNode
	var l2Next *ListNode
	switch {
	case l1 == nil && l2 == nil:
		if carry != 0 {
			answer.Val = carry
			return &answer
		}
		return nil
	case l1 == nil:
		sum = l2.Val + carry
		l2Next = l2.Next
		//modulo := sum % 10
		//answer.Val = modulo
		//answer.Next = addTwoNumbersWithCarry(l1.Next, nil, carryover)
	case l2 == nil:
		sum = l1.Val + carry
		l1Next = l1.Next
	default:
		sum = l1.Val + l2.Val + carry
		l1Next = l1.Next
		l2Next = l2.Next
	}
	modulo := sum % 10
	carryover := sum / 10
	answer.Val = modulo
	answer.Next = addTwoNumbersWithCarry(l1Next, l2Next, carryover)

	return &answer
}
