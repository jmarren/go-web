package tui

func (t *Tui) WriteSsh(p []byte) {
	terminal := t.terminal()
	(*terminal).AppendText(p)
}
