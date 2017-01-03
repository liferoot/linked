package linked

import (
	"fmt"
	"testing"
)

func TestList_New(t *testing.T) {
	exp := []interface{}{1, 2, 3}
	l := NewList()
	chkempty(t, ``, l)
	l.Add(exp...)
	chklist(t, ``, l, exp)
	l.Clear()
	chkempty(t, ``, l)
}

func TestList_Add(t *testing.T) {
	l := new(List)
	if n := l.Add(); n != nil {
		t.Errorf("expected node %+v, got %+v", nil, n)
	}
	l.Add(1)
	l.Add(NewNode(2))
	chklist(t, `case 2`, l, []interface{}{1, 2})
	l.Add(3, &Node{Value: 4}, 5)
	chklist(t, `case 3`, l, []interface{}{1, 2, 3, 4, 5})
}

func TestList_AddFromAnotherList(t *testing.T) {
	l := NewList(1, 2, 3)
	ll := NewList(4, 5)
	l.Add(ll.First())
	chklist(t, `case 1/l`, l, []interface{}{1, 2, 3, 4})
	chklist(t, `case 1/ll`, ll, []interface{}{5})
	l.Add(ll.Last())
	chklist(t, `case 2/l`, l, []interface{}{1, 2, 3, 4, 5})
	chkempty(t, `case 2/ll`, ll)
	ll.Add(l.First(), l.Last())
	chklist(t, `case 3/l`, l, []interface{}{2, 3, 4})
	chklist(t, `case 3/ll`, ll, []interface{}{1, 5})
}

func TestList_AddAfter(t *testing.T) {
	exp := []interface{}{1, &Node{Value: 2}, 3, 4, NewNode(5)}
	l := new(List)
	m := l.AddAfter(nil)
	if m != nil {
		t.Errorf("case 1: expected node %+v, got %+v", nil, m)
	}
	if m = l.AddAfter(new(Node)); m != nil {
		t.Errorf("case 2: expected node %+v, got %+v", nil, m)
	}
	m = l.AddAfter(nil, exp[0])
	chklist(t, `case 3`, l, exp[:1])
	m = l.AddAfter(m, m, exp[1])
	chklist(t, `case 4`, l, exp[:2])
	l.AddAfter(m, exp[2:]...)
	chklist(t, `case 5`, l, exp)
}

func TestList_AddBefore(t *testing.T) {
	exp := []interface{}{5, &Node{Value: 4}, 3, NewNode(2), 1}
	l := new(List)
	m := l.AddBefore(nil)
	if m != nil {
		t.Errorf("case 1: expected node %+v, got %+v", nil, m)
	}
	if m = l.AddBefore(new(Node)); m != nil {
		t.Errorf("case 2: expected node %+v, got %+v", nil, m)
	}
	m = l.AddBefore(nil, exp[0])
	chklist(t, `case 3`, l, exp[:1])
	m = l.AddBefore(m, exp[1])
	chklist(t, `case 4`, l, []interface{}{4, 5})
	l.AddBefore(m, exp[2:]...)
	chklist(t, `case 5`, l, []interface{}{3, 2, 1, 4, 5})
}

func TestList_Pop(t *testing.T) {
	l := NewList()
	if v := l.Pop(); v != nil {
		t.Errorf("expected value %v, got %d", nil, v)
	}
	l.Add(1, 2, 3)
	if v := l.Pop(); v != 1 {
		t.Errorf("expected value %d, got %d", 1, v)
	}
	chklist(t, ``, l, []interface{}{2, 3})
	if v := l.Pop(); v != 2 {
		t.Errorf("expected value %d, got %d", 2, v)
	}
	chklist(t, ``, l, []interface{}{3})
}

func TestList_Push(t *testing.T) {
	l := new(List)
	l.Push(3)
	chklist(t, `case 1`, l, []interface{}{3})
	l.Push(NewNode(2))
	l.Push(1)
	chklist(t, `case 2`, l, []interface{}{1, 2, 3})
}

func TestList_PushFromAnotherList(t *testing.T) {
	l, ll := new(List), NewList(1, 2, 3)
	l.Push(ll.First())
	chklist(t, `case 1/l`, l, []interface{}{1})
	chklist(t, `case 1/ll`, ll, []interface{}{2, 3})
	l.Push(ll.First())
	chklist(t, `case 2/l`, l, []interface{}{2, 1})
	chklist(t, `case 2/ll`, ll, []interface{}{3})
}

func TestList_Remove(t *testing.T) {
	l := new(List)
	if v := l.Remove(); v != nil {
		t.Errorf("expected value %v, got %v", nil, v)
	}
	l.Add(1, 2, 3, 4, 5)
	if v := l.Remove(); v != 5 {
		t.Errorf("expected last removed value %v, got %v", 1, v)
	}
	chklist(t, `case 1`, l, []interface{}{1, 2, 3, 4})
	if v := l.Remove(l.First(), l.Last()); v != 4 {
		t.Errorf("expected last removed value %v, got %v", 4, v)
	}
	chklist(t, `case 2`, l, []interface{}{2, 3})
}

func TestList_RemoveFromAnotherList(t *testing.T) {
	l, ll := NewList(1, 2), NewList(3, 4)
	if v := l.Remove(ll.First()); v != nil {
		t.Errorf("expected value %v, got %d", nil, v)
	}
	chklist(t, `l`, l, []interface{}{1, 2})
	chklist(t, `ll`, ll, []interface{}{3, 4})
}

func chkempty(t *testing.T, prefix string, list *List) {
	if chklen(t, prefix, list, 0) {
		if len(prefix) > 0 {
			prefix += `: `
		}
		if list.First() != nil {
			t.Errorf("expected first node <nil>, got %v", list.First())
		}
		if list.Last() != nil {
			t.Errorf("expected last node <nil>, got %v", list.Last())
		}
		if root := &list.root; root != list.root.next || root != list.root.prev {
			t.Errorf("incorrect root %p: %+v", list, root)
		}
	}
}

func chklen(t *testing.T, prefix string, list *List, length int) (ok bool) {
	if ok = list.Len() == length; !ok {
		if len(prefix) > 0 {
			prefix += `: `
		}
		t.Errorf("%sexpected length %d, got %d", prefix, length, list.Len())
	}
	return
}

func chklist(t *testing.T, prefix string, list *List, exp []interface{}) {
	if !chklen(t, prefix, list, len(exp)) {
		return
	}
	var (
		node *Node
		ok   bool
		v    interface{}
	)
	if len(prefix) > 0 {
		prefix += `: `
	}
	for i, n := 0, list.First(); i < len(exp); i, n = i+1, n.Next() {
		if node, ok = exp[i].(*Node); ok {
			v = node.Value
		} else {
			v = exp[i]
		}
		if list != n.List() {
			t.Errorf("%svalue %v is in the wrong list", prefix, v)
		}
		if n.Value != v {
			t.Errorf("%sexpected value %v, got %v", prefix, v, n.Value)
		}
		if p := n.Prev(); p != nil {
			if node, ok = exp[i-1].(*Node); ok {
				v = node.Value
			} else {
				v = exp[i-1]
			}
			if p.Value != v {
				t.Errorf("%sexpected previous value %v, got %v", prefix, v, p.Value)
			}
		}
	}
}

func dumplist(l *List, depth int) {
	fmt.Printf("%p: %+v\n\n", l, l)
	for i, n := 0, l.First(); n != nil || i < depth; i, n = i+1, n.Next() {
		fmt.Printf("%p: %+v\n", n, n)
	}
	println()
}
