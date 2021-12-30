package state

import "fmt"

type LoseState struct{}

func (l *LoseState) ExecuteState(c *GameContext) bool {
	fmt.Printf("You lose. The correct number was: %d\n", c.SecretNumber)
	return false
}
