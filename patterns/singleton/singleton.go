// sigleton.go
package singleton

import "sync"

type singleton struct {
	value interface{}
}

var instance *singleton
var once sync.Once

func New() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
