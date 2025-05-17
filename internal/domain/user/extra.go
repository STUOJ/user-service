package user

import (
	"fmt"
	"stuoj-common/infrastructure/persistence/entity"
	"stuoj-common/infrastructure/persistence/repository/option"
)

const acCountSQL = "(SELECT COUNT(DISTINCT(problem_id)) FROM tbl_submission WHERE tbl_submission.user_id = tbl_user.id AND tbl_submission.status = 3) AS ac_count"

func QueryUserACCount() option.QueryContextOption {
	return func(pqm option.QueryContext) option.QueryContext {
		field := pqm.GetField()
		if field == nil {
			return pqm
		}
		acCountSelector := option.NewSelector(acCountSQL)
		field.AddSelect(*acCountSelector)
		return pqm
	}
}

var blogCountSQL = fmt.Sprintf("(SELECT COUNT(*) FROM tbl_blog WHERE tbl_blog.user_id = tbl_user.id AND tbl_blog.status >= %d) AS blog_count", entity.BlogPublic)

func QueryUserBlogCount() option.QueryContextOption {
	return func(pqm option.QueryContext) option.QueryContext {
		field := pqm.GetField()
		if field == nil {
			return pqm
		}
		blogCountSelector := option.NewSelector(blogCountSQL)
		field.AddSelect(*blogCountSelector)
		return pqm
	}
}

const submissionCountSQL = "(SELECT COUNT(*) FROM tbl_submission WHERE tbl_submission.user_id = tbl_user.id) AS submit_count"

func QueryUserSubmissionCount() option.QueryContextOption {
	return func(pqm option.QueryContext) option.QueryContext {
		field := pqm.GetField()
		if field == nil {
			return pqm
		}
		submissionCountSelector := option.NewSelector(submissionCountSQL)
		field.AddSelect(*submissionCountSelector)
		return pqm
	}
}
