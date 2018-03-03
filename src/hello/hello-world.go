package main

import (
       "fmt"
       "endgame"
       "endgame/container"
)

func main() {
     fmt.Println("Hello, World!", endgame.Foo(4, 5))

     //endgame.RunDFS()
     //endgame.QueueTest()
     var q container.SliceQueue
     q.Push("Hello, ")
     q.Push("World!")
}