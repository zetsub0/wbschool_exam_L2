package calendar

import (
	"sync"
)

var LastID = 0
var LastIdMutex = sync.Mutex{}

type Result struct {
	Result []Event `json:"result"`
}
