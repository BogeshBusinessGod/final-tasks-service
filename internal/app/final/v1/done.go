package v1

import (
	"context"
	tsk1 "final/pkg/proto/sync/final/v1"
	// ← должен быть именно этот
)

func (s *Server) DoneTask(ctx context.Context, req *tsk1.DoneTaskRequest) (*tsk1.DoneTaskResponse, error) {
	resp, err := s.svc.DoneTask(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
