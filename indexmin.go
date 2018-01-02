package pq // import "kkn.fi/pq"

// IndexMin struct represents an indexed priority queue of float32 keys. It
// supports the usual Insert, DelMin, DecreaseKey and IncreaseKey functions.
// In order to let the client refer to keys on the priority queue, an integer
// between 0 and max-1 is associated with each key the client uses this integer
// to specify which key to delete or change. It also has a function for testing
// if the priority queue is empty.
//
// This implementation uses a binary heap along with an slice to associate keys
// with floats in the given range. The Insert, DelMin, DecreaseKey and IncreaseKey
// operations take logarithmic time. The IsEmpty and Len operations take
// constant time. Construction takes time proportional to the specified
// capacity.
//
// For additional documentation, see Section 2.4 of Algorithms, 4th Edition by
// Robert Sedgewick and Kevin Wayne.
type IndexMin struct {
	max  int
	len  int
	pq   []int
	qp   []int
	keys []float32 // keys[i] = priority of i
}

// NewIndexMin creates a new and empty minimium priority queue with given
// maximum value for keys. The maximum key index value is max - 1.
func NewIndexMin(max int) *IndexMin {
	qp := make([]int, max+1)
	for i := 0; i <= max; i++ {
		qp[i] = -1
	}
	return &IndexMin{
		max:  max,
		len:  0,
		pq:   make([]int, max+1),
		qp:   qp,
		keys: make([]float32, max+1),
	}
}

// Clear resets and initializes the priority queue to empty state.
func (pq *IndexMin) Clear() {
	qp := make([]int, pq.max+1)
	for i := 0; i <= pq.max; i++ {
		qp[i] = -1
	}
	pq.len = 0
	pq.pq = make([]int, pq.max+1)
	pq.qp = qp
	pq.keys = make([]float32, pq.max+1)
}

func (pq IndexMin) greater(i, j int) bool {
	return pq.keys[pq.pq[i]] > pq.keys[pq.pq[j]]
}

func (pq *IndexMin) exch(i, j int) {
	pq.pq[j], pq.pq[i] = pq.pq[i], pq.pq[j]
	pq.qp[pq.pq[i]] = i
	pq.qp[pq.pq[j]] = j
}

func (pq *IndexMin) swim(k int) {
	for k > 1 && pq.greater(k>>1, k) {
		pq.exch(k, k>>1)
		k = k >> 1
	}
}

func (pq *IndexMin) sink(k int) {
	for k<<1 <= pq.len {
		j := k << 1
		if j < pq.len && pq.greater(j, j+1) {
			j++
		}
		if !pq.greater(k, j) {
			break
		}
		pq.exch(k, j)
		k = j
	}
}

// Insert a key into the priority queue.
func (pq *IndexMin) Insert(i int, key float32) {
	pq.len++
	pq.qp[i] = pq.len
	pq.pq[pq.len] = i
	pq.keys[i] = key
	pq.swim(pq.len)
}

// DelMin deletes and returns smallest key.
func (pq *IndexMin) DelMin() int {
	min := pq.pq[1]
	pq.exch(1, pq.len)
	pq.len--
	pq.sink(1)
	pq.qp[min] = -1
	pq.pq[pq.len+1] = -1
	return min
}

// DecreaseKey decreases indexes priority.
func (pq *IndexMin) DecreaseKey(i int, key float32) {
	if pq.keys[i] <= key {
		return
	}
	pq.keys[i] = key
	pq.swim(pq.qp[i])
}

// IncreaseKey increases key associated with index to the specified priority.
func (pq *IndexMin) IncreaseKey(i int, key float32) {
	if pq.keys[i] >= key {
		return
	}
	pq.keys[i] = key
	pq.sink(pq.qp[i])
}

// Contains returns true if queue contains index and false otherwise.
func (pq IndexMin) Contains(i int) bool {
	if i < 0 || i >= pq.max {
		return false
	}
	return pq.qp[i] != -1
}

// Len returns the size of the queue.
func (pq IndexMin) Len() int {
	return pq.len
}

// IsEmpty returns true if queue is empty and false otherwise.
func (pq IndexMin) IsEmpty() bool {
	return pq.len == 0
}
