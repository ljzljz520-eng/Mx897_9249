package model

const (
	TaskPending     = "pending"
	TaskValidated   = "validated"
	TaskAwaiting    = "awaiting_confirmation"
	TaskCompleted   = "completed"
	TaskFailed      = "failed"
	DecisionApprove = "approve"
	DecisionReject  = "reject"
)

func IsTerminal(s string) bool { return s == TaskCompleted || s == TaskFailed }
func AllowedTransition(from, to string) bool {
	switch from {
	case "":
		return to == TaskPending
	case TaskPending:
		return to == TaskValidated || to == TaskFailed
	case TaskValidated:
		return to == TaskAwaiting || to == TaskFailed
	case TaskAwaiting:
		return to == TaskCompleted || to == TaskFailed
	default:
		return false
	}
}
