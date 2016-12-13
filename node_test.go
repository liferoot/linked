package linked

import "testing"

func TestNode_AttachItself(t *testing.T) {
	n := NewNode(1)
	if n != n.Attach(n) {
		t.Errorf(`incorrect attachment`)
	}
}

func TestNode_AttachInSeries(t *testing.T) {
	first := NewNode(1)
	i, m, n := 2, 5, first
	for ; i < m; i++ {
		n = n.Attach(NewNode(i))
	}
	for ; n != nil; n = n.Prev() {
		if i--; i != n.Value {
			t.Errorf("backwards: expected value %d, got %d", i, n.Value)
		}
	}
	for n = first; n != nil; n = n.Next() {
		if i != n.Value {
			t.Errorf("forwards: expected value %d, got %d", i, n.Value)
		}
		i++
	}
	if i != m {
		t.Errorf("expected iterations %d, got %d", m, i)
	}
}

func TestNode_AttachInside(t *testing.T) {
	first, i := NewNode(1), 0
	first.Attach(NewNode(3))
	first.Attach(NewNode(2))
	for n := first; n != nil; n = n.Next() {
		if i++; i != n.Value {
			t.Errorf("expected value %d, got %d", i, n.Value)
		}
	}
	if i != 3 {
		t.Errorf("expected iterations %d, got %d", 3, i)
	}
}

func TestNode_Detach(t *testing.T) {
	first := NewNode(1)
	touch := first.Attach(NewNode(2))
	last := touch.Attach(NewNode(3))
	for i, n := 0, first; n != nil; n = n.Next() {
		if i++; i != n.Value {
			t.Errorf("expected value %d, got %d", i, n.Value)
		}
	}
	if touch.Detach(); first != last.Prev() || first.Next() != last {
		t.Errorf("first and last nodes not properly linked")
	}
}
