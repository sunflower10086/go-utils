package queue

type base interface {
	push()
	pop()
	IsEmpty()
}

type Queue struct {
	hh  int
	tt  int
	arr []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		hh:  0,
		tt:  -1,
		arr: make([]interface{}, 0),
	}
}

func (q *Queue) push(x ...interface{}) {
	for _, item := range x {
		q.tt++
		q.arr = append(q.arr, item)
	}
}

func (q *Queue) pop() {
	q.hh++
}

func (q *Queue) IsEmpty() bool {
	return q.hh > q.tt
}

// front 返回队头
func (q *Queue) front() interface{} {
	return q.arr[q.hh]
}

func (q *Queue) size() int {
	return len(q.arr)
}
