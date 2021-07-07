package core

import (
	"sync"

	"github.com/markhuang1212/code-grader/backend/internal/types"
)

type GradeResultCache struct {
	Data map[string]types.GradeResult
	Lock sync.RWMutex
}

func NewGradeResultCache() *GradeResultCache {
	ret := GradeResultCache{
		Data: make(map[string]types.GradeResult),
	}

	return &ret
}

func (c *GradeResultCache) Add(key string, val types.GradeResult) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.Data[key] = val
}

func (c *GradeResultCache) Del(key string) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	delete(c.Data, key)
}

func (c *GradeResultCache) Get(key string) (types.GradeResult, bool) {
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	val, ok := c.Data[key]
	return val, ok
}
