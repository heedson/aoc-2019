package day6

import (
	"fmt"
	"os"
	"strings"

	"github.com/heedson/aoc-2019/util"
)

// Day is the implementation of day 6.
type Day struct {
	orbits        []string
	totalOrbits   int
	santaDistance int
}

// New returns a new instantiation of a Day.
func New() *Day {
	reader, err := os.Open("./day6/day6.txt")
	if err != nil {
		panic(err)
	}
	orbits, err := util.StringsFromReader(reader, "\n")
	if err != nil {
		panic(err)
	}
	return &Day{orbits: orbits}
}

// P1 runs the day's Part 1 for the day. It prints the challenge output.
func (d *Day) P1() error {
	root := buildTree(d.orbits)
	d.totalOrbits = walk(root)
	fmt.Println("Total orbits:", d.totalOrbits)
	return nil
}

// P2 runs the day's Part 2 for the day. It prints the challenge output.
func (d *Day) P2() error {
	root := buildTree(d.orbits)
	d.santaDistance = dist(root, "YOU", "SAN")
	fmt.Println("Distance to Santa:", d.santaDistance)
	return nil
}

func buildTree(orbits []string) *node {
	var root *node
	nodes := make(map[string]*node)
	for _, o := range orbits {
		bodyNames := strings.Split(o, ")")
		parentBody, ok := nodes[bodyNames[0]]
		if !ok {
			parentBody = &node{
				name: bodyNames[0],
			}
			nodes[bodyNames[0]] = parentBody
		}
		childBody, ok := nodes[bodyNames[1]]
		if !ok {
			childBody = &node{
				name: bodyNames[1],
			}
			nodes[bodyNames[1]] = childBody
		}
		root = parentBody
		parentBody.children = append(parentBody.children, childBody)
		childBody.parent = parentBody
	}
	for {
		if root.parent == nil {
			break
		}
		root = root.parent
	}
	return root
}

type node struct {
	name     string
	parent   *node
	children []*node
}

func walk(n *node) int {
	if n == nil {
		return 0
	}
	var count int
	p := n.parent
	for p != nil {
		p = p.parent
		count++
	}
	for _, c := range n.children {
		count += walk(c)
	}
	return count
}

func dist(n *node, start, target string) int {
	if n == nil || start == target {
		return 0
	}
	startNode := find(n, start)
	targetNode := find(n, target)

	var startParentNodes []*node
	for startNode.parent != nil {
		startParentNodes = append(startParentNodes, startNode.parent)
		startNode = startNode.parent
	}
	var targetParentNodes []*node
	for targetNode.parent != nil {
		targetParentNodes = append(targetParentNodes, targetNode.parent)
		targetNode = targetNode.parent
	}

	for i, s := range startParentNodes {
		for j, t := range targetParentNodes {
			if s.name == t.name {
				return i + j
			}
		}
	}

	return 0
}

func find(n *node, target string) *node {
	if n == nil {
		return nil
	}
	if n.name == target {
		return n
	}
	var targetNode *node
	for _, c := range n.children {
		nn := find(c, target)
		if nn != nil && nn.name == target {
			targetNode = nn
			break
		}
	}
	return targetNode
}
