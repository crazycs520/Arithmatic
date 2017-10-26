/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
*/
package main

import "fmt"

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	c := l1.Val + l2.Val
	result = &ListNode{Val: c % 10}
	c = c / 10
	res := result
	l1 = l1.Next
	l2 = l2.Next
	x := 0
	y := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			x = l1.Val
		} else {
			x = 0
		}
		if l2 != nil {
			y = l2.Val
		} else {
			y = 0
		}

		result.Next = &ListNode{}

		c = x + y + c
		result.Next.Val = c % 10
		c = c / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		result = result.Next
	}
	if c != 0 {
		result.Next = &ListNode{}
		result.Next.Val = c
	}
	return res
}

func makeListNode(data []int) *ListNode {
	var result = &ListNode{}
	res := result
	for i, v := range data {
		result.Val = v

		if i != len(data)-1 {
			result.Next = &ListNode{}
			result = result.Next
		}

	}
	return res
}

func printListNode(l1 *ListNode) {
	fmt.Printf("%d", l1.Val)
	l1 = l1.Next
	for l1 != nil {
		fmt.Printf("-->%d", l1.Val)
		l1 = l1.Next
	}
	fmt.Println()
}

func main() {
	l1 := makeListNode([]int{1, 2, 3})
	l2 := makeListNode([]int{9, 8, 7, 9})

	printListNode(l1)
	printListNode(l2)
	printListNode(addTwoNumbers(l1, l2))
}
