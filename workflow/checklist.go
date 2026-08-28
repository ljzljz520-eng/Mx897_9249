package workflow

import (
	"fmt"
	"librarytransfer/model"
	"librarytransfer/policy"
	"librarytransfer/transfer"
)

type Checklist struct {
	Items    []string
	Complete []bool
}

func NewChecklist(items []string) Checklist {
	return Checklist{Items: append([]string{}, items...), Complete: make([]bool, len(items))}
}
func (c *Checklist) Mark(i int) bool {
	if i < 0 || i >= len(c.Complete) {
		return false
	}
	c.Complete[i] = true
	return true
}
func (c Checklist) Done() bool {
	for _, v := range c.Complete {
		if !v {
			return false
		}
	}
	return true
}
func (c Checklist) Remaining() int {
	n := 0
	for _, v := range c.Complete {
		if !v {
			n++
		}
	}
	return n
}
func (c Checklist) Status() string {
	if c.Done() {
		return "complete"
	}
	return fmt.Sprintf("%d remaining", c.Remaining())
}
func (c Checklist) Names() []string { return append([]string{}, c.Items...) }
func BuildTransferChecklist() Checklist {
	return NewChecklist([]string{"connection", "mapping", "validation", "persistence", "confirmation"})
}
func EvaluateReader(r model.Reader) bool { return policy.CanTransfer(r) }
func EvaluateReaders(rs []model.Reader) int {
	n := 0
	for _, r := range rs {
		if EvaluateReader(r) {
			n++
		}
	}
	return n
}
func RunSteps(s *transfer.Service, rs []model.Reader) Checklist {
	c := BuildTransferChecklist()
	for i, x := range s.Prepare(rs) {
		if x.OK {
			c.Mark(i)
		}
	}
	if len(rs) > 0 {
		c.Mark(0)
	}
	return c
}
