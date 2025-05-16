package user

import (
	"common/application/dto"
	"common/application/dto/request"
	"common/application/dto/response"
	"common/infrastructure/persistence/entity"
	"common/infrastructure/persistence/repository/querycontext"
	"common/pkg/errors"
	"common/pkg/utils"
	"log"
	"user-service/internal/domain/user"
)

type UserPage struct {
	Users []response.UserData `json:"users"`
	dto.Page
}

// SelectById 根据ID查询用户
func SelectById(id int64, reqUser request.ReqUser) (response.UserQueryData, error) {
	var resp response.UserQueryData

	// 查询
	qc := querycontext.UserQueryContext{}
	qc.Id.Add(id)
	qc.Field.SelectId().SelectUsername().SelectRole().SelectEmail().SelectAvatar().SelectSignature().SelectCreateTime().SelectUpdateTime()
	u0, map_, err := user.Query.SelectOne(qc, user.QueryUserACCount(), user.QueryUserBlogCount(), user.QueryUserSubmissionCount())

	if err != nil {
		return resp, err
	}

	if reqUser.Id != u0.Id.Value() && reqUser.Role < entity.RoleAdmin {
		u0.Email.Set("")
	}

	resp.UserData = domain2Resp(u0)
	resp.UserStatistics = response.Map2UserStatistics(map_)
	return resp, nil
}

// SelectByEmail 根据邮箱查询用户
func SelectByEmail(email string, reqUser request.ReqUser) (response.UserData, error) {
	var resp response.UserData
	qc := querycontext.UserQueryContext{}
	qc.Email.Set(email)
	qc.Field.SelectId().SelectUsername().SelectRole().SelectEmail().SelectAvatar().SelectSignature().SelectCreateTime().SelectUpdateTime()

	// 查询
	dmUser, _, err := user.Query.SelectOne(qc)
	if err != nil {
		return resp, err
	}

	resp = domain2Resp(dmUser)
	return resp, nil
}

// Select 查询所有用户
func Select(params request.QueryUserParams, reqUser request.ReqUser) (UserPage, error) {
	var resp UserPage

	// 查询
	qc := params2Query(params)
	qc.Field.SelectId().SelectUsername().SelectRole().SelectEmail().SelectAvatar().SelectSignature().SelectCreateTime().SelectUpdateTime()
	users, _, err := user.Query.Select(qc)
	if err != nil {
		return resp, err
	}

	for _, u := range users {
		respUser := domain2Resp(u)
		resp.Users = append(resp.Users, respUser)
	}

	resp.Page.Page = qc.Page.Page
	resp.Size = qc.Page.PageSize
	resp.Page.Total, err = Count(qc)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// LoginByEmail 根据邮箱验证密码
func LoginByEmail(req request.UserLoginReq, reqUser request.ReqUser) (string, error) {
	qc := querycontext.UserQueryContext{}
	qc.Email.Set(req.Email)
	qc.Field.SelectId().SelectPassword()

	// 查询
	u0, _, err := user.Query.SelectOne(qc)
	if err != nil {
		return "", err
	}

	// 验证密码
	err = u0.Password.VerifyHash(req.Password)
	if err != nil {
		log.Println(err)
		return "", errors.ErrUnauthorized.WithMessage("密码错误")
	}

	// 生成token
	token, err := utils.GenerateToken(uint64(u0.Id.Value()))
	if err != nil {
		log.Println(err)
		return "", errors.ErrInternalServer.WithMessage("生成Token失败")
	}

	return token, nil
}

func SelectRoleById(id int64) (entity.Role, error) {
	qc := querycontext.UserQueryContext{}
	qc.Id.Add(id)
	qc.Field.SelectRole()

	u0, _, err := user.Query.SelectOne(qc)
	if err != nil {
		return 0, err
	}

	return u0.Role.Value(), nil
}

func Statistics(params request.UserStatisticsParams, reqUser request.ReqUser) (response.StatisticsRes, error) {
	qc := params2Query(params.QueryUserParams)
	qc.GroupBy = params.GroupBy
	resp, err := user.Query.GroupCount(qc)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
