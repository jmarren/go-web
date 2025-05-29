package ui

type Ui struct {
	Layout    *EasyGrid
	InnerGrid *EasyGrid
	InnerLeft *EasyGrid
	Terminal  *Terminal
}

func New() *Ui {
	return &Ui{}
}
