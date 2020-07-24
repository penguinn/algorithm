package graph

// 广度优先搜索算法
func BFS(graph *Graph, startIndex, targetIndex int) []int {
	length := len(graph.Nodes)
	if startIndex == targetIndex || startIndex > length-1 || targetIndex > length-1 {
		return nil
	}
	queue := make([]int, 0, length)
	visited := make([]bool, length)
	prevNodes := make([]int, length)
	for i, _ := range prevNodes {
		prevNodes[i] = -1
	}

	queue = append(queue, startIndex)
	var find bool
	for len(queue) != 0 {
		nodeIndex := queue[0]
		queue = queue[1:]
		visited[nodeIndex] = true

		subNodeIndexs := graph.Nodes[nodeIndex]
		for _, subNodeIndex := range subNodeIndexs {
			if visited[subNodeIndex] {
				continue
			}
			prevNodes[subNodeIndex] = nodeIndex
			if subNodeIndex == targetIndex {
				find = true
				break
			}
			queue = append(queue, subNodeIndex)
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

func PrintPath(result []int, index int, prevNodes []int) []int {
	prevIndex := prevNodes[index]
	if prevIndex != -1 {
		result = PrintPath(result, prevIndex, prevNodes)
		result = append(result, index)
		return result
	} else {
		result = append(result, index)
		return result
	}
}
