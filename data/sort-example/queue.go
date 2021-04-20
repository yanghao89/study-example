package sort_example

type  QueueExample struct {
	Head *DoubleNode `json:"head"`
	Tail *DoubleNode `json:"tail"`
	Count int `json:"count"`
}


func New() *QueueExample  {
	return &QueueExample{
		Head: nil,
		Tail: nil,
	}
}

func (q *QueueExample) AddFromHead(num int)  {
	cur := new(DoubleNode)
	cur.Value = num
	if q.Head == nil{
		q.Head = cur
		q.Tail = cur
	}else {
		cur.Next = q.Head
		q.Head.Last = cur
		q.Head = cur
	}
	q.Count++
}

func (q *QueueExample) AddFromTail(num int)  {
	cur := new(DoubleNode)
	cur.Value = num
	if q.Head == nil{
		q.Head = cur
		q.Tail = cur
	}else {
		q.Tail.Next = cur
		cur.Last = q.Tail
		q.Tail = cur
	}
	q.Count++
}

//栈 操作 先进后出
func (q *QueueExample) PopFromHead() int {
	if q.Head == nil{
		return 0
	}
	cur := q.Head
	if q.Head == q.Tail{
		q.Head = nil
		q.Tail = nil
	}else {
		q.Head = q.Head.Next
		q.Head.Last = nil
		cur.Next = nil
	}
	q.Count --
	return cur.Value
}



func (q *QueueExample) PopFromBottom() int  {
	if q.Head == nil{
		return 0
	}
	cur := q.Tail
	if q.Head == q.Tail{
		q.Head = nil
		q.Tail = nil
	}else {
		q.Tail = cur.Last
		q.Tail.Next = nil
		cur.Last = nil
	}
	return cur.Value
}

type MyStack struct {
	Stack *QueueExample
}

func NewStack()*MyStack  {
	return &MyStack{
		Stack:New(),
	}
}

func (m *MyStack) Push(v int)  {
	m.Stack.AddFromHead(v)
}

func (m *MyStack) Pop() int {
	return m.Stack.PopFromHead()
}

func (m *MyStack) Count()int  {
	return  m.Stack.Count
}

type  MyQueue struct {
	Queue *QueueExample
}
func NewQueue() *MyQueue {
	return &MyQueue{
		Queue: New(),
	}
}

func (q *MyQueue) Push(v int)  {
	q.Queue.AddFromHead(v)
}

func (q *MyQueue) Poll()int  {
	return q.Queue.PopFromBottom()
}

func (q *MyQueue) Count() int  {
	return q.Queue.Count
}