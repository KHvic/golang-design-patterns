package singleton

/*
	Type: Creational
	Purpose: Ensure a class only has one instance, and provide a global point of access to it
	Additional: Global variables are bad because it can be affected by code from anywhere
				Make testing harder
*/

import "sync"

var (
	once sync.Once

	instance []int
)

func Get() []int {
	once.Do(func() {
		instance = make([]int, 0, 0)
	})
	return instance
}
