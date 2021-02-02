package access

import (
	"smooth-bear.live/lib/database/access/errors"
	"smooth-bear.live/lib/model"
)

func (d *_default) CreateUser(user *model.User) (*model.User, error) {
	result := d.tx.Create(user)

	if user, ok := result.Value.(*model.User); ok {
		return user, result.Error
	}

	if result.Error != nil {
		result.Error = errors.UserAssertionError
	}

	return nil, result.Error
}
