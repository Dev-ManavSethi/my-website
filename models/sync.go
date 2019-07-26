package models

import "sync"

var(
	GlobalMutex sync.Mutex
)
