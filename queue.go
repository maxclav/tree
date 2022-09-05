package tree

type NodeQueue struct {
	items []*Node
}

func NewNodeQueue() *NodeQueue {
	return &NodeQueue{
		items: make([]*Node, 0),
	}
}

func (q *NodeQueue) WithNodes(nodes []*Node) *NodeQueue {
	q.items = nodes
	return q
}

func (q *NodeQueue) Enqueue(n *Node) {
	q.items = append(q.items, n)
}

func (q *NodeQueue) Dequeue() *Node {
	if q.IsEmpty() {
		return nil
	}
	n := q.items[0]
	q.items = q.items[1:]
	return n
}

func (q *NodeQueue) Size() int {
	return len(q.items)
}

func (q *NodeQueue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *NodeQueue) Nodes() []*Node {
	sol := make([]*Node, q.Size())
	copy(q.items, sol)
	return sol
}

func (q *NodeQueue) ToStack() *NodeStack {
	return NewNodeStack().WithNodes(q.items)
}
