package Structure

import "fmt"

type MyList struct {
	Elem interface{}
	Next *MyList
}

func NewList() *MyList {
	head := new(MyList)
	head.Elem = 0
	head.Next = nil
	return head
}

func (l *MyList) Add(elem interface{}) *MyList {
	if l == nil {
		_ = fmt.Errorf("%s\n", "You need to create a linked list before adding elements. Call the New method to get an empty linked list containing the head node.")
		return nil
	}
	a := new(MyList)
	a.Elem = elem
	a.Next = l.Next
	l.Next = a
	if v, ok := l.Elem.(int); ok == true {
		l.Elem = v + 1
		return l
	} else {
		_ = fmt.Errorf("%s\n", "The head node element is wrong, please make sure that the element of the head node is of type int.")
		return nil
	}
}

func NewListByDatas(elems ...interface{}) *MyList {
	head := NewList()
	for i := len(elems) - 1; i >= 0; i-- {
		head.Add(elems[i])
	}
	return head
}

func (l *MyList) String() string {
	s := "MyList[len: "
	s = s + fmt.Sprintf("%d", l.Elem.(int)) + "; Elements: "
	for e := l.Next; e != nil; e = e.Next {
		s = s + fmt.Sprintf("%v", e.Elem) + ", "
	}
	s = s[:len(s)-2] + "]"
	return s
}
