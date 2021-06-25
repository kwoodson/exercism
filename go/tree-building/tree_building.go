package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Cyclic(start int, records []Record) bool {
	// visited := make(map[int]bool)
	r := records[start]
	count := len(records)
	for {
		if r.ID == 0 {
			return false
		}
		if count < 0 {
			return true
		}
		r = records[r.Parent]
		count--
	}
}

func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}
	sort.Slice(records[:], func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	if records[0].ID != 0 {
		return nil, errors.New("No root node found in records.")
	}
	if len(records) > 1 && records[0].ID == records[1].ID {
		return nil, errors.New("Duplicate root.")
	}

	dedupe := make(map[int]bool, 0)
	nodes := make([]Node, 0)
	for _, r := range records {
		if dedupe[r.ID] {
			return nil, errors.New("Duplicate node found.")
		} else {
			dedupe[r.ID] = true
		}
		if r.ID < r.Parent {
			return nil, errors.New("higher id parent of lower id")
		}
		n := Node{
			ID: r.ID,
		}
		nodes = append(nodes, n)
	}

	// build tree
	for _, r := range records {
		if r.ID == 0 {
			if r.Parent != 0 {
				return nil, errors.New("Not a valid records set.")
			}
			continue
		}
		if r.Parent < 0 || r.Parent > len(records)-1 {
			return nil, errors.New("Parent outside of records.")
		}
		n := nodes[r.Parent]
		if n.Children == nil {
			children := make([]*Node, 0)
			n.Children = children
		}
		if r.ID < 0 || r.ID > len(records)-1 {
			return nil, errors.New("Non contigous record.")
		}
		if n.ID == nodes[r.ID].ID {
			return nil, errors.New("Cannot add child to self.")
		}
		n.Children = append(n.Children, &nodes[r.ID])
		nodes[r.Parent] = n
	}

	for _, r := range records {
		if Cyclic(r.ID, records) {
			return nil, errors.New("Detected a cycle in tree")
		}
	}
	return &nodes[0], nil
}
