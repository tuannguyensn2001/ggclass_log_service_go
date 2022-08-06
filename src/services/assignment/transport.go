package assignment

import (
	"context"
	"ggclass_log_service/src/models"
	assignmentpb "ggclass_log_service/src/pb/assignment"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IService interface {
	Create(ctx context.Context, input createAssignmentLogInput) error
	GetAll(ctx context.Context) ([]models.Assignment, error)
	GetByAssignmentId(ctx context.Context, assignmentId int) ([]models.Assignment, error)
}

type transport struct {
	service IService
	assignmentpb.UnimplementedLogAssignmentServiceServer
}

func NewTransport(service IService) *transport {
	return &transport{service: service}
}

func (t *transport) GetLogAssignmentByAssignment(ctx context.Context, request *assignmentpb.GetLogAssignmentByAssignmentRequest) (*assignmentpb.GetLogAssignmentByAssignmentResponse, error) {

	result, err := t.service.GetByAssignmentId(ctx, int(request.AssignmentId))
	if err != nil {
		return nil, status.Error(codes.Internal, "error")
	}

	response := make([]*assignmentpb.LogAssignment, len(result))

	for index, item := range result {
		response[index] = &assignmentpb.LogAssignment{
			AssignmentId: int64(item.AssignmentId),
			Id:           item.Id,
			Action:       item.Action,
			CreatedAt:    timestamppb.New(*item.CreatedAt),
			UpdatedAt:    timestamppb.New(*item.UpdatedAt),
		}
	}

	return &assignmentpb.GetLogAssignmentByAssignmentResponse{
		Data: response,
	}, nil
}

func (t *transport) CreateLogAssignment(ctx context.Context, request *assignmentpb.CreateLogAssignmentRequest) (*assignmentpb.CreateLogAssignmentResponse, error) {

	input := createAssignmentLogInput{
		AssignmentId: int(request.AssignmentId),
		Action:       request.Action,
	}

	err := t.service.Create(ctx, input)

	if err != nil {
		return nil, status.Error(codes.Internal, "create error")
	}

	return &assignmentpb.CreateLogAssignmentResponse{
		Message: "done",
	}, nil
}
