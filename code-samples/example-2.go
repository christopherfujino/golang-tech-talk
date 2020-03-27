package main

import "fmt"

type LinkedNode struct {
  Value int
  Next *LinkedNode
}

func traverseList(node *LinkedNode) {
  for node.Next != nil {
    fmt.Println(node.Value)
    node = node.Next
  }
}

func main() {
  var startNode, currentNode *LinkedNode;
  startNode = new(LinkedNode)
  currentNode = startNode

  for i := 10; i >= 0; i-- {
    currentNode.Value = i
    currentNode.Next = new(LinkedNode)
    currentNode = currentNode.Next
  }

  traverseList(startNode)
}
