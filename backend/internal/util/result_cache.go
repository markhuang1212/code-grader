package util

import (
	"sync"
	"time"

	"github.com/markhuang1212/code-grader/backend/internal/types"
)

type ResultCache struct {
	timeout time.Duration
	lock    sync.RWMutex
	data    map[string]*types.GradeResult
}

func (c *ResultCache) Set(key string, val types.GradeResult) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[key] = &val
	go func() {
		time.Sleep(c.timeout)

	}()
}

func (c *ResultCache) Get(key string) (types.GradeResult, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	v, ok := c.data[key]
	if ok {
		return *v, true
	} else {
		return types.GradeResult{}, false
	}
}
