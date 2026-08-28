package workflow

import (
	"librarytransfer/admin"
	"librarytransfer/model"
	"librarytransfer/storage"
	"path/filepath"
	"testing"
)

func TestWorkflowQueryAndReport(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x.db"))
	defer s.Close()
	a := admin.New(s)
	c := New(a)
	if !c.ValidateOnly([]model.Reader{model.NewReader("R", "A", "D", 0, "idle")}) {
		t.Fatal("validation")
	}
}
