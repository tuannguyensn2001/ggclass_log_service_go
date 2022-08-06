package routes

import (
	"context"
	assignmentpb "ggclass_log_service/src/pb/assignment"
	"ggclass_log_service/src/services/assignment"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func RegisterGrpcServer(server *grpc.Server) {
	assignmentpb.RegisterLogAssignmentServiceServer(server, buildAssignmentTransport())
}

func RegisterGrpcGateway(ctx context.Context, m *runtime.ServeMux, conn *grpc.ClientConn) error {
	err := assignmentpb.RegisterLogAssignmentServiceHandler(ctx, m, conn)
	if err != nil {
		return err
	}

	return nil
}

func buildAssignmentTransport() assignmentpb.LogAssignmentServiceServer {

	transport := assignment.NewTransport(nil)

	return transport

}
