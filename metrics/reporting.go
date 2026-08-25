package metrics

func (c *Counter) Value(name string) int64 { c.mu.Lock(); defer c.mu.Unlock(); return c.Values[name] }
func (c *Counter) Names() []string {
	c.mu.Lock()
	defer c.mu.Unlock()
	out := []string{}
	for k := range c.Values {
		out = append(out, k)
	}
	return out
}
func (c *Counter) Reset(name string) { c.mu.Lock(); defer c.mu.Unlock(); delete(c.Values, name) }
func (c *Counter) ResetAll()         { c.mu.Lock(); defer c.mu.Unlock(); c.Values = map[string]int64{} }
func (c *Counter) Merge(other *Counter) {
	for k, v := range other.Snapshot() {
		c.mu.Lock()
		c.Values[k] += v
		c.mu.Unlock()
	}
}
func (c *Counter) Has(name string) bool { return c.Value(name) > 0 }
func (c *Counter) Total() int64 {
	var n int64
	for _, v := range c.Snapshot() {
		n += v
	}
	return n
}
