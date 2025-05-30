package ui

type TreeNode struct {
	id       ID
	children []EasyPrimitive
	Pos      *GridPosition
	class    Class
}

func (t *TreeNode) Id() ID {
	if t.id == 0 {

	}
	return t.id
}

func (t *TreeNode) Children() []EasyPrimitive {
	return t.children
}

func (t *TreeNode) Class() Class {
	return t.class
}

type ITreeNode interface {
	Id() ID
	Children() []EasyPrimitive
	Class() Class
	init() *GridPosition
}
