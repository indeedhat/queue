package queue

import (
  "errors"
)

func Filo() *FiloQueue {
  return &FiloQueue{}
}

type FiloQueue struct {
  nodes   []*Node
  count   int
  maxSize int
}

func (q *FiloQueue) Push(node *Node) error {
  if q.Full() {
    return errors.New("Queue is full")
  }

  q.nodes = append(q.nodes[:q.count], node)
  q.count++

  return nil
}

func (q * FiloQueue) Pop() (*Node, error) {
  if q.Empty() {
    return nil, errors.New("Queue is empty")
  }

  q.count--
  return q.nodes[q.count], nil
}

func (q *FiloQueue) Empty() bool {
  return 0 >= q.count
}

func (q *FiloQueue) Full() bool {
  if 0 >= q.maxSize || q.count >= q.maxSize {
    return false
  }

  return true
}