package Node

import (
	"fmt"
	"strconv"
	"strings"
)

type Neighbourhood struct {
	Top    *Node
	Bottom *Node
	Left   *Node
	Right  *Node
}

type Group struct {
	Nodes []*Node
}

func (group Group) String() string {
	stringValues := []string{}

	for _, node := range group.Nodes {
		stringValues = append(stringValues, fmt.Sprintf("%d", node.Value))
	}

	return strings.Join(stringValues, ",")
}

type Node struct {
	Value          int
	Neighbourhood  Neighbourhood
	Group          *Group
	PossibleValues []int
}

func (node Node) String() string {

	top := fmt.Sprintf("  %v  \n", node.Neighbourhood.Top != nil)
	middle := fmt.Sprintf("%v %d %v\n", node.Neighbourhood.Left != nil, node.Value, node.Neighbourhood.Right != nil)
	bottom := fmt.Sprintf("  %v  \n", node.Neighbourhood.Bottom != nil)

	return top + middle + bottom
}

func CreateSudokuFromString(input string) *Node {
	stringValues := strings.Split(input, ";")

	head := &Node{Neighbourhood: Neighbourhood{}}

	var current *Node = head
	var previous *Node = nil
	var top *Node = nil

	for index, stringValue := range stringValues {

		value, err := strconv.Atoi(stringValue)

		if err != nil {
			panic("Invalid input")
		}

		current.Value = value

		if previous != nil {
			current.Neighbourhood.Left = previous
			previous.Neighbourhood.Right = current
		}

		if top != nil {
			current.Neighbourhood.Top = top
			top.Neighbourhood.Bottom = current
		}

		if index%3 == 0 {
			if index%27 == 0 || top == nil {
				current.Group = &Group{Nodes: []*Node{current}}
			} else {
				top.Group.Nodes = append(top.Group.Nodes, current)
				current.Group = top.Group
			}
		} else {
			if previous != nil {
				previous.Group.Nodes = append(previous.Group.Nodes, current)
				current.Group = previous.Group
			}
		}

		if (index+1)%9 == 0 && index != 0 {
			// New row
			previous = nil

			for top = current; top.Neighbourhood.Left != nil; top = top.Neighbourhood.Left {
			}
		} else {
			previous = current

			if top != nil {
				top = top.Neighbourhood.Right
			}
		}

		current = &Node{Neighbourhood: Neighbourhood{}}
	}

	return head
}
