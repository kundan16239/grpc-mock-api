package server

import (
	"context"
	"go-project/pkg/models"
	"go-project/pkg/services"
	pb "go-project/proto_gen"

	"github.com/google/uuid"
)

type UserGrpcServer struct {
	UserService *services.UserService
	pb.UnimplementedUserServiceServer
}

func NewUserGrpcServer(userService *services.UserService) *UserGrpcServer {
	return &UserGrpcServer{UserService: userService}
}

func (s *UserGrpcServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := &models.User{
		ID:    uuid.New().String(), // generate new ID
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}

	if err := s.UserService.CreateUser(user); err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// GetUserById
func (s *UserGrpcServer) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	user, err := s.UserService.GetUserByID(req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.GetUserByIdResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// UpdateUser
func (s *UserGrpcServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user := &models.User{
		ID:    req.GetId(),
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}
	if err := s.UserService.UpdateUser(user); err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// DeleteUser
func (s *UserGrpcServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	if err := s.UserService.DeleteUser(req.GetId()); err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{Success: true}, nil
}

// GetAllUsers
func (s *UserGrpcServer) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	users, err := s.UserService.GetAllUsers()
	if err != nil {
		return nil, err
	}
	resp := &pb.GetAllUsersResponse{}
	for _, user := range users {
		resp.Users = append(resp.Users, &pb.CreateUserResponse{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return resp, nil
}
