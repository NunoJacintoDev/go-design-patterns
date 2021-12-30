package main

import "go-design-patterns/behavioral-patterns/state"

func main() {
	start := state.StartState{}
	game := state.GameContext{
		Next: &start,
	}
	for game.Next.ExecuteState(&game) {
	}
}
