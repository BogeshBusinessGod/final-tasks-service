package v1

import (
	"context"

	tsk1 "final/pkg/proto/sync/final/v1"
)

func (s *Server) CreateTask(ctx context.Context, req *tsk1.CreateTaskRequest) (*tsk1.CreateTaskResponse, error) {
	resp, err := s.svc.CreateTask(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
