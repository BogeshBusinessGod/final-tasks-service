package models

type TaskStatus string

const (
	StatusNew        TaskStatus = "new"
	StatusInProgress TaskStatus = "in_progress"
	StatusDone       TaskStatus = "done"
	StatusDeleted    TaskStatus = "deleted"
	StatusError      TaskStatus = "error"
)
