package state

type WinState struct{}

func (w *WinState) ExecuteState(c *GameContext) bool {
	println("Congrats, you won")
	return false
}
