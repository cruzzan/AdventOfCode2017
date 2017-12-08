package main

import (
	"os"
	"bufio"
	"regexp"
	"strconv"
	"fmt"
)

// MODEL FOR NODES
type node struct {
	name string
	weight int
	children []string
	parent string
	branchWeight int
}

func (n node) hasChildren() bool {
	return len(n.children) > 0
}

func (n node) hasParent() bool {
	return n.parent != ""
}

func (n *node) setParent(parentName string) {
	n.parent = parentName
}


// MAIN
func main() {
	nodes := readFromFile()
	nodes = createParentLink(nodes)

	fmt.Println("Task 1: ", task1(nodes), "Task 2: ", task2(nodes))
}

func checkError(err error)  {
	if err != nil {
		panic(err)
	}
}

// PARSE INPUT INTO NODE-TREE-LIST

func readFromFile() map[string]node {
	var nodes = make(map[string]node)
	file, err := os.Open("input.txt")
	checkError(err)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		name := findValueInString("([a-z]+)\\s\\(", line)[0][1]
		weight, err := strconv.Atoi(findValueInString("\\((\\d+)\\)", line)[0][1])
		var childrenList []string
		children := findValueInString("->\\s(.*)", line)

		if len(children) > 0 {
			for _, child := range findValueInString("(\\w+)", children[0][1]) {
				childrenList = append(childrenList, child[1])
			}
		}

		checkError(err)

		nodes[name] = node{name:name, weight:weight, children:childrenList}
	}

	return nodes
}

func findValueInString(pattern string, haystack string) [][]string {
	regex := regexp.MustCompile(pattern)
	return regex.FindAllStringSubmatch(haystack, -1)
}

func createParentLink(nodes map[string]node) map[string]node {
	for _, node := range nodes {
		if node.hasChildren() {
			for _, child := range node.children {
				childNode := nodes[child]
				childNode.setParent(node.name)

				nodes[child] = childNode
			}
		}
	}

	return nodes
}

// TASK 1

func task1(nodes map[string]node) string {
	nodesCopy := nodes
	firstNode := findFirstNode(nodesCopy)

	return findParent(nodesCopy, nodesCopy[firstNode].name)
}

// TASK 2
func task2(nodes map[string]node) int {
	nodesCopy := nodes
	for _, node := range nodesCopy {
		node.branchWeight = sumBranch(nodesCopy, node.name)
		nodesCopy[node.name] = node
	}

	rootNode := findParent(nodes, findFirstNode(nodes))

	return findBalancedLevelValue(nodes, rootNode, 0)
}

func findBalancedLevelValue(nodes map[string]node, levelBase string, balanced int) int {
	var siblings = nodes[levelBase].children
	var branches = make(map[int][]node)

	for _, sibling := range siblings {
		siblingNode := nodes[sibling]
		branches[siblingNode.branchWeight] = append(branches[siblingNode.branchWeight], siblingNode)
	}

	var correctWeight int
	var wrongWeight = 0
	var unbalancedNode string
	for branchWeight, matchedNodes := range branches {
		if len(matchedNodes) == 1 {
			wrongWeight = branchWeight
			unbalancedNode = matchedNodes[0].name
		} else {
			correctWeight = branchWeight
		}
	}

	if wrongWeight == 0 {
		return balanced
	} else {
		bal := nodes[unbalancedNode].weight+(correctWeight-wrongWeight)
		return findBalancedLevelValue(nodes, unbalancedNode, bal)
	}
}

func sumBranch(nodes map[string]node, nodeName string) int {
	node := nodes[nodeName]
	if node.hasChildren() {
		var total = 0
		for _, child := range node.children {
			total += sumBranch(nodes, child)
		}
		return total + node.weight
	}

	return node.weight
}

//MISC
func findFirstNode(nodes map[string]node) string {
	var firstNode string
	for node := range nodes {
		firstNode = node
		break
	}

	return firstNode
}

func findParent(nodes map[string]node, name string) string {
	if nodes[name].hasParent() {
		return findParent(nodes, nodes[name].parent)
	}

	return name
}
