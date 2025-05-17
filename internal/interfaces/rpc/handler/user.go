package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"stuoj-api/api/pb"
	"stuoj-api/application/converter"
	"user-service/internal/application/service/user"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

// Update 更新用户基本信息
func (s *UserServer) Update(ctx context.Context, req *pb.UserUpdateRequest) (*pb.UserUpdateResponse, error) {
	reqUser := converter.PBToReqUser(req.ReqUser)
	updateReq := converter.PBToUserUpdateReq(req)

	err := user.Update(updateReq, reqUser)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UserUpdateResponse{}, nil
}

// UpdatePassword 更新用户密码
func (s *UserServer) UpdatePassword(ctx context.Context, req *pb.UserForgetPasswordRequest) (*pb.UserForgetPasswordResponse, error) {
	reqUser := converter.PBToReqUser(req.ReqUser)
	updateReq := converter.PBToUserForgetPasswordReq(req)

	err := user.UpdatePassword(updateReq, reqUser)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UserForgetPasswordResponse{}, nil
}

// UpdateRole 更新用户角色
func (s *UserServer) UpdateRole(ctx context.Context, req *pb.UserUpdateRoleRequest) (*pb.UserUpdateRoleResponse, error) {
	reqUser := converter.PBToReqUser(req.ReqUser)
	updateReq := converter.PBToUserUpdateRoleReq(req)

	err := user.UpdateRole(updateReq, reqUser)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UserUpdateRoleResponse{}, nil
}

// SelectById 根据ID查询用户
func (s *UserServer) SelectById(ctx context.Context, req *pb.UserSelectByIdRequest) (*pb.UserSelectByIdResponse, error) {
	reqUser := converter.PBToReqUser(req.ReqUser)

	userData, err := user.SelectById(req.Id, reqUser)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pbUserData := converter.UserDataToPB(userData.UserData)
	pbStats := converter.UserStatisticsToPB(userData.UserStatistics)

	return &pb.UserSelectByIdResponse{
		UserData:       pbUserData,
		UserStatistics: pbStats,
	}, nil
}

// SelectByEmail 根据邮箱查询用户
func (s *UserServer) SelectByEmail(ctx context.Context, req *pb.UserSelectByEmailRequest) (*pb.UserSelectByEmailResponse, error) {
	reqUser := converter.PBToReqUser(req.ReqUser)

	userData, err := user.SelectByEmail(req.Email, reqUser)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pbUserData := converter.UserDataToPB(userData)

	return &pb.UserSelectByEmailResponse{
		UserData: pbUserData,
	}, nil
}

// Select 查询用户列表
func (s *UserServer) Select(ctx context.Context, req *pb.UserSelectRequest) (*pb.UserSelectResponse, error) {
	params := converter.PBToQueryUserParams(req.Params)
	reqUser := converter.PBToReqUser(req.ReqUser)

	result, err := user.Select(params, reqUser)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pbUsers := make([]*pb.UserData, 0, len(result.Users))
	for _, user := range result.Users {
		pbUsers = append(pbUsers, converter.UserDataToPB(user))
	}

	pbPage := converter.PageToPB(result.Page)

	return &pb.UserSelectResponse{
		Users: pbUsers,
		Page:  pbPage,
	}, nil
}

// LoginByEmail 用户登录
func (s *UserServer) LoginByEmail(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	reqUser := converter.PBToReqUser(req.ReqUser)
	loginReq := converter.PBToUserLoginReq(req)

	token, err := user.LoginByEmail(loginReq, reqUser)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return &pb.UserLoginResponse{
		Token: token,
	}, nil
}

// SelectRoleById 查询用户角色
func (s *UserServer) SelectRoleById(ctx context.Context, req *pb.UserSelectRoleByIdRequest) (*pb.UserSelectRoleByIdResponse, error) {
	role, err := user.SelectRoleById(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UserSelectRoleByIdResponse{
		Role: uint32(role),
	}, nil
}

// Statistics 获取用户统计信息
func (s *UserServer) Statistics(ctx context.Context, req *pb.UserStatisticsRequest) (*pb.UserStatisticsResponse, error) {
	params := converter.PBToUserStatisticsParams(req.Params)
	reqUser := converter.PBToReqUser(req.ReqUser)

	stats, err := user.Statistics(params, reqUser)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	items := converter.StatisticsItemsToPB(stats)

	return &pb.UserStatisticsResponse{
		Items: items,
	}, nil
}

// Register 用户注册
func (s *UserServer) Register(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	reqUser := converter.PBToReqUser(req.ReqUser)
	registerReq := converter.PBToUserRegisterReq(req)

	id, err := user.Register(registerReq, reqUser)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UserRegisterResponse{
		Id: id,
	}, nil
}

// Delete 删除用户
func (s *UserServer) Delete(ctx context.Context, req *pb.UserDeleteRequest) (*pb.UserDeleteResponse, error) {
	reqUser := converter.PBToReqUser(req.ReqUser)

	err := user.Delete(req.Id, reqUser)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UserDeleteResponse{}, nil
}

// Count 统计用户数量
func (s *UserServer) Count(ctx context.Context, req *pb.UserCountRequest) (*pb.UserCountResponse, error) {
	params := converter.PBToQueryUserParams(req.Params)

	qc := converter.UserParams2Query(params)
	count, err := user.Count(qc)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UserCountResponse{
		Count: count,
	}, nil
}
