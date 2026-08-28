package metrics

import (
	"librarytransfer/model"
	"sync"
)

type Counter struct {
	mu     sync.Mutex
	Values map[string]int64
}

func New() *Counter                { return &Counter{Values: map[string]int64{}} }
func (c *Counter) Inc(name string) { c.mu.Lock(); defer c.mu.Unlock(); c.Values[name]++ }
func (c *Counter) AddTask(t model.TransferTask) {
	c.Inc("tasks")
	if t.Status == model.TaskCompleted {
		c.Inc("completed")
	}
	if t.Status == model.TaskFailed {
		c.Inc("failed")
	}
}
func (c *Counter) Snapshot() map[string]int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	out := map[string]int64{}
	for k, v := range c.Values {
		out[k] = v
	}
	return out
}
