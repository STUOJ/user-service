package user

import (
	"common/application/dto/request"
	"common/application/dto/response"
	"common/infrastructure/persistence/repository/dao"
	"common/infrastructure/persistence/repository/option"
	"common/infrastructure/persistence/repository/querycontext"
	"time"
	"user-service/internal/domain/user"
)

func domain2Resp(dm user.User) (resp response.UserData) {
	resp = response.UserData{
		Id:         dm.Id.Value(),
		Username:   dm.Username.Value(),
		Role:       uint8(dm.Role.Value()),
		Email:      dm.Email.Value(),
		Avatar:     dm.Avatar.Value(),
		Signature:  dm.Signature.Value(),
		CreateTime: dm.CreateTime.String(),
		UpdateTime: dm.UpdateTime.String(),
	}
	return
}

func params2Query(params request.QueryUserParams) (query querycontext.UserQueryContext) {
	if params.EndTime != nil {
		t, err := time.Parse("2006-01-02 15:04:05", *params.EndTime)
		if err == nil {
			query.EndTime.Set(t)
		}
	}
	if params.StartTime != nil {
		t, err := time.Parse("2006-01-02 15:04:05", *params.StartTime)
		if err == nil {
			query.StartTime.Set(t)
		}
	}
	if params.Id != nil {
		ids, err := utils.StringToInt64Slice(*params.Id)
		if err == nil {
			query.Id.Set(ids)
		}
	}
	if params.Role != nil {
		role, err := dao.StringToRoleSlice(*params.Role)
		if err == nil {
			query.Role.Set(role)
		}
	}
	if params.Username != nil {
		query.Username.Set(*params.Username)
	}
	if params.Email != nil {
		query.Email.Set(*params.Email)
	}
	if params.Page != nil && params.Size != nil {
		query.Page = option.NewPagination(*params.Page, *params.Size)
	}
	if params.Order != nil && params.OrderBy != nil {
		query.Sort = option.NewSortQuery(*params.OrderBy, *params.Order)
	}
	return query
}
