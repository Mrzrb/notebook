
/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 *
 * https://leetcode.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (43.49%)
 * Total Accepted:    320.2K
 * Total Submissions: 721.8K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, reverse the nodes of a linked list k at a time and
 * return its modified list.
 *
 * k is a positive integer and is less than or equal to the length of the
 * linked list. If the number of nodes is not a multiple of k then left-out
 * nodes, in the end, should remain as it is.
 *
 * Follow up:
 *
 *
 * Could you solve the problem in O(1) extra memory space?
 * You may not alter the values in the list's nodes, only nodes itself may be
 * changed.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4,5], k = 2
 * Output: [2,1,4,3,5]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1,2,3,4,5], k = 3
 * Output: [3,2,1,4,5]
 *
 *
 * Example 3:
 *
 *
 * Input: head = [1,2,3,4,5], k = 1
 * Output: [1,2,3,4,5]
 *
 *
 * Example 4:
 *
 *
 * Input: head = [1], k = 1
 * Output: [1]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range sz.
 * 1 <= sz <= 5000
 * 0 <= Node.val <= 1000
 * 1 <= k <= sz
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type reverse func(head *ListNode) *ListNode

//actually applied reverse function
var reverseFn = reverseRecursion

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	n := head
	h := head
	for i := 0; i < k; i++ {
		if n == nil {
			return head
		}
		n = n.Next
	}
	newHead := reverseBetween(h, n)
	h.Next = reverseKGroup(n, k)
	return newHead
}

func reverseBetween(head *ListNode, n *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == n {
		return head
	}
	last := reverseBetween(head.Next, n)
	head.Next.Next = head
	head.Next = nil
	return last
}

// reverse function applied via recursion
func reverseRecursion(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	last := reverseRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

//func reverseIter(head *ListNode) *ListNode {
//var prev, cur, next *ListNode
//cur = head
//next = head
//}
