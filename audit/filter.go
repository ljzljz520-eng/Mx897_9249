package audit

import "librarytransfer/model"

func FilterLevel(logs []model.TaskLog, level string) []model.TaskLog {
	out := []model.TaskLog{}
	for _, l := range logs {
		if l.Level == level {
			out = append(out, l)
		}
	}
	return out
}
func FilterMessage(logs []model.TaskLog, term string) []model.TaskLog {
	out := []model.TaskLog{}
	for _, l := range logs {
		if l.Message == term {
			out = append(out, l)
		}
	}
	return out
}
func CountErrors(logs []model.TaskLog) int { return len(FilterLevel(logs, "error")) }
func CountInfo(logs []model.TaskLog) int   { return len(FilterLevel(logs, "info")) }
func Any(logs []model.TaskLog, p func(model.TaskLog) bool) bool {
	for _, l := range logs {
		if p(l) {
			return true
		}
	}
	return false
}
func All(logs []model.TaskLog, p func(model.TaskLog) bool) bool {
	for _, l := range logs {
		if !p(l) {
			return false
		}
	}
	return true
}
func Copy(logs []model.TaskLog) []model.TaskLog { return append([]model.TaskLog{}, logs...) }
func Empty(logs []model.TaskLog) bool           { return len(logs) == 0 }
func NonEmpty(logs []model.TaskLog) bool        { return len(logs) > 0 }
func LevelsList(logs []model.TaskLog) []string {
	out := []string{}
	for _, l := range logs {
		out = append(out, l.Level)
	}
	return out
}
func TaskIDs(logs []model.TaskLog) []string {
	out := []string{}
	for _, l := range logs {
		out = append(out, l.TaskID)
	}
	return out
}
func Validate(logs []model.TaskLog) bool {
	for _, l := range logs {
		if l.TaskID == "" || l.Message == "" {
			return false
		}
	}
	return true
}
