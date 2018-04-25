package queue

type Queue struct {
	capacity uint
	start    uint
	end      uint
	T        []interface{}
}

func NewQueue(capacity uint) *Queue {
	return &Queue{
		capacity: capacity,
		T:        make([]interface{}, capacity),
	}
}
func (q *Queue) Push(t interface{}) bool {
	if (q.end+1)%q.capacity == q.start {
		return false
	}
	q.T[q.end] = t
	q.end = (q.end + 1) % q.capacity
	return true
}

func (q *Queue) Pop() (interface{}, bool) {
	if q.end == q.start {
		return nil, false
	}
	temp := q.T[q.start]
	q.start = (q.start + 1) % q.capacity
	return temp, true
}
func (q *Queue) isEmpty() bool {
	return q.start == q.end
}
func (q *Queue) Size() uint {
	return (q.end - q.start + q.capacity) % q.capacity
}
