package graph

import()

const(
  NoPathError = "No path exists between nodes"
)

type PathError struct {
    Error string
    Target Node
    Source Node
}

func Dijkstra(g Graph, source, target ID) ([]ID, map[ID]float64, error) {
  return nil, nil, nil
}
