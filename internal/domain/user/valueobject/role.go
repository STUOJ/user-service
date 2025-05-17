package valueobject

import (
	"stuoj-common/domain/shared"
	"stuoj-common/infrastructure/persistence/entity"
	"stuoj-common/pkg/errors"
)

type Role struct {
	shared.Valueobject[entity.Role]
}

func NewRole(value entity.Role) Role {
	var s Role
	s.Set(value)
	return s
}

func (s Role) Verify() error {
	if !s.Value().IsValid() {
		return errors.ErrStatus
	}
	return nil
}
