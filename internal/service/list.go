package service

import (
	"context"
	"final/internal/conv"
	tsk1 "final/pkg/proto/sync/final-boss/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) ListTasks(ctx context.Context, _ *tsk1.ListTasksRequest) (*tsk1.ListTasksResponse, error) {
	tasks, err := s.DB.ListTasks(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to list tasks")
	}

	resp := &tsk1.ListTasksResponse{
		Tasks: make([]*tsk1.Task, 0, len(tasks)),
	}

	for _, t := range tasks {
		resp.Tasks = append(resp.Tasks, &tsk1.Task{
			Id:      t.ID,
			Title:   t.Title,
			Content: t.Content,
			Status:  conv.TaskStatusToProto(t.Status),
		})
	}

	return resp, nil
}
