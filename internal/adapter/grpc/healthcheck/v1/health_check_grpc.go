package healthv1

import (
	"context"
)

type HealthCheckService struct {
	UnimplementedHealthServer
}

func NewHealthCheckService() *HealthCheckService {
	return &HealthCheckService{}
}

func (h HealthCheckService) Check(ctx context.Context, request *CheckRequest) (*CheckResponse, error) {
	return &CheckResponse{
		Success: true,
	}, nil
}
