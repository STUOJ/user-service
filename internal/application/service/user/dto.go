package user

import (
	"stuoj-api/application/dto/response"
	"stuoj-common/pkg/utils"
	"user-service/internal/domain/user"
)

func domain2UserSimpleData(u user.User) response.UserSimpleData {
	return response.UserSimpleData{
		Avatar: u.Avatar.String(),
		Id:     u.Id.Value(),
		Name:   u.Username.String(),
		Role:   uint8(u.Role.Value()),
	}
}

func map2UserSimpleData(u map[string]any) response.UserSimpleData {
	var res response.UserSimpleData
	utils.SafeTypeAssert(u["avatar"], &res.Avatar)
	utils.SafeTypeAssert(u["id"], &res.Id)
	utils.SafeTypeAssert(u["username"], &res.Name)
	utils.SafeTypeAssert(u["role"], &res.Role)
	return res
}

func userDomain2Resp(dm user.User) (resp response.UserData) {
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

func map2UserStatistics(m map[string]any) response.UserStatistics {
	var res response.UserStatistics
	utils.SafeTypeAssert(m["ac_count"], &res.ACCount)
	utils.SafeTypeAssert(m["blog_count"], &res.BlogCount)
	utils.SafeTypeAssert(m["submit_count"], &res.SubmitCount)
	return res
}
