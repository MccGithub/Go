package main

import (
	"fmt"
	"github.com/chive-chan/Go/Structure"
)

// 反转一个单链表。
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL

func reverseList(head *Structure.MyList) *Structure.MyList {
	if head == nil || head.Next == nil {
		return nil
	}

	var pCur *Structure.MyList  // 当前节点
	var pNext *Structure.MyList // 后继结点

	pCur = head.Next
	head.Next = nil
	for {
		pNext = pCur.Next
		pCur.Next = head.Next
		head.Next = pCur
		pCur = pNext
		if pCur == nil {
			break
		}
	}
	return head
}

func main() {
	//l := Structure.NewListByDatas(1, 2.8, "string", []int{1, 2, 3}, []string{"A", "B", "C"})
	l := Structure.NewListByDatas("A", "B", "C", "D", "E")
	fmt.Println(l)
	fmt.Println(reverseList(l))
}
