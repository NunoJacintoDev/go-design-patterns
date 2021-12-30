package channel_singleton

import "sync"

/*
 In this singleton instance,
 we are using the RWMutex type instead of the already known sync.Mutex type.
 The main difference here is that the RWMutex type has two types of locks
 a read lock and a write lock. The read lock, executed by calling the RLock method,
 only waits if a write lock is currently active. At the same time,
 it only blocks a write lock, so that many read actions can be done in parallel.

 It makes a lot of sense; we don't want to block a Goroutine that wants
 to read a value just because another Goroutine is also reading the value-it won't change.
 The sync.RWMutex type helps us to achieve this logic in our code.
*/

type singleton struct {
	count int
	sync.RWMutex
}

var instance singleton

func GetInstance() *singleton {
	return &instance
}

func (s *singleton) AddOne() {
	s.Lock()
	defer s.Unlock()
	s.count++
}

func (s *singleton) GetCount() int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}

func (s *singleton) Stop() {
	quitCh <- true
	close(addCh)
	close(getCountCh)
	close(quitCh)
}
