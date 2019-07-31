package graph

import ("fmt"
        "endgame/container")

type Node struct {
     Id int
     Content container.Type
     Neighbours []*Node
}

type Edge struct {
}

type Graph []Node

func BreadthFirstTraversal(id int, graph Graph) {
     if graph == nil {
     	return
     }

     visited := make([]bool, len(graph))
     var queue container.SliceQueue

     queue.Push(index)

     for !queue.Empty() {
     	 node_id := queue.Pop().(int)
	 fmt.Println(node_id)

	 visited[node_id] = true

	 neighbours := graph[node_id].Neighbours
	 for _, neighbour := range neighbours {
	     if !visited[neighbour.Id] {
	     	queue.Push(neighbour.Id)
	     }
	 }
     }
}

func DepthFirstTraversal(index int, graph Graph) {
     if graph == nil {
     	return
     }

     visited := make([]bool, len(graph))
     var stack container.SliceStack

     stack.Push(index)

     for !stack.Empty() {
     	 node_id := stack.Pop().(int)
	 fmt.Println(node_id)

	 visited[node_id] = true

	 neighbours := graph[node_id].Neighbours
	 for _, neighbour := range neighbours {
	     if !visited[neighbour.Id] {
	     	stack.Push(neighbour.Id)
	     }
	 }
     }
}

func RunBFS() {
     var g Graph = Graph{
     	 {0, nil},
	 {1, nil},
	 {2, nil},
	 {3, nil},
	 {4, nil},
     }

     g[0].Neighbours = []*Node{&g[1], &g[2]}
     g[1].Neighbours = []*Node{&g[3], &g[4]}

     BreadthFirstTraversal(0, g)   
}

func RunDFS() {
     var g Graph = Graph{
     	 {0, nil},
	 {1, nil},
	 {2, nil},
	 {3, nil},
	 {4, nil},
     }

     g[0].Neighbours = []*Node{&g[1], &g[2]}
     g[1].Neighbours = []*Node{&g[3], &g[4]}

     DepthFirstTraversal(0, g)   
}
