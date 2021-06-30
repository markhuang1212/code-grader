package cmd

import (
	"sync"
	"time"

	"github.com/markhuang1212/code-grader/types"
)

type cachedItem struct {
	Result   *types.GradeResult
	Deadline time.Time
}

type TestResultCache struct {
	cachedData map[string]*cachedItem
	lock       sync.RWMutex
	Timeout    time.Duration
}

func NewTestResultCache() *TestResultCache {

	cache := TestResultCache{
		cachedData: make(map[string]*cachedItem),
		Timeout:    720 * time.Second,
	}

	return &cache
}

func (c *TestResultCache) Add(id string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cachedData[id] = &cachedItem{
		Result:   nil,
		Deadline: time.Now().Add(c.Timeout),
	}
	go func() {
		delete(c.cachedData, id)
	}()
}

func (c *TestResultCache) Update(id string, result types.GradeResult) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.cachedData[id] != nil {
		c.cachedData[id].Result = &result
	}
}

func (c *TestResultCache) Get(id string) (*cachedItem, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	item, ok := c.cachedData[id]
	return item, ok
}
