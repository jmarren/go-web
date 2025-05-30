package ui

type Cell struct {
	row int
	col int
}

func (c *Cell) MoveLeft() {
	c.col--
}

func (c *Cell) MoveRight() {
	c.col++
}

func (c *Cell) MoveDown() {
	c.row++
}

func (c *Cell) MoveUp() {
	c.row--
}
