package chain_of_responsibilities

type ChainLogger interface {
	Next(string)
}
