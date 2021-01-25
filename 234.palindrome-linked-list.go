/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 *
 * https://leetcode.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (39.88%)
 * Total Accepted:    543.5K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,2]'
 *
 * Given a singly linked list, determine if it is a palindrome.
 *
 * Example 1:
 *
 *
 * Input: 1->2
 * Output: false
 *
 * Example 2:
 *
 *
 * Input: 1->2->2->1
 * Output: true
 *
 * Follow up:
 * Could you do it in O(n) time and O(1) space?
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	nomal := []int{}
	reverseL := []int{}

	c := head
	for c != nil {
		nomal = append(nomal, c.Val)
		c = c.Next
	}
	reverseList := reverse(head)
	r := reverseList
	c = r
	for c != nil {
		reverseL = append(reverseL, c.Val)
		c = c.Next
	}

	for idx, v := range nomal {
		if v != reverseL[idx] {
			return false
		}
	}

	return true
}

func reverse(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	last := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
