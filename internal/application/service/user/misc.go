package user

import (
	"common/application/dto/request"
	"common/infrastructure/persistence/entity"
	"common/pkg/errors"
)

// 检查权限
func isAdminPermission(reqUser request.ReqUser) error {
	if reqUser.Role < entity.RoleAdmin {
		return &errors.ErrUnauthorized
	}
	return nil
}

// 检查权限
func isPermission(id int64, reqUser request.ReqUser) error {
	if reqUser.Id != id && reqUser.Role < entity.RoleAdmin {
		return &errors.ErrUnauthorized
	}
	return nil
}
