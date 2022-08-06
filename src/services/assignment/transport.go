package assignment

import (
	"context"
	assignmentpb "ggclass_log_service/src/pb/assignment"
)

type IService interface {
}

type transport struct {
	service IService
	assignmentpb.UnimplementedLogAssignmentServiceServer
}

func NewTransport(service IService) *transport {
	return &transport{service: service}
}

func (t *transport) GetLogAssignmentByAssignment(ctx context.Context, request *assignmentpb.GetLogAssignmentByAssignmentRequest) (*assignmentpb.GetLogAssignmentByAssignmentResponse, error) {
	return &assignmentpb.GetLogAssignmentByAssignmentResponse{
		Data: nil,
	}, nil
}

func (t *transport) CreateLogAssignment(ctx context.Context, request *assignmentpb.CreateLogAssignmentRequest) (*assignmentpb.CreateLogAssignmentResponse, error) {
	return &assignmentpb.CreateLogAssignmentResponse{
		Message: "done",
	}, nil
}
