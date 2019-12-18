package singleton

/*
	Type: Creational
	Purpose: Ensure a class only has one instance, and provide a global point of access to it
	Additional: - Global variables are bad because it can be affected by code from anywhere
				- Global shared instance which prevents an object and resources used by this object from being deallocated
				- Creates tightly coupled code. The clients of the Singleton become difficult to test
*/

import "sync"

var (
	once sync.Once

	singleton []int
)

func Get() []int {
	once.Do(func() {
		singleton = make([]int, 0, 0)
	})
	return singleton
}
