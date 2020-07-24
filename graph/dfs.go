package graph

// 深度优先搜索算法（没有使用递归，使用的栈）
func DFS(graph *Graph, startIndex, targetIndex int) []int {
	length := len(graph.Nodes)
	if startIndex == targetIndex || startIndex > length-1 || targetIndex > length-1 {
		return nil
	}
	stack := make([]int, 0, length)
	visited := make([]bool, length)
	prevNodes := make([]int, length)
	for i, _ := range prevNodes {
		prevNodes[i] = -1
	}

	stack = append(stack, startIndex)
	var find bool
	for len(stack) != 0 {
		nodeIndex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited[nodeIndex] = true

		subNodeIndexs := graph.Nodes[nodeIndex]
		l := len(subNodeIndexs)
		for i, _ := range subNodeIndexs {
			if visited[subNodeIndexs[l-1-i]] {
				continue
			}
			prevNodes[subNodeIndexs[l-1-i]] = nodeIndex
			if subNodeIndexs[l-1-i] == targetIndex {
				find = true
				break
			}
			stack = append(stack, subNodeIndexs[l-1-i])
		}
		if find == true {
			break
		}
	}
	var result []int
	if find {
		result = PrintPath(result, targetIndex, prevNodes)
		return result
	} else {
		return []int{}
	}
}
