package v1

import (
	"context"
	tsk1 "final/pkg/proto/sync/final-boss/v1"
)

func (s *Server) ListTasks(ctx context.Context, req *tsk1.ListTasksRequest) (*tsk1.ListTasksResponse, error) {
	resp, err := s.svc.ListTasks(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
