package models

type TaskStatus string

const (
	StatusNew   TaskStatus = "new"
	StatusDone  TaskStatus = "done"
	StatusError TaskStatus = "error"
)
