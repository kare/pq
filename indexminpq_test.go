package pq // import "kkn.fi/pq"

import "testing"

func TestIndexMinPQ(t *testing.T) {
	values := []float32{0.7, 0.123, 0.453, 0.23, 0.657, 0.120, 0.4246, 0.12, 0.9999, 0.123123}
	pq := NewIndexMin(50)
	for i, v := range values {
		pq.Insert(i, v)
	}
	if pq.Len() != len(values) {
		t.Errorf("expected length %d, but got %d", len(values), pq.Len())
	}
	n := 0
	expected := []float32{0.12, 0.120, 0.123, 0.123123, 0.23, 0.4246, 0.453, 0.657, 0.7, 0.9999}
	for !pq.IsEmpty() {
		i := pq.DelMin()
		if expected[n] != values[i] {
			t.Errorf("expected value %2.2f, but got %2.2f", expected[n], values[i])
		}
		n++
	}
}

func TestIndexMinPQBetween(t *testing.T) {
	values := []float32{0.7, 0.123, 0.3, 0.453}
	pq := NewIndexMin(20)
	for i, v := range values {
		pq.Insert(i, v)
	}
	i := pq.DelMin()
	if i != 1 {
		t.Errorf("expected 1, but got %v", i)
	}
	pq.Insert(5, 22)
	i = pq.DelMin()
	if i != 2 {
		t.Errorf("expected 2, but got %v", i)
	}
	i = pq.DelMin()
	if i != 3 {
		t.Errorf("expected 3, but got %v", i)
	}
}