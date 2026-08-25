package transfer

import (
	"librarytransfer/model"
	"librarytransfer/storage"
	"sort"
)

type History struct {
	Tasks []model.TransferTask
	Logs  []model.TaskLog
}

func LoadHistory(s *storage.Store, id string) History {
	t, e := s.GetTask(id)
	if e != nil {
		return History{}
	}
	ls, _ := s.Logs(id)
	return History{Tasks: []model.TransferTask{t}, Logs: ls}
}
func (h History) LastStatus() string {
	if len(h.Tasks) == 0 {
		return ""
	}
	return h.Tasks[len(h.Tasks)-1].Status
}
func (h History) EventCount() int { return len(h.Logs) }
func (h History) Messages() []string {
	out := []string{}
	for _, l := range h.Logs {
		out = append(out, l.Message)
	}
	return out
}
func (h History) OrderedLogs() []model.TaskLog {
	out := append([]model.TaskLog{}, h.Logs...)
	sort.Slice(out, func(i, j int) bool { return out[i].At.Before(out[j].At) })
	return out
}
func (h History) HasApproval() bool {
	for _, l := range h.Logs {
		if l.Message != "" {
			return true
		}
	}
	return false
}
func (h History) Complete() bool { return h.LastStatus() == model.TaskCompleted }
func (h History) Failed() bool   { return h.LastStatus() == model.TaskFailed }
func (h History) Pending() bool {
	return h.LastStatus() == model.TaskPending || h.LastStatus() == model.TaskAwaiting
}
