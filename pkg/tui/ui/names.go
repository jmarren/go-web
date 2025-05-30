package ui

type ID int

const (
	Layout ID = iota
	InnerLayout
	InnerLeft
	TerminalArea
	InstanceTable
)

var IdStr = map[ID]string{
	Layout:        "Layout",
	InnerLayout:   "InnerLayout",
	InnerLeft:     "InnerLeft",
	TerminalArea:  "TerminalArea",
	InstanceTable: "InstanceTable",
}

func (n ID) String() string {
	return IdStr[n]
}

type Class int

const (
	Black Class = iota
	GenericBox
)

// TerminalID
