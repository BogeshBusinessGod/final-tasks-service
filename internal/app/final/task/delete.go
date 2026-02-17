package task

import (
	"context"
	tsk1 "final/pkg/proto/sync/final-boss/v1"
)

func (s *Server) DeleteTask(ctx context.Context, req *tsk1.DeleteTaskRequest) (*tsk1.DeleteTaskResponse, error) {
	resp, err := s.svc.DeleteTask(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
