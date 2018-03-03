package endgame

import "fmt"

type Node struct {
     Label int
     Neighbours []*Node
}

type Graph []Node

func DepthFirstTraversal(index int, graph Graph, visited []bool) {
     if graph == nil {
     	return
     }

     fmt.Println(index)

     neighbours := graph[index].Neighbours

     for _, n := range neighbours {
     	 if !visited[n.Label] {
	    visited[n.Label] = true
	    DepthFirstTraversal(n.Label, graph, visited)
	 }
     }
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

     DepthFirstTraversal(0, g, []bool{false, false, false, false, false})   
}
