package transfer

import (
	"fmt"
	"librarytransfer/model"
	"librarytransfer/storage"
	"librarytransfer/validate"
	"sync"
	"time"
)

type Service struct {
	Store *storage.Store
	mu    sync.Mutex
}

func New(s *storage.Store) *Service { return &Service{Store: s} }
func (s *Service) CreateTask(id, source, target string, readers []model.Reader) (model.TransferTask, []validate.Issue, error) {
	t := model.TransferTask{ID: id, SourceName: source, TargetName: target, Status: model.TaskPending, CreatedAt: time.Unix(0, 0).UTC(), Total: len(readers)}
	if e := s.Store.SaveTask(t); e != nil {
		return t, nil, e
	}
	issues := validate.ValidateBatch(readers)
	if len(issues) > 0 {
		t.Status = model.TaskFailed
		_ = s.Store.SaveTask(t)
		return t, issues, nil
	}
	t.Status = model.TaskValidated
	for _, r := range readers {
		r.Department = validate.NormalizeDepartment(r.Department)
		if e := s.Store.SaveReader(r); e != nil {
			return t, nil, e
		}
	}
	t.Status = model.TaskAwaiting
	_ = s.Store.SaveTask(t)
	_ = s.Store.SaveLog(model.TaskLog{TaskID: id, Level: "info", Message: "validated", At: time.Unix(0, 0).UTC()})
	return t, nil, nil
}
func (s *Service) Confirm(id, admin, decision string) (model.TransferTask, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	t, e := s.Store.GetTask(id)
	if e != nil {
		return t, e
	}
	if t.Status != model.TaskAwaiting {
		return t, fmt.Errorf("task not awaiting confirmation")
	}
	c := model.Confirmation{TaskID: id, Administrator: admin, Decision: decision, At: time.Unix(0, 0).UTC()}
	if e = s.Store.SaveConfirmation(c); e != nil {
		return t, e
	}
	if decision == model.DecisionApprove {
		t.Status = model.TaskCompleted
		t.Accepted = t.Total
	} else {
		t.Status = model.TaskFailed
		t.Rejected = t.Total
	}
	if e = s.Store.SaveTask(t); e != nil {
		return t, e
	}
	_ = s.Store.SaveLog(model.TaskLog{TaskID: id, Level: "info", Message: "confirmed by " + admin, At: time.Unix(0, 0).UTC()})
	return t, nil
}
func (s *Service) TransferReader(card string, task model.TransferTask) error {
	r, e := s.Store.GetReader(card)
	if e != nil {
		return e
	}
	if task.Status != model.TaskCompleted {
		return fmt.Errorf("task incomplete")
	}
	r.Enabled = true
	return s.Store.UpdateReaderAndTask(r, task)
}
