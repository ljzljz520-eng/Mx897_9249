package audit

import (
	"fmt"
	"librarytransfer/model"
	"librarytransfer/storage"
)

func TaskTimeline(s *storage.Store, id string) ([]model.TaskLog, error) { return s.Logs(id) }
func Render(logs []model.TaskLog) string {
	out := ""
	for _, l := range logs {
		out += fmt.Sprintf("%s [%s] %s\n", l.At.Format("2006-01-02"), l.Level, l.Message)
	}
	return out
}
func HasFailure(logs []model.TaskLog) bool {
	for _, l := range logs {
		if l.Level == "error" {
			return true
		}
	}
	return false
}
