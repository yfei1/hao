package yelp

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/sherr—99/hao/codes"
	pb "github.com/sherr—99/hao/examples/yelp"
	"github.com/sherr—99/hao/status"
)

var userMap map[string]*pb.User = map[string]*pb.User{}

type yelpServer struct {
	pb.UnimplementedYelpServiceServer
}

func (*yelpServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	uuid := uuid.New().String()
	userMap[uuid] = req.GetUser()

	return req.GetUser(), nil
}
func (*yelpServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*yelpServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (*yelpServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*yelpServer) CreateBusiness(ctx context.Context, req *pb.CreateBusinessRequest) (*pb.Business, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBusiness not implemented")
}
func (*yelpServer) UpdateBusiness(ctx context.Context, req *pb.UpdateBusinessRequest) (*pb.Business, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBusiness not implemented")
}
func (*yelpServer) DeleteBusiness(ctx context.Context, req *pb.DeleteBusinessRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBusiness not implemented")
}
func (*yelpServer) GetBusiness(ctx context.Context, req *pb.GetBusinessRequest) (*pb.Business, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBusiness not implemented")
}
func (*yelpServer) ListBusinesses(ctx context.Context, req *pb.ListBusinessesRequest) (*pb.ListBusinessesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBusinesses not implemented")
}
func (*yelpServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
