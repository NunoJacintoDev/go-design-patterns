package state

type GameContext struct {
	SecretNumber int
	Retries      int
	Won          bool
	Next         GameState
}
