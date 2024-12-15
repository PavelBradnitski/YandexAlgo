package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Node struct {
	Value    string
	Children []*Node
}

func BuildTree(pairs [][2]string) *Node {
	nodes := make(map[string]*Node)

	for _, pair := range pairs {
		child, parent := pair[0], pair[1]
		if _, exists := nodes[child]; !exists {
			nodes[child] = &Node{Value: child}
		}
		if _, exists := nodes[parent]; !exists {
			nodes[parent] = &Node{Value: parent}
		}
		nodes[parent].Children = append(nodes[parent].Children, nodes[child])
	}

	hasParent := make(map[string]bool)
	for _, pair := range pairs {
		hasParent[pair[0]] = true
	}

	var root *Node
	for node := range nodes {
		if !hasParent[node] {
			root = nodes[node]
			break
		}
	}

	return root
}

func PrintTreeinOrderWithChildCount(node *Node, level int, mapNodeLevel map[string]int) {
	if node == nil {
		return
	}
	mapNodeLevel[node.Value] = len(node.Children)
	for _, child := range node.Children {
		PrintTreeinOrderWithChildCount(child, level+1, mapNodeLevel)
		mapNodeLevel[node.Value] += mapNodeLevel[child.Value]
	}
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var length, num int
	fmt.Fscan(in, &length, &num)
	pairs := make([][2]string, length-1)
	for i := 0; i < length-1; i++ {
		fmt.Fscan(in, &pairs[i][0], &pairs[i][1])
	}
	root := BuildTree(pairs)
	mapNodeLevel := make(map[string]int, length-1)
	PrintTreeinOrderWithChildCount(root, 0, mapNodeLevel)
	outNodeNames := make([]string, 0, len(mapNodeLevel))
	for key := range mapNodeLevel {
		outNodeNames = append(outNodeNames, key)
	}
	sort.Strings(outNodeNames)
	for _, v := range outNodeNames {
		fmt.Fprintf(out, "%s %d\n", v, mapNodeLevel[v])
	}
}

func PrintTreeinOrder(node *Node, level int) ([]string, map[string]int) {
	if node == nil {
		return nil, nil
	}
	var outNodeNames []string
	mapNodeLevel := make(map[string]int)
	outNodeNames = append(outNodeNames, node.Value)
	mapNodeLevel[node.Value] = level
	for _, child := range node.Children {
		o, m := PrintTreeinOrder(child, level+1)
		outNodeNames = append(outNodeNames, o...)
		for k, v := range m {
			mapNodeLevel[k] = v
		}
	}
	return outNodeNames, mapNodeLevel
}

//	func PrintTreeinOrderWithChildCount(node *Node, level int) ([]string, map[string]int) {
//		if node == nil {
//			return nil, nil
//		}
//		var outNodeNames []string
//		mapNodeLevel := make(map[string]int)
//		outNodeNames = append(outNodeNames, node.Value)
//		mapNodeLevel[node.Value] = len(node.Children)
//		for _, child := range node.Children {
//			o, m := PrintTreeinOrderWithChildCount(child, level+1)
//			outNodeNames = append(outNodeNames, o...)
//			for k, v := range m {
//				mapNodeLevel[k] = v
//			}
//			mapNodeLevel[node.Value] += mapNodeLevel[child.Value]
//		}
//		return outNodeNames, mapNodeLevel
//	}
