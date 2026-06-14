package domain

type status string

var (
	PendStatus status = "pending"
	ProcStatus status = "processing"
	DoneStatus status = "done"
)
