package graph

import (
  "fmt"
)

type ID interface {
        String() string
}

type StringID string

func (s StringID) String() string {
        return string(s)
}

type Node interface {
          ID() ID
          String() string
}

type node struct {
          id string
}

var nodeCount uint64

func NewNode(id string) Node {
  return &node{
    id: id,
  }
}

func (n *node) ID() ID {
  return StringID(n.id)
}

func (n *node) String() string {
  return n.id
}

type Edge interface {
  Source() Node
  Target() Node
  Weight() float64
  String() string
}

type edge struct {
    source Node
    target Node
    weight float64
}

func NewEdge(source, target Node, weight float64) Edge {
  return &edge{
    source: source,
    target: target,
    weight: weight,
  }
}

func (e *edge) Source() Node {
  return e.source
}

func (e *edge) Target() Node {
  return e.target
}

func (e *edge) Weight() float64 {
  return e.weight
}

func (e *edge) String() string {
	return fmt.Sprintf("%s -- %.3f -→ %s\n", e.source, e.weight, e.target) //string representatinh of edge
}





