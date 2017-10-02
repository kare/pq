package pq

import (
	"testing"
)

func TestDelMin(t *testing.T) {
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

func TestDelMinInsertBetweenDelete(t *testing.T) {
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

func TestContains(t *testing.T) {
	values := []float32{0.7, 0.123, 0.3, 0.453}
	pq := NewIndexMin(20)
	for i, v := range values {
		pq.Insert(i, v)
	}
	if pq.Contains(10) {
		t.Error("pq doesn't contain 10, but reports that it does")
	}
	if !pq.Contains(1) {
		t.Error("pq contains 1, but reports that it doesn't")
	}
	if pq.Contains(100) {
		t.Error("pq doesn't contain 100, but reports that it does")
	}
}

func TestDecreaseKey(t *testing.T) {
	values := []float32{0.7, 0.123, 0.3, 0.453}
	pq := NewIndexMin(20)
	for i, v := range values {
		pq.Insert(i, v)
	}
	pq.DecreaseKey(2, 2)
	pq.DecreaseKey(2, 0.1)
	d := pq.DelMin()
	if d != 2 {
		t.Errorf("expected key 2, but got %d", d)
	}
}

func TestIncreaseKey(t *testing.T) {
	values := []float32{1, 0, 2}
	pq := NewIndexMin(20)
	for i, v := range values {
		pq.Insert(i, v)
	}
	pq.IncreaseKey(1, 7)
	pq.IncreaseKey(0, 0)
	d := pq.DelMin()
	if d != 0 {
		t.Errorf("expected key 0, but got %d", d)
	}
	d = pq.DelMin()
	if d != 2 {
		t.Errorf("expected key 2, but got %d", d)
	}
	d = pq.DelMin()
	if d != 1 {
		t.Errorf("expected key 1, but got %d", d)
	}
}

var r int

func BenchmarkInsertAndDelMin(b *testing.B) {
	testData := []struct {
		i   int
		key float32
	}{
		{1, 1.2},
		{2, 2.2},
		{3, 2.1},
		{4, 1.1},
		{5, 4.9},
		{6, 2.7},
		{7, 3.3},
		{8, 0.1},
		{9, 9.3},
		{10, 6.3},
	}
	b.ResetTimer()
	p := NewIndexMin(len(testData))
	for n := 0; n < b.N; n++ {
		for _, testCase := range testData {
			p.Insert(testCase.i, testCase.key)
		}
		for range testData {
			r = p.DelMin()
		}
	}
}

func BenchmarkDecreaseKey(b *testing.B) {
	testData := []struct {
		i   int
		key float32
	}{
		{1, 1.2},
		{2, 2.2},
		{3, 2.1},
		{4, 1.1},
		{5, 4.9},
		{6, 2.7},
		{7, 3.3},
		{8, 0.1},
		{9, 9.3},
		{10, 6.3},
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		p := NewIndexMin(len(testData))
		for _, testCase := range testData {
			p.Insert(testCase.i, testCase.key)
		}
		for _, testCase := range testData {
			p.DecreaseKey(testCase.i, testCase.key-0.1)
		}
	}
}
