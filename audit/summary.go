package audit

import (
	"librarytransfer/model"
	"sort"
	"strings"
)

func Levels(logs []model.TaskLog) map[string]int {
	out := map[string]int{}
	for _, l := range logs {
		out[l.Level]++
	}
	return out
}
func Messages(logs []model.TaskLog) []string {
	out := []string{}
	for _, l := range logs {
		out = append(out, l.Message)
	}
	return out
}
func Sort(logs []model.TaskLog) []model.TaskLog {
	out := append([]model.TaskLog{}, logs...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].At.Before(out[j].At) })
	return out
}
func JoinMessages(logs []model.TaskLog) string {
	parts := Messages(Sort(logs))
	return strings.Join(parts, " | ")
}
func First(logs []model.TaskLog) (model.TaskLog, bool) {
	if len(logs) == 0 {
		return model.TaskLog{}, false
	}
	return Sort(logs)[0], true
}
func Last(logs []model.TaskLog) (model.TaskLog, bool) {
	if len(logs) == 0 {
		return model.TaskLog{}, false
	}
	x := Sort(logs)
	return x[len(x)-1], true
}
