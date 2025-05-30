package ui

// import "github.com/rivo/tview"
//
// type Terminal struct {
// 	*tview.TextArea
// 	text []byte
// }
//
// func NewTerminal() *Terminal {
// 	return &Terminal{
// 		TextArea: tview.NewTextArea(),
// 		text:     []byte{},
// 	}
// }
//
// func (t *Terminal) AppendText(p []byte) {
//
// 	// append p to CurrentText
// 	t.text = append(t.text, p...)
//
// 	// if text is longer than 500,
// 	// take only the last 500 bytes
// 	length := len(t.text)
// 	if length > 500 {
// 		t.text = t.text[length-500:]
// 	}
// 	// set text
// 	t.SetText(string(t.text), true)
// }
