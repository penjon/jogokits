package utils

import (
	"math/rand"
)

type WeightRandom struct {
	maxWeightValue int
	weight         []int
	nodes          []*weightRandomNode
}
type weightRandomNode struct {
	weight int
	low    int
	height int
}

func (i *WeightRandom) UpdateWeight(args []int) {
	i.weight = args
	i.nodes = make([]*weightRandomNode, len(args))

	for idx := 0; idx < len(args); idx++ {
		i.nodes[idx] = &weightRandomNode{
			weight: args[idx],
		}
	}
	i.InitNodes()
}

func (i *WeightRandom) InitNodes() {
	for idx, node := range i.nodes {
		if idx == 0 {
			node.low = 0
			node.height = node.weight
		} else {
			node.low = i.nodes[idx-1].height
			node.height = node.low + node.weight
		}
		i.maxWeightValue = node.height
	}
}

func (i *WeightRandom) RemoveNode(index int) {
	offset := 0
	for idx, node := range i.nodes {
		if idx != index {
			i.nodes[offset] = node
			offset++
		}
	}

	i.nodes = i.nodes[:offset]
}

func (i *WeightRandom) NextIndex() int {
	value := rand.Intn(i.maxWeightValue)
	for idx, node := range i.nodes {
		if value >= node.low && value < node.height {
			return idx
		}
	}
	return len(i.nodes) - 1
}
