package state

type FinishState struct{}

func (f *FinishState) ExecuteState(c *GameContext) bool {
	if c.Won {
		c.Next = &WinState{}
	} else {
		c.Next = &LoseState{}
	}
	return true
}
