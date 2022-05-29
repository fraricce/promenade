// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

type Vertex struct {
	id          int
	name        string
	adjacencies []*Vertex
}

type Graph struct {
	vertices []*Vertex
}

func (g *Graph) AddVertex(v *Vertex) {
	g.vertices = append(g.vertices, v)
}

func (g *Graph) AddEdge(vFrom, vTo *Vertex) {
	vFrom.adjacencies = append(vFrom.adjacencies, vTo)
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Print(v.name)
		for _, vv := range v.adjacencies {
			fmt.Printf("-> %s", vv.name)
		}
		fmt.Println()
	}
}

func (g *Graph) GetNode(v *Vertex) *Vertex {
	for _, n := range g.vertices {
		if n.id == v.id {
			return v
		}
	}
	return nil
}

func (g *Graph) GetNodeById(id int) *Vertex {
	for _, v := range g.vertices {
		if id == v.id {
			return v
		}
	}
	return nil
}

func (g *Graph) IsVisited(id int, visited []int) bool {
	for _, v := range visited {
		if id == v {
			return true
		}
	}
	return false
}

func (g *Graph) FindPath(start, end int) []*Vertex {
	path := []*Vertex{}
	var startNode *Vertex
	var visitedNodes []int
	canVisitNode := true

	for _, v := range g.vertices {
		if v.id == start {
			startNode = v
			break
		}
	}

	if startNode == nil {
		return []*Vertex{}
	}

	currNode := g.GetNode(startNode)
	if currNode == nil {
		canVisitNode = false
		return []*Vertex{}
	}

	path = append(path, currNode)
	visitedNodes = append(visitedNodes, currNode.id)

	for canVisitNode {
		// look adjancencies
		for _, n := range currNode.adjacencies {
			if n.id == end {
				path = append(path, g.GetNodeById(n.id))
				visitedNodes = append(visitedNodes, n.id)
				return path
			}
		}

		// select a node to visit
		for _, n := range currNode.adjacencies {
			if !g.IsVisited(n.id, visitedNodes) {
				path = append(path, g.GetNodeById(n.id))
				visitedNodes = append(visitedNodes, n.id)
				currNode = g.GetNodeById(n.id)
				break
			}
		}

		// test if we have nodes to explore

		canVisitNode = false

		for _, v := range currNode.adjacencies {
			if !g.IsVisited(v.id, visitedNodes) {
				canVisitNode = true
				break
			}
		}

		if !canVisitNode {
			currNode = path[(len(path)-1)-1]
			path = path[:len(path)-1]
			canVisitNode = true
			continue
		}
	}

	return path
}

func main() {
	g := &Graph{}
	v := &Vertex{id: 1, name: "Porte Dauphine"}
	v2 := &Vertex{id: 2, name: "Victor Hugo"}
	v3 := &Vertex{id: 3, name: "Charles De Gaulle Etoile"}
	v4 := &Vertex{id: 4, name: "George V"}
	v5 := &Vertex{id: 5, name: "Ternes"}
	v6 := &Vertex{id: 6, name: "Monceau"}
	v7 := &Vertex{id: 7, name: "Courcelles"}
	v8 := &Vertex{id: 8, name: "Villiers"}
	v9 := &Vertex{id: 9, name: "Pl.De Clichy"}
	v10 := &Vertex{id: 10, name: "Europe"}

	g.AddVertex(v)
	g.AddVertex(v2)
	g.AddVertex(v3)
	g.AddVertex(v4)
	g.AddVertex(v5)
	g.AddVertex(v6)
	g.AddVertex(v7)
	g.AddVertex(v8)
	g.AddVertex(v9)
	g.AddVertex(v10)

	g.AddEdge(v, v2)
	g.AddEdge(v2, v3)
	g.AddEdge(v3, v4)
	g.AddEdge(v3, v5)
	g.AddEdge(v5, v6)
	g.AddEdge(v6, v7)
	g.AddEdge(v7, v8)
	g.AddEdge(v8, v9)
	g.AddEdge(v8, v10)

	//g.Print()
	fmt.Println()
	path := g.FindPath(2, 9)
	for _, n := range path {
		fmt.Printf(" -> %s", n.name)
	}
}
