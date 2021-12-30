package mutex

import "sync"

type Counter struct {
	sync.Mutex
	Value int
}
