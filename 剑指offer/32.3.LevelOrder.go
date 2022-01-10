package main

import "container/list"

func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	queue := list.New()
	queue.PushBack(root)
	level := 0

	for queue.Len() > 0 {
		curLevelNodes := queue.Len()
		level++
		nodeVals := []int{}
		for i := 0; i < curLevelNodes; i++ {
			if level%2 == 1 {
				node := queue.Remove(queue.Front()).(*TreeNode)
				nodeVals = append(nodeVals, node.Val)
				if node.Left != nil {
					queue.PushBack(node.Left)
				}
				if node.Right != nil {
					queue.PushBack(node.Right)
				}
			} else {
				node := queue.Remove(queue.Back()).(*TreeNode)
				nodeVals = append(nodeVals, node.Val)
				if node.Right != nil {
					queue.PushFront(node.Right)
				}
				if node.Left != nil {
					queue.PushFront(node.Left)
				}
			}
		}
		res = append(res, nodeVals)
	}
	return res
}

func main() {

}
