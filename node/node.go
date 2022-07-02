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

type Row struct {
	Nodes []*Node
}

type Column struct {
	Nodes []*Node
}

type Node struct {
	Value         int
	Neighbourhood *Neighbourhood
	Group         *Group
	Row           *Row
	Column        *Column
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
			value := current.Value
			if value == 0 {
				fmt.Printf(" __ ")
				continue
			}

			fmt.Printf(" %2d ", value)
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

func (node *Node) AddToRow(row *Row) {
	node.Row = row
	row.Nodes = append(row.Nodes, node)
}

func (node *Node) AddToColumn(column *Column) {
	node.Column = column
	column.Nodes = append(column.Nodes, node)
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
			current.AddToRow(previous.Row)
		} else {
			current.AddToRow(&Row{Nodes: []*Node{}})
		}

		if top != nil {
			current.Neighbourhood.Top = top
			top.Neighbourhood.Bottom = current
			current.AddToColumn(top.Column)
		} else {
			current.AddToColumn(&Column{Nodes: []*Node{}})
		}

		if index%3 == 0 {
			if (len(current.Column.Nodes)-1)%3 == 0 || top == nil {
				current.Group = &Group{Nodes: []*Node{current}}
			} else {
				current.AddToGroup(top.Group)
			}
		} else {
			current.AddToGroup(previous.Group)
		}

		if (index+1)%9 == 0 && index != 0 {
			// New row
			previous = nil

			top = current.Row.Nodes[0]
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
	for start := node; start != nil; start = start.Neighbourhood.Bottom {
		for current := start; current != nil; current = current.Neighbourhood.Right {
			value := current.Value
			if value == 0 {
				continue
			}

			current.TakeOutValueInOthers(value)
		}
	}

	node.SolveForState()
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
