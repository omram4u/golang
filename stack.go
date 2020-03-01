package testgitsrc

type Stack struct {
	ll *Linklist
}

func (s *Stack) Push(val int) {
	s.ll.AddNode(val)
}

func Newlist() *Stack {
	return &Stack{ll: &Linklist{}}
}

func (s *Stack) PrintStack() {
	s.ll.PrintNode()
}
func (s *Stack) Empty() bool {
	return s.ll.root == nil
}

func (s *Stack) Last() int {
	return s.ll.tail.val
}
func (s *Stack) Pop() int {
	l := s.Last()
	if s.ll.root != s.ll.tail {
		s.ll.RemoveNode(s.ll.tail)
		return l
	}
	s.ll.root = nil
	return l
}
