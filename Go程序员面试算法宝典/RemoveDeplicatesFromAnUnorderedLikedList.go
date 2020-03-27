//package main
//
//import (
//	"fmt"
//	"github.com/chive-chan/Go/Structure" // 引入自定义的数据结构
//	"sync"
//)
//
//// 从无序链表中移除重复项
////
////给定一个没有排序的链表， 去掉其重复项， 并保留原顺序。
////示例:
////原链表: 1->3->1->5->5->7
////移除后: 1->3->5->7
//func RemoveDeplicatesLikedList(head *Structure.MyList) *Structure.MyList {
//	if head == nil || head.Next == nil {
//		return nil
//	}
//	outerCur := head.Next          //用于外层循环, 指向链表的第一个节点
//	var innerCur *Structure.MyList //用于内层循环遍历 outerCur 后面的节点
//	var innerPre *Structure.MyList //innerCur的前驱节点
//	for ; outerCur != nil; outerCur = outerCur.Next {
//		for innerCur, innerPre = outerCur.Next, outerCur; innerCur != nil; innerCur = innerCur.Next {
//			if outerCur.Elem == innerCur.Elem {
//				innerPre.Next = innerCur.Next
//			} else {
//				innerPre = innerCur
//			}
//		}
//	}
//	return head
//}
//
//func main() {
//	//head := Structure.NewListByDatas(1, 3, 1, 5, 5, 7)
//	//fmt.Println(head)
//	//fmt.Println(RemoveDeplicatesLikedList(head))
//	wg := sync.WaitGroup{}
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go func() {
//			defer fmt.Println(i)
//			wg.Done()
//		}()
//	}
//	wg.Wait()
//}
package main

import (
	"fmt"
	"sort"
)

func Find(arr []int, num int) bool {
	sort.Ints(arr)
	fmt.Println(arr)
	size := len(arr)
	i := 0, j = size-1
	for  i != j {
		if arr[i] + arr[j] < num {
			i++
		} else if arr[i] + arr[j] > num {
			j--
		} else if arr [i] + arr[j] == num{
			return true
		}
	}
	return false
}

func main() {
	//a := 0
	//fmt.Scan(&a)
	//fmt.Printf("%d\n", a)
	//fmt.Printf("Hello World!\n");
	arr := []int{9, 7, -2, 3, 5}
	fmt.Println(Find(arr, 8))
	fmt.Println(Find(arr, 5))
	fmt.Println(Find(arr, 100))
}
