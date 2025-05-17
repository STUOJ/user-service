package user

import (
	"stuoj-common/infrastructure/persistence/repository/option"
	"stuoj-common/infrastructure/persistence/repository/querycontext"
	"user-service/internal/domain/user"
)

// Count 统计用户数量
func Count(query querycontext.UserQueryContext) (int64, error) {
	query.Page = option.NewPagination(0, 0)
	count, err := user.Query.Count(query)
	if err != nil {
		return 0, err
	}

	return count, nil
}
