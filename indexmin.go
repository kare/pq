package pq // import "kkn.fi/pq"

// IndexMin struct represents an indexed priority queue of int keys. It
// supports the usual Insert and DelMin and DecreaseKey functions. In order to
// let the client refer to keys on the priority queue, an integer between 0 and
// max-1 is associated with each key the client uses this integer to specify
// which key to delete or change. It also has a function for testing if the
// priority queue is empty.
//
// This implementation uses a binary heap along with an slice to associate keys
// with floats in the given range. The Insert, DelMin and DecreaseKey
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
	keys []float32 // priorities
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

func (pq IndexMin) greater(i, j int) bool {
	return pq.keys[pq.pq[i]] > pq.keys[pq.pq[j]]
}

func (pq *IndexMin) exch(i, j int) {
	pq.pq[j], pq.pq[i] = pq.pq[i], pq.pq[j]
	pq.qp[pq.pq[i]] = i
	pq.qp[pq.pq[j]] = j
}

func (pq *IndexMin) swim(k int) {
	for k > 1 && pq.greater(k/2, k) {
		pq.exch(k, k/2)
		k = k / 2
	}
}

func (pq *IndexMin) sink(k int) {
	for 2*k <= pq.len {
		j := 2 * k
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
// If key is out of bounds or exists in the queue function will simply return.
func (pq *IndexMin) Insert(key int, priority float32) {
	if key < 0 || key >= pq.max {
		return
	}
	if pq.qp[key] != -1 {
		return
	}
	pq.len++
	pq.qp[key] = pq.len
	pq.pq[pq.len] = key
	pq.keys[key] = priority
	pq.swim(pq.len)
}

// DelMin deletes and returns smallest key.
// If queue is empty the function returns max.
func (pq *IndexMin) DelMin() int {
	if pq.len == 0 {
		return pq.max
	}
	min := pq.pq[1]
	pq.exch(1, pq.len)
	pq.len--
	pq.sink(1)
	pq.qp[min] = -1
	pq.pq[pq.len+1] = -1
	return min
}

// DecreaseKey decreases keys priority.
// If key is out of bounds or doesn't exist in the queue function will simply return.
func (pq *IndexMin) DecreaseKey(key int, priority float32) {
	if key < 0 || key >= pq.max {
		return
	}
	if pq.qp[key] == -1 {
		return
	}
	pq.keys[key] = priority
	pq.swim(pq.qp[key])
}

// Contains returns true if queue contains key and false otherwise.
func (pq IndexMin) Contains(key int) bool {
	if key < 0 || key >= pq.max {
		panic("index out of bounds")
	}
	return pq.qp[key] != -1
}

// Len returns the size of the queue.
func (pq IndexMin) Len() int {
	return pq.len
}

// IsEmpty returns true if queue is empty and false otherwise.
func (pq IndexMin) IsEmpty() bool {
	return pq.len == 0
}
