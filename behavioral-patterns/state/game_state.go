package state

type GameState interface {
	ExecuteState(c *GameContext) bool
}
