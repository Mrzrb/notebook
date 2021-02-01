/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 *
 * https://leetcode.com/problems/sort-list/description/
 *
 * algorithms
 * Medium (45.16%)
 * Total Accepted:    333.1K
 * Total Submissions: 724.2K
 * Testcase Example:  '[4,2,1,3]'
 *
 * Given the head of a linked list, return the list after sorting it in
 * ascending order.
 *
 * Follow up: Can you sort the linked list in O(n logn) time and O(1) memory
 * (i.e. constant space)?
 *
 *
 * Example 1:
 *
 *
 * Input: head = [4,2,1,3]
 * Output: [1,2,3,4]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [-1,5,3,4,0]
 * Output: [-1,0,3,4,5]
 *
 *
 * Example 3:
 *
 *
 * Input: head = []
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range [0, 5 * 10^4].
 * -10^5 <= Node.val <= 10^5
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

// 时间复杂度 O(n^2)
func sortList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h := head

	sorted := sortList(head.Next)
	p := sorted
	prev := sorted

	for p != nil && head.Val > p.Val {
		h = sorted
		prev = p
		p = p.Next
	}

	if p != sorted {
		head.Next = p
		prev.Next = head
	} else {
		h.Next = sorted
	}

	return h
}

var dummyNode = &ListNode{}

//时间复杂度O(nlogn)
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { // 递归的出口，不用排序 直接返回
		return head
	}
	slow, fast := head, head // 快慢指针
	var preSlow *ListNode    // 保存slow的前一个结点
	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow = slow.Next      // 慢指针走一步
		fast = fast.Next.Next // 快指针走两步
	}
	preSlow.Next = nil  // 断开，分成两链
	l := sortList(head) // 已排序的左链
	r := sortList(slow) // 已排序的右链
	return merge(l, r)  // 合并已排序的左右链，一层层向上返回
}

func merge(l1, l2 *ListNode) *ListNode {
	defer func() {
		dummyNode.Next = nil
	}()
	prev := dummyNode

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	if l1 != nil {
		prev.Next = l1
	}

	if l2 != nil {
		prev.Next = l2
	}
	return dummyNode.Next
}
