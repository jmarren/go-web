package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var structureTwo = &Grid{
	Grid:    tview.NewGrid(),
	Rows:    []int{-1, -10, -1},
	Columns: []int{-1, -10, -1},
	TreeNode: &TreeNode{
		id: Layout,
		Pos: &GridPosition{
			RowSpan: -1,
			ColSpan: -1,
		},
		children: []EasyPrimitive{
			&Grid{
				Grid:            tview.NewGrid(),
				Rows:            []int{},
				Columns:         []int{-1, -1},
				Title:           " Go Web! ",
				BorderColor:     tcell.ColorWhite,
				BackgroundColor: tcell.ColorBlack,
				TreeNode: &TreeNode{
					id: InnerLayout,
					Pos: &GridPosition{
						Row:     1,
						Column:  1,
						RowSpan: 1,
						ColSpan: 1,
						Focus:   true,
					},
					children: []EasyPrimitive{
						&Grid{
							Grid:            tview.NewGrid(),
							Rows:            []int{-1, -1},
							Columns:         []int{-1, -50, -1},
							Title:           "InnerLeft",
							BackgroundColor: tcell.ColorBlack,
							TreeNode: &TreeNode{
								id: InnerLeft,
								Pos: &GridPosition{
									Row:           0,
									Column:        0,
									RowSpan:       1,
									ColSpan:       1,
									MinGridWidth:  0,
									MinGridHeight: 0,
									Focus:         true,
								},
								children: []EasyPrimitive{
									&Box{
										Box:             tview.NewBox(),
										BackgroundColor: tcell.ColorBlue,
										BorderColor:     tcell.ColorYellow,
										TreeNode: &TreeNode{
											class: Class(GenericBox),
											Pos: &GridPosition{
												Row:           1,
												Column:        1,
												RowSpan:       1,
												ColSpan:       1,
												MinGridWidth:  0,
												MinGridHeight: 0,
												Focus:         true,
											},
										},
									},
									&Box{
										Box:             tview.NewBox(),
										BackgroundColor: tcell.ColorBlack,
										TreeNode: &TreeNode{
											class: GenericBox,
											Pos: &GridPosition{
												Row:     0,
												Column:  0,
												RowSpan: 2,
												ColSpan: 1,
											},
										},
									},
									&TextArea{
										TextArea:        tview.NewTextArea(),
										BackgroundColor: tcell.ColorBlack,
										BorderColor:     tcell.ColorWhite,
										TreeNode: &TreeNode{
											id: TerminalArea,
											Pos: &GridPosition{
												Row:     0,
												Column:  1,
												RowSpan: 1,
												ColSpan: 1,
											},
										},
									},
									&Table{
										Table:           tview.NewTable(),
										BackgroundColor: tcell.ColorBlack,
										Title:           " instances ",
										selected: &Cell{
											row: 1,
											col: 1,
										},
										Data: [][]string{
											{"instance", "status", "uptime", "playbooks"},
											{"devdb", "online", "12m 15s", "db"},
											{"devapp", "online", "10m 05s", "app"},
											{"app", "offline", "--", "--"},
											{"db", "offline", "--", "--"},
										},
										TreeNode: &TreeNode{
											id: InstanceTable,
											Pos: &GridPosition{
												Row:     1,
												Column:  1,
												RowSpan: 1,
												ColSpan: 1,
												Focus:   true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

// var structure = &Grid{
// 	Grid:    tview.NewGrid(),
// 	Name:    Layout,
// 	Rows:    []int{-1, -10, -1},
// 	Columns: []int{-1, -10, -1},
// 	Pos: &GridPosition{
// 		RowSpan: -1,
// 		ColSpan: -1,
// 	},
// 	Items: []ITreeNode{
// 		&Grid{
// 			Grid:            tview.NewGrid(),
// 			Name:            InnerLayout,
// 			Rows:            []int{},
// 			Columns:         []int{-1, -1},
// 			Title:           " Go Web! ",
// 			BorderColor:     tcell.ColorWhite,
// 			BackgroundColor: tcell.ColorBlack,
// 			Items: []ITreeNode{
// 				&Grid{
// 					Grid:            tview.NewGrid(),
// 					Name:            InnerLeft,
// 					Rows:            []int{-1, -1},
// 					Columns:         []int{-1, -50, -1},
// 					Title:           "InnerLeft",
// 					BackgroundColor: tcell.ColorBlack,
// 					Pos: &GridPosition{
// 						Row:           0,
// 						Column:        0,
// 						RowSpan:       1,
// 						ColSpan:       1,
// 						MinGridWidth:  0,
// 						MinGridHeight: 0,
// 						Focus:         true,
// 					},
// 					Items: []EasyPrimitive{
// 						&Box{
// 							Name:            GenericBox,
// 							Box:             tview.NewBox(),
// 							BackgroundColor: tcell.ColorBlue,
// 							BorderColor:     tcell.ColorYellow,
// 							Pos: &GridPosition{
// 								Row:           1,
// 								Column:        1,
// 								RowSpan:       1,
// 								ColSpan:       1,
// 								MinGridWidth:  0,
// 								MinGridHeight: 0,
// 								Focus:         true,
// 							},
// 						},
// 						&Box{
// 							Box:             tview.NewBox(),
// 							BackgroundColor: tcell.ColorBlack,
// 							Pos: &GridPosition{
// 								Row:     0,
// 								Column:  0,
// 								RowSpan: 2,
// 								ColSpan: 1,
// 							},
// 						},
// 						&TextArea{
// 							Name:            TerminalNode,
// 							TextArea:        tview.NewTextArea(),
// 							BackgroundColor: tcell.ColorAliceBlue,
// 							BorderColor:     tcell.ColorRed,
// 							Pos: &GridPosition{
// 								Row:     0,
// 								Column:  1,
// 								RowSpan: 1,
// 								ColSpan: 1,
// 							},
// 						},
// 						&Table{
// 							Name:            TableNode,
// 							Table:           tview.NewTable(),
// 							BackgroundColor: tcell.ColorBlack,
// 							Title:           " instances ",
// 							selected: &Cell{
// 								row: 1,
// 								col: 1,
// 							},
// 							Data: [][]string{
// 								{"instance", "status", "uptime", "playbooks"},
// 								{"devdb", "online", "12m 15s", "db"},
// 								{"devapp", "online", "10m 05s", "app"},
// 								{"app", "offline", "--", "--"},
// 								{"db", "offline", "--", "--"},
// 							},
// 							Pos: &GridPosition{
// 								Row:     1,
// 								Column:  1,
// 								RowSpan: 1,
// 								ColSpan: 1,
// 								Focus:   true,
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	},
// }
