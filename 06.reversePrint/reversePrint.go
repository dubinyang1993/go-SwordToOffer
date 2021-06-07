package main

import (
	"container/list"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	b := ListNode{
		2,
		nil,
	}
	a := ListNode{
		1,
		&b,
	}
	fmt.Println(reversePrint(&a), a.Next, b.Next) // 因为翻转链表，ab数据发生变化
	// fmt.Println(reversePrint4(&a), a.Next, b.Next) // 只是遍历链表，ab数据不变
}

// 翻转链表，遍历数组
func reversePrint(head *ListNode) []int {
	var newListNode *ListNode
	for head != nil {
		head, head.Next, newListNode = head.Next, newListNode, head
	}
	var list []int
	for newListNode != nil {
		list = append(list, newListNode.Val)
		newListNode = newListNode.Next
	}
	return list
}

// 遍历链表，翻转数组
func reversePrint2(head *ListNode) []int {
	var list []int
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	for i, j := 0, len(list)-1; i < j; {
		list[i], list[j] = list[j], list[i]
		i++
		j--
	}
	return list
}

// 利用list，后进先出
func reversePrint3(head *ListNode) []int {
	l := list.New()
	for head != nil {
		l.PushFront(head.Val)
		head = head.Next
	}

	var list []int
	for e := l.Front(); e != nil; e = e.Next() {
		list = append(list, e.Value.(int))
	}
	return list
}

// 一次遍历计数，二次遍历根据索引给数组赋值
func reversePrint4(head *ListNode) []int {
	var count int
	newHead := head
	for head != nil {
		count++
		head = head.Next
	}

	result := make([]int, count)
	i := 0
	for newHead != nil {
		result[count-i-1] = newHead.Val
		i++
		newHead = newHead.Next
	}
	return result
}
