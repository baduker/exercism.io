package tree

import "fmt"

const rootID = 0

// Record contains its ID & Parent
type Record struct {
	ID     int
	Parent int
}

// Node contains its ID & Children
type Node struct {
	ID       int
	Children []*Node
}

// Build converts an unordered slice of Records into a tree of Nodes.
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	positions, err := getPositions(records)
	if err != nil {
		return nil, err
	}

	nodes := make([]Node, len(records))
	for index := range positions {
		record := records[positions[index]]
		if err := checkRecord(record, index); err != nil {
			return nil, err
		}
		nodes[index].ID = index
		addChildren(index, nodes, record)
	}
	return &nodes[0], nil
}

// addChildren appends a node to its parent
func addChildren(index int, nodes []Node, record Record) {
	if index != rootID {
		parent := &nodes[record.Parent]
		parent.Children = append(parent.Children, &nodes[index])
	}
}

// checkRecord validates if a record is in sequence or has a valid parent
func checkRecord(record Record, index int) error {
	if record.ID != index {
		return fmt.Errorf("node %d is out of sequence", record.ID)
	}
	hasValidParent := (record.ID > record.Parent) || (record.ID == rootID && record.Parent == rootID)
	if !hasValidParent {
		return fmt.Errorf("node %d can't belong to %d", record.ID, record.Parent)
	}
	return nil
}

// getPositions
func getPositions(records []Record) ([]int, error) {
	positions := make([]int, len(records))
	for index, record := range records {
		if record.ID < rootID || record.ID >= len(records) {
			return nil, fmt.Errorf("invalid record id %d", record.ID)
		}
		positions[record.ID] = index
	}
	return positions, nil
}
