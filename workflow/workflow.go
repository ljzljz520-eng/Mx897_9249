package workflow

import (
	"fmt"
	"librarytransfer/admin"
	"librarytransfer/model"
	"librarytransfer/transfer"
)

type Chain struct {
	Service *transfer.Service
	Admin   *admin.Manager
}

func New(a *admin.Manager) *Chain { return &Chain{Service: a.Service(), Admin: a} }
func (c *Chain) Run(id string, readers []model.Reader) (model.TransferTask, error) {
	t, issues, e := c.Service.CreateTask(id, "legacy", "new", readers)
	if e != nil {
		return t, e
	}
	if len(issues) > 0 {
		return t, fmt.Errorf("validation failed: %d issues", len(issues))
	}
	return c.Service.Confirm(id, "system", model.DecisionApprove)
}
func (c *Chain) Publish(card string) error {
	t := model.TransferTask{Status: model.TaskCompleted}
	return c.Service.TransferReader(card, t)
}
func (c *Chain) Approve(id, admin string) error {
	_, e := c.Service.Confirm(id, admin, model.DecisionApprove)
	return e
}
func (c *Chain) Reject(id, admin string) error {
	_, e := c.Service.Confirm(id, admin, model.DecisionReject)
	return e
}
func (c *Chain) ValidateOnly(readers []model.Reader) bool {
	_, issues, _ := c.Service.CreateTask("validation-only", "legacy", "new", readers)
	return len(issues) == 0
}
