package usecase

import (
	"context"

	"github.com/seki-shinnosuke/grpc-gateway/gen/go/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TodoUsecase struct {
	grpc.UnimplementedTodoApiServer
}

func NewTodoUsecase() (*TodoUsecase, error) {
	return &TodoUsecase{}, nil
}

func (t TodoUsecase) GetTasks(request *emptypb.Empty, stream grpc.TodoApi_GetTasksServer) error {
	return nil
}

func (t TodoUsecase) GetTask(ctx context.Context, request *grpc.GetTaskRequest) (*grpc.TodoDetailResponse, error) {
	var response grpc.TodoDetailResponse
	return &response, nil
}

func (t TodoUsecase) CreateTask(ctx context.Context, request *grpc.CreateTaskRequest) (*grpc.TodoDetailResponse, error) {
	var response grpc.TodoDetailResponse
	return &response, nil
}

func (t TodoUsecase) UpdateTask(ctx context.Context, request *grpc.UpdateTaskRequest) (*grpc.TodoDetailResponse, error) {
	var response grpc.TodoDetailResponse
	return &response, nil
}

func (t TodoUsecase) DeleteTask(ctx context.Context, request *grpc.DeleteTaskRequest) (*emptypb.Empty, error) {
	var response emptypb.Empty
	return &response, nil
}
