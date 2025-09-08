package main

import "sync"

type Singleton struct {
	counter int
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	// We use sync.Once to ensure the instance is created only once, even in concurrent environments.
	once.Do(func() {
		instance = &Singleton{
			counter: 0,
		}
	})

	return instance
}
func (s *Singleton) Increment() {
	s.counter++
}

func (s *Singleton) GetCounter() int {
	return s.counter
}
