// https://leetcode.com/problems/add-two-numbers/

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	currentNode1 := l1
	currentNode2 := l2
	resultNodeHead := &ListNode{}
	carry := 0

	for resultNode := resultNodeHead; currentNode1 != nil || currentNode2 != nil || carry > 0; resultNode = resultNode.Next {
		sum := 0

		if currentNode1 != nil {
			sum += currentNode1.Val
			currentNode1 = currentNode1.Next
		}

		if currentNode2 != nil {
			sum += currentNode2.Val
			currentNode2 = currentNode2.Next
		}

		sum += carry
		carry = sum / 10

		resultNode.Next = &ListNode{Val: sum % 10}
	}

	return resultNodeHead.Next
}
