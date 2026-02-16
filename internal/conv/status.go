package conv

import (
	"final/internal/models"
	tsk1 "final/pkg/proto/sync/final-boss/v1"
)

func TaskStatusToProto(s models.TaskStatus) tsk1.Status {
	switch s {
	case models.StatusNew:
		return tsk1.Status_STATUS_NEW
	case models.StatusInProgress:
		return tsk1.Status_STATUS_IN_PROGRESS
	case models.StatusDone:
		return tsk1.Status_STATUS_DONE
	case models.StatusDeleted:
		return tsk1.Status_STATUS_DELETED
	case models.StatusError:
		return tsk1.Status_STATUS_ERROR
	default:
		return tsk1.Status_STATUS_UNSPECIFIED
	}
}
