package internal

import (
	tsk1 "final/pkg/proto/sync/final/v1"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidation_CreateTaskRequest(t *testing.T) {
	tests := []struct {
		name    string
		req     *tsk1.CreateTaskRequest
		wantErr bool
	}{
		{
			name: "valid request",
			req: &tsk1.CreateTaskRequest{
				Title:   "New Task",
				Content: "Some text",
			},
			wantErr: false,
		},
		{
			name: "too short title",
			req: &tsk1.CreateTaskRequest{
				Title:   "a",
				Content: "Valid",
			},
			wantErr: true,
		},
		{
			name: "too long content",
			req: &tsk1.CreateTaskRequest{
				Title:   "Valid title",
				Content: string(make([]byte, 2000)), // слишком длинный
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.ValidateAll()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
