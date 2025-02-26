package AddTwoNumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	ll1      *ListNode
	ll2      *ListNode
	expected *ListNode
}

//func LinkedListFromInt(num int) *ListNode {
//	// Handle the case where the number is 0
//	if num == 0 {
//		return &ListNode{Val: 0}
//	}
//
//	var head *ListNode
//	for num > 0 {
//		// Get the last digit of the number
//		digit := num % 10
//		// Remove the last digit from the number
//		num /= 10
//		// Create a new node and point it to the current head
//		newNode := &ListNode{Val: digit}
//		newNode.Next = head
//		// Update the head to the new node
//		head = newNode
//	}
//
//	return head
//}

// createLinkedListFromInt takes an integer and constructs a linked list with digits in reverse order
func LinkedListFromInt(num int) *ListNode {
	// Handle the case where the number is 0
	if num == 0 {
		return &ListNode{Val: 0}
	}

	var head *ListNode
	var current *ListNode
	for num > 0 {
		// Get the last digit of the number
		digit := num % 10
		// Remove the last digit from the number
		num /= 10
		// Create a new node with this digit
		newNode := &ListNode{Val: digit}
		// Append the new node at the beginning of the list
		if head == nil {
			head = newNode
			current = newNode
		} else {
			current.Next = newNode
			current = newNode
		}
	}

	return head
}

func TestAddTwoNumbers(t *testing.T) {
	testCases := []testCase{
		{
			name:     "example1",
			ll1:      LinkedListFromInt(342),
			ll2:      LinkedListFromInt(465),
			expected: LinkedListFromInt(807),
		},
		{
			name:     "adding zeroes",
			ll1:      LinkedListFromInt(0),
			ll2:      LinkedListFromInt(0),
			expected: LinkedListFromInt(0),
		},
		{
			name:     "sum adds to less than 10",
			ll1:      LinkedListFromInt(1),
			ll2:      LinkedListFromInt(2),
			expected: LinkedListFromInt(3),
		},
		{
			name:     "ll1 is empty",
			ll2:      LinkedListFromInt(9),
			expected: LinkedListFromInt(9),
		},
		{
			name:     "ll2 is empty",
			ll1:      LinkedListFromInt(9),
			expected: LinkedListFromInt(9),
		},
		{
			name:     "sum adds to more than 10",
			ll1:      LinkedListFromInt(5),
			ll2:      LinkedListFromInt(6),
			expected: LinkedListFromInt(11),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := addTwoNumbers(tc.ll1, tc.ll2)

			assert.NotNil(t, actual)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
