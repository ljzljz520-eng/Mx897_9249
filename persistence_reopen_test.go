package librarytransfer

import (
	"librarytransfer/model"
	"librarytransfer/storage"
	"path/filepath"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "library.db")
	s, e := storage.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r := model.NewReader("REOPEN", "Persisted", "Science", 0, "idle")
	s.SaveReader(r)
	s.SaveTask(model.TransferTask{ID: "T", Status: model.TaskPending})
	s.Close()
	s, e = storage.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	got, e := s.GetReader("REOPEN")
	if e != nil || got.Name != "Persisted" {
		t.Fatalf("%+v %v", got, e)
	}
}
