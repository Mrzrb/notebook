//package main

import "fmt"

/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 *
 * https://leetcode.com/problems/linked-list-cycle-ii/description/
 *
 * algorithms
 * Medium (38.84%)
 * Total Accepted:    411.6K
 * Total Submissions: 1M
 * Testcase Example:  '[3,2,0,-4]\n1'
 *
 * Given a linked list, return the node where the cycle begins. If there is no
 * cycle, return null.
 *
 * There is a cycle in a linked list if there is some node in the list that can
 * be reached again by continuously following the next pointer. Internally, pos
 * is used to denote the index of the node that tail's next pointer is
 * connected to. Note that pos is not passed as a parameter.
 *
 * Notice that you should not modify the linked list.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [3,2,0,-4], pos = 1
 * Output: tail connects to node index 1
 * Explanation: There is a cycle in the linked list, where tail connects to the
 * second node.
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1,2], pos = 0
 * Output: tail connects to node index 0
 * Explanation: There is a cycle in the linked list, where tail connects to the
 * first node.
 *
 *
 * Example 3:
 *
 *
 * Input: head = [1], pos = -1
 * Output: no cycle
 * Explanation: There is no cycle in the linked list.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of the nodes in the list is in the range [0, 10^4].
 * -10^5 <= Node.val <= 10^5
 * pos is -1 or a valid index in the linked-list.
 *
 *
 *
 * Follow up: Can you solve it using O(1) (i.e. constant) memory?
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//type ListNode struct {
//Val  int
//Next *ListNode
//}

func detectCycle(head *ListNode) *ListNode {
	// 这时候没有环
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head

	// 先判断有没有环
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			index1 := head
			for index1 != slow {
				index1 = index1.Next
				slow = slow.Next
			}
			return index1
		}
	}

	return nil
}

// =================================== 下面的是暴力解法

var slowNode = &ListNode{}
var fastNode = &ListNode{}

var tailNode = &ListNode{}
var is = false

func detectCycle1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	p := head
	for p != nil {
		if isCircle, h := check(p); isCircle && h == p {
			fmt.Printf("isCircle %s, node is %+v", isCircle, h)
			return h
		}
		p = p.Next
	}
	return nil
}

// check return first boll indicate weither list is a circle, second  returns head if circle list start with head
func check(head *ListNode) (bool, *ListNode) {
	if head == nil || head.Next == nil {
		return false, nil
	}

	if head.Next == head {
		return true, head
	}

	slowNode = head.Next
	fastNode = head.Next.Next

	for fastNode != slowNode && fastNode != nil && fastNode.Next != nil {
		if fastNode.Next == head {
			return true, head
		}
		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next
		if fastNode == head {
			return true, head
		}
	}

	if slowNode == fastNode {
		return true, nil
	}

	return false, nil
}
