package msg

import (
	"fmt"
	"sync"
)

// Mute ...
var (
	Mute = false
	m    sync.Mutex
)

// Printf ...
func Printf(format string, a ...interface{}) {
	if !Mute {
		m.Lock()
		fmt.Printf(format, a...)
		m.Unlock()
	}
}

// Println ...
func Println(a ...interface{}) {
	if !Mute {
		m.Lock()
		fmt.Println(a...)
		m.Unlock()
	}
}
