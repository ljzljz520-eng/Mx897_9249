package librarytransfer

import (
	"librarytransfer/model"
	"librarytransfer/storage"
	"librarytransfer/transfer"
	"path/filepath"
	"testing"
	"time"
)

func TestBusinessChain24(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x.db"))
	defer s.Close()
	svc := transfer.New(s)
	r := model.NewReader("CHAIN", "Chain", "D", 0, "idle")
	_, _, _ = svc.CreateTask("CHAIN-T", "old", "new", []model.Reader{r})
	done := make(chan error, 1)
	go func() { _, e := svc.Confirm("CHAIN-T", "admin", model.DecisionApprove); done <- e }()
	select {
	case e := <-done:
		if e != nil {
			t.Fatal(e)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("workflow blocked by lock ordering")
	}
}
