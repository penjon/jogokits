package dfa

import (
	"strings"
)

type Tire struct {
	root        *node
	replaceHold string
}

func NewTire() *Tire {
	t := &Tire{
		root:        &node{},
		replaceHold: "*",
	}
	return t
}

type node struct {
	children map[uint8]*node
	isEnd    bool
}

func (t *Tire) ReplaceHold(char string) *Tire {
	t.replaceHold = char
	return t
}

func (t *Tire) Build(data []string) {
	root := &node{
		children: make(map[uint8]*node),
	}
	t.root = root
	for _, value := range data {
		parent := root
		v := strings.ToLower(value)
		for i := 0; i < len(v); i++ {
			parent = parse(parent, v, i)
		}
		parent.isEnd = true
	}
}

func (t *Tire) Add(value string) {
	parent := t.root
	v := strings.ToLower(value)
	for i := 0; i < len(v); i++ {
		parent = parse(parent, v, i)
	}
	parent.isEnd = true
}

func (t *Tire) Filter(word string) string {

	value := strings.ToLower(word)
	for i := 0; i < len(value); i++ {
		node := t.root
		hit := false
		hitIdx := i
		for s := i; s < len(value); s++ {
			node = findNode(node, value, s)
			if node == nil {
				break
			}
			if node.isEnd && (s+1 >= len(value) || value[s+1:s+2] == " ") {
				hit = true
				hitIdx = s + 1
				break
			}
		}
		if hit {
			sw := word[i:hitIdx]
			b := strings.Builder{}
			for ii := 0; ii < len(sw); ii++ {
				b.WriteString(t.replaceHold)
			}
			b.String()
			word = strings.Replace(word, sw, b.String(), -1)
			i = hitIdx

		}
	}
	return word
}

func findNode(parent *node, value string, index int) *node {
	char := value[index]
	children := parent.children
	if nil == children {
		return nil
	}
	child := children[char]
	return child
}

func parse(parent *node, value string, index int) *node {
	char := value[index]
	children := parent.children
	if nil == children {
		children = make(map[uint8]*node)
	}
	child, ok := children[char]
	if ok {
		//有节点
		return child
	}
	child = &node{}
	children[char] = child
	parent.children = children
	return child
}
