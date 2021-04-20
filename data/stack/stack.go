package stack

import "container/list"

type Stack struct {
	list *list.List
}

func NewStack() *Stack  {
	return &Stack{
		list: list.New(),
	}
}

func (s *Stack) Push( value interface{})  {
	s.list.PushBack(value)
}

func (s *Stack) Pop() interface{}  {
 	e := s.list.Back()
	if e != nil{
		s.list.Remove(e)
		return e.Value
	}
	return  nil
}

func (s *Stack) Peak()interface{}  {
	e := s.list.Back()
	if e != nil{
		s.list.Remove(e)
		return e.Value
	}
	return  nil
}

func (s *Stack) Len()int  {
	return  s.list.Len()
}

func (s *Stack) IsEmpty() bool  {
	return  s.list.Len() == 0
}