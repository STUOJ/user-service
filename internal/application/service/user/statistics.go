package user

import (
	"stuoj-api/application/converter"
	"stuoj-api/application/dto/request"
	"stuoj-api/application/dto/response"
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

func Statistics(params request.UserStatisticsParams, reqUser request.ReqUser) (response.StatisticsRes, error) {
	qc := converter.UserParams2Query(params.QueryUserParams)
	qc.GroupBy = params.GroupBy
	resp, err := user.Query.GroupCount(qc)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
