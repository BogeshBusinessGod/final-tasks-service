package v1

import (
	"context"
	tsk1 "final/pkg/proto/sync/final-boss/v1"
)

func (s *Server) GetTask(ctx context.Context, req *tsk1.GetTaskRequest) (*tsk1.GetTaskResponse, error) {
	resp, err := s.svc.GetTask(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
