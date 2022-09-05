package tree

type NodeStack struct {
	items []*Node
}

func NewNodeStack() *NodeStack {
	return &NodeStack{
		items: make([]*Node, 0),
	}
}

func (s *NodeStack) WithNodes(nodes []*Node) *NodeStack {
	s.items = nodes
	return s
}

func (s *NodeStack) Push(n *Node) {
	s.items = append(s.items, n)
}

func (s *NodeStack) Pop() *Node {
	if s.IsEmpty() {
		return nil
	}

	n := s.items[s.Size()-1]
	s.items = s.items[:s.Size()-1]
	return n
}

func (s *NodeStack) Size() int {
	return len(s.items)
}

func (s *NodeStack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *NodeStack) Nodes() []*Node {
	sol := make([]*Node, s.Size())
	copy(s.items, sol)
	return sol
}

func (s *NodeStack) ToQueue() *NodeQueue {
	return NewNodeQueue().WithNodes(s.items)
}
