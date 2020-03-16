package storage

import "ged/model"

type UserStore interface {
	Save(user model.User)
	Find(id uint) model.User
	Delete(user model.User)
	List() []model.User
}
