package ui

import ()

type Ui struct {
	*Grid
}

func New() *Ui {
	u := &Ui{}
	u.SetGrid()
	u.init()
	return u
}

// func (u *Ui) WriteSsh(p []byte) {
// 	terminal, err := GetByName[*TextArea](u.Grid, TerminalNode)
// 	if err != nil {
// 		panic(err)
// 	}
// 	(*terminal).AppendText(p)
// }

func (u *Ui) SetGrid() {
	u.Grid = structureTwo
}
