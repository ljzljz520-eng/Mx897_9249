package transfer

import (
	"librarytransfer/model"
	"librarytransfer/storage"
	"path/filepath"
	"testing"
)

func TestWorkflowConnectionTransferConfirmation(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x.db"))
	defer s.Close()
	svc := New(s)
	r := model.NewReader("R1", "A", "D", 0, "idle")
	task, issues, e := svc.CreateTask("T1", "old", "new", []model.Reader{r})
	if e != nil || len(issues) != 0 || task.Status != model.TaskAwaiting {
		t.Fatalf("%+v %v %v", task, issues, e)
	}
	done, e := svc.Confirm("T1", "admin", model.DecisionApprove)
	if e != nil || done.Status != model.TaskCompleted {
		t.Fatalf("%+v %v", done, e)
	}
}
