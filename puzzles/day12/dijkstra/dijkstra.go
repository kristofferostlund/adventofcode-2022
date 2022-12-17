package dijkstra

import (
	"fmt"
	"math"
	"sync"
)

const initialCost = math.MaxInt64

type Node[T comparable] struct {
	t       T
	cost    int
	through *Node[T]
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("Node{cost: %d, t: %v, through: %v}", n.cost, n.t, n.through)
}

func (n *Node[T]) reset() {
	n.cost = initialCost
	n.through = nil
}

func (n *Node[T]) Value() T {
	return n.t
}

type Edge[T comparable] struct {
	node   *Node[T]
	weight int
}

type Graph[T comparable] struct {
	Nodes     []*Node[T]
	nodeIndex map[T]int
	Edges     map[T][]*Edge[T]

	mutex sync.RWMutex
}

func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{
		Nodes:     make([]*Node[T], 0),
		Edges:     make(map[T][]*Edge[T]),
		nodeIndex: make(map[T]int),
		mutex:     sync.RWMutex{},
	}
}

func NewNode[T comparable](val T) *Node[T] {
	return &Node[T]{val, initialCost, nil}
}

func (g *Graph[T]) GetNode(val T) *Node[T] {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	idx, ok := g.nodeIndex[val]
	if !ok {
		return nil
	}
	return g.Nodes[idx]
}

func (g *Graph[T]) AddNode(n *Node[T]) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	i := len(g.Nodes)
	g.nodeIndex[n.t] = i
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph[T]) AddEdge(from, to *Node[T], weight int) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.Edges[from.t] = append(g.Edges[from.t], &Edge[T]{to, weight})
}

func (g *Graph[T]) ShortestPath(start, dest T) (int, []T, bool) {
	g.dijkstra(start)

	idx, ok := g.nodeIndex[dest]
	if !ok {
		return 0, nil, false
	}

	node := g.Nodes[idx]
	path := make([]T, 0)

	for n := node; n.through != nil; n = n.through {
		path = append(path, n.t)
	}

	return node.cost, path, true
}

type SetHeap[T comparable] struct {
	elemLookup map[*Node[T]]struct{}
	elements   []*Node[T]
	mutex      sync.RWMutex
}

func (h *SetHeap[T]) Size() int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return len(h.elements)
}

func (h *SetHeap[T]) Push(element *Node[T]) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	if _, exists := h.elemLookup[element]; exists {
		return
	}
	h.elemLookup[element] = struct{}{}

	h.elements = append(h.elements, element)

	for i := len(h.elements) - 1; h.elements[i].cost < h.elements[parent(i)].cost; i = parent(i) {
		h.swap(i, parent(i))
	}
}

func (h *SetHeap[T]) Pop() *Node[T] {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	element := h.elements[0]

	delete(h.elemLookup, element)

	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	h.rearrange(0)

	return element
}

func (g *Graph[T]) dijkstra(start T) {
	// Reset the nodes so we can calculate the fastest path several times.
	for _, n := range g.Nodes {
		n.reset()
	}

	visited := make(map[*Node[T]]struct{}, len(g.Nodes))
	heap := &SetHeap[T]{elemLookup: make(map[*Node[T]]struct{})}

	startNode := g.GetNode(start)
	startNode.cost = 0

	heap.Push(startNode)

	for heap.Size() > 0 {
		current := heap.Pop()
		visited[current] = struct{}{}
		edges := g.Edges[current.t]

		for _, edge := range edges {
			if _, isVisited := visited[edge.node]; !isVisited {
				heap.Push(edge.node)

				if current.cost+edge.weight < edge.node.cost {
					edge.node.cost = current.cost + edge.weight
					edge.node.through = current
				}
			}
		}
	}
}

func (h *SetHeap[T]) rearrange(i int) {
	smallest := i
	left, right, size := leftChild(i), rightChild(i), len(h.elements)
	if left < size && h.elements[left].cost < h.elements[smallest].cost {
		smallest = left
	}
	if right < size && h.elements[right].cost < h.elements[smallest].cost {
		smallest = right
	}
	if smallest != i {
		h.swap(i, smallest)
		h.rearrange(smallest)
	}
}

func (h *SetHeap[T]) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}
