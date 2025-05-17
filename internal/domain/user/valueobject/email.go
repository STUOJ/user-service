package valueobject

import (
	"errors"
	"regexp"
	"stuoj-common/domain/shared"
)

type Email struct {
	shared.Valueobject[string]
}

func (e Email) Verify() error {
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(e.Value()) {
		return errors.New("邮箱格式不正确！")
	}
	return nil
}

func NewEmail(email string) Email {
	var e Email
	e.Set(email)
	return e
}
