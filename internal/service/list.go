package service

import (
	"context"
	tsk1 "final/pkg/proto/sync/final/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) ListTasks(ctx context.Context, req *tsk1.ListTasksRequest) (*tsk1.ListTasksResponse, error) {
	tasks, err := s.DB.ListTasks(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to list tasks")
	}
	resp := &tsk1.ListTasksResponse{}
	for _, t := range tasks {
		resp.Tasks = append(resp.Tasks, &tsk1.Task{
			Id:      t.ID,
			Title:   t.Title,
			Content: t.Content.String,
			Status:  t.Status,
			Done:    t.Done,
		})
	}
	return resp, nil
}
