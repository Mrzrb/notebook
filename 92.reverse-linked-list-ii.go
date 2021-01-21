/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 *
 * https://leetcode.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (39.76%)
 * Total Accepted:    320.9K
 * Total Submissions: 796.8K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * Reverse a linked list from position m to n. Do it in one-pass.
 *
 * Note: 1 ≤ m ≤ n ≤ length of list.
 *
 * Example:
 *
 *
 * Input: 1->2->3->4->5->NULL, m = 2, n = 4
 * Output: 1->4->3->2->5->NULL
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

var p *ListNode

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == 1 {
		return reverseN(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	last := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		p = head.Next
		return head
	}

	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = p
	return last
}
