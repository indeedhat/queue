package queue

import(
  "errors"
)

func Fifo() *FifoQueue {
  return &FifoQueue{}
}

type FifoQueue struct {
  nodes   []*Node
  count   int
  maxSize int
}

func (q *FifoQueue) Push(node *Node) error {
  if q.Full() {
    return errors.New("Queue is full")
  }

  q.nodes = append(q.nodes[:q.count], node)
  q.count++

  return nil
}

func (q *FifoQueue) Pop() (*Node, error) {
  if q.Empty() {
    return nil, errors.New("Queue is empty")
  }

  q.count--
  node := q.nodes[0]

  q.nodes = q.nodes[1:]

  return node, nil
}

func (q *FifoQueue) Empty() bool {
  return 0 >= q.count
}

func (q *FifoQueue) Full() bool {
  if 0 >= q.maxSize || q.count <= q.maxSize {
    return false
  }

  return true
}
