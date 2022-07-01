package Node

import (
	"fmt"
	"strings"
	. "sudoku/set"
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
	Value         int
	Neighbourhood *Neighbourhood
	Group         *Group
	TakenValues   *Set
}

func NewNode() *Node {
	return &Node{
		Neighbourhood: &Neighbourhood{},
		TakenValues:   NewSet(),
	}
}

func (node Node) PrintBoard() {
	fmt.Printf("\n")
	for start := &node; start != nil; start = start.Neighbourhood.Bottom {
		for current := start; current != nil; current = current.Neighbourhood.Right {
			fmt.Printf(" %2d ", len(current.Group.Nodes))
			// value := current.Value
			// if value == 0 {
			// 	fmt.Printf(" __ ")
			// 	continue
			// }

			// fmt.Printf(" %2d ", value)
		}

		fmt.Printf("\n")
		fmt.Printf("\n")
	}
}

func (node Node) String() string {

	top := fmt.Sprintf("  %v  \n", node.Neighbourhood.Top != nil)
	middle := fmt.Sprintf("%v %d %v\n", node.Neighbourhood.Left != nil, node.Value, node.Neighbourhood.Right != nil)
	bottom := fmt.Sprintf("  %v  \n", node.Neighbourhood.Bottom != nil)

	return top + middle + bottom
}

func (node *Node) AddToGroup(group *Group) {
	node.Group = group
	group.Nodes = append(group.Nodes, node)
}

func CreateSudoku(input []int) *Node {
	head := NewNode()

	var current *Node = head
	var previous *Node = nil
	var top *Node = nil

	for index, value := range input {

		current.Value = value

		if previous != nil {
			current.Neighbourhood.Left = previous
			previous.Neighbourhood.Right = current
		}

		if top != nil {
			current.Neighbourhood.Top = top
			top.Neighbourhood.Bottom = current
		}

		// 0    3    6

		// 27   30   33

		// 54   57   60

		if index%3 == 0 {
			if index%27 == 0 || (index-3)%27 == 0 || (index-6)%27 == 0 || top == nil {
				current.Group = &Group{Nodes: []*Node{current}}
			} else {
				current.AddToGroup(top.Group)
			}
		} else {
			current.AddToGroup(previous.Group)
		}

		// if index%3 == 0 {
		// 	if index%27 == 0 || top == nil {
		// 		current.Group = &Group{Nodes: []*Node{current}}
		// 	} else {
		// 		top.Group.Nodes = append(top.Group.Nodes, current)
		// 		current.Group = top.Group
		// 	}
		// } else {
		// 	if previous != nil {
		// 		previous.Group.Nodes = append(previous.Group.Nodes, current)
		// 		current.Group = previous.Group
		// 	}
		// }

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

		current = NewNode()

	}

	return head
}

func (group *Group) TakeOutValue(value int, origin *Node) {
	for _, node := range group.Nodes {
		if node.Value != 0 {
			continue
		}

		if node == origin {
			continue
		}

		node.TakeOutValue(value)
	}
}

func (node *Node) TakeOutValue(value int) {
	node.TakenValues.Insert(value)
}

func (node *Node) TakeOutValueInColumn(value int) {
	for current := node.Neighbourhood.Top; current != nil; current = current.Neighbourhood.Top {
		current.TakeOutValue(value)
	}

	for current := node.Neighbourhood.Bottom; current != nil; current = current.Neighbourhood.Bottom {
		current.TakeOutValue(value)
	}
}

func (node *Node) TakeOutValueInRow(value int) {
	for current := node.Neighbourhood.Right; current != nil; current = current.Neighbourhood.Right {
		current.TakeOutValue(value)
	}

	for current := node.Neighbourhood.Left; current != nil; current = current.Neighbourhood.Left {
		current.TakeOutValue(value)
	}
}

func (node *Node) InsertValue(value int) {
	node.Value = value
	node.TakenValues = NewSet()
	node.TakeOutValueInOthers(value)
}

func (node *Node) TakeOutValueInOthers(value int) {
	node.Group.TakeOutValue(value, node)
	node.TakeOutValueInColumn(value)
	node.TakeOutValueInRow(value)
}

func (node *Node) Solve() {
	// for start := node; start != nil; start = start.Neighbourhood.Bottom {
	// 	for current := start; current != nil; current = current.Neighbourhood.Right {
	// 		value := current.Value
	// 		if value == 0 {
	// 			continue
	// 		}

	// 		current.TakeOutValueInOthers(value)
	// 	}
	// }

	row := 0
	for start := node; start != nil; start = start.Neighbourhood.Bottom {
		if row == 6 {
			return
		}

		for current := start; current != nil; current = current.Neighbourhood.Right {
			value := current.Value
			if value == 0 {
				continue
			}

			current.TakeOutValueInOthers(value)
		}

		row++
	}

	// node.SolveForState()
}

func (node *Node) SolveForState() {
	for start := node; start != nil; start = start.Neighbourhood.Bottom {
		for current := start; current != nil; current = current.Neighbourhood.Right {
			value := current.Value

			if value != 0 {
				continue
			}

			allValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

			if current.TakenValues.Len() != len(allValues)-1 {
				continue
			}

			for _, testValue := range allValues {
				if current.TakenValues.Has(testValue) {
					continue
				}

				current.InsertValue(testValue)

				node.SolveForState()
				return
			}

			fmt.Printf("Taken values: %v not valid", current.TakenValues)
		}
	}
}
