package queue

type Queue struct {
	data []interface{}
}

func (q * Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q * Queue) Pop() {
	q.data = q.data[1:]
}

func (q * Queue) Push(value interface{}) {
	q.data = append(q.data, value)
}

func (q * Queue) Front() interface{} {
	return q.data[0]
}

